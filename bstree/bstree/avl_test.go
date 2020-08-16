package bstree

import (
	"fmt"
	"math/rand"
	"testing"
)

var inputForAvlInsert = "81 387 347 59 318 425 40 456 300 194 11 162 89 228 274 211 445 237 106 495 466 28 258 47 447 287 388 290 15 41 408 331 429 356 131485 26 413 90 63 433 147 78 324 159 353 457 221 189 199 0 205 38 203 355 451 10 105 156 266 328 61 202 283 246 376 2 218 94 77 463 496 420 123 453 137 133 241 33 143 391 378 336 46 107 440 3 52 343 98 351 257 187 410 285 132 53 91 82 384 297 267 271 394 226 302 481 79 66 270 493 86 319 175 385 210 249 403 224 112 32 116 339 286 51 76 140 344 364 305 183 301 102 231 154 322 223 342 208 243 468 166 35 404 157 415 371 39 430 13 200 359 220 370 484 247 65 329 48 256 195"
var inputForAvlDel = "129 92 331 77 386 320 399 162 447 292 388 111 103 318 256 19 307 157 152 175 181 353 295 417 393 470 496 132 20 260422 29 161 60 420 79 454 464 51 257 16 100 39 137 33 61 304 185 9 215 219 14 40 240 0 284 173 235 443 421 390 74 369 70 336 338 472 44 395 174 237 24 293 106148 352 122 371 117 151 265 182 479 231 221 355 11 343 53 166 359 367 408 8 28 153 274 375 374 37 286 467 114 158 259 46 223 224 87 31 146 483 428 426 477 412 36 105 262 289 198 387 430 409 128 159 361 102 418 345 120 81 405 276 466 56 195 80 334 267 478 493 380 308 131 58 384 442 23 253 27 48 13 306 332 177 402 154 330 186 187 462 104 193 398 456"

func TestAvlFind(t *testing.T) {

}

func TestAvlLeftRotate(t *testing.T) {
	tree := NewAVLTree()
	input := make([]int, 0)
	for i := 1; i < 20; i++ {
		input = append(input, i)
		tree.Insert(i, nil)
	}
	//dumpTree(tree.root)
	if !isBalance(tree.root) {
		t.Error("not balance")
	}
	doFindTest(t, tree, input)
}

func TestAvlRightRotate(t *testing.T) {
	tree := NewAVLTree()
	input := make([]int, 0)
	for i := 20; i >= 1; i-- {
		input = append(input, i)
		tree.Insert(i, nil)
	}
	//dumpNodeHeight(tree.root)
	dumpTree(tree.root)
	if !isBalance(tree.root) {
		t.Error("not balance")
	}
	doFindTest(t, tree, input)
}

func TestAvlDoubleRotate(t *testing.T) {
	tree := NewAVLTree()
	input := make([]int, 0)
	for i := 1; i < 8; i++ {
		tree.Insert(i, nil)
		input = append(input, i)
	}

	for i := 15; i >= 8; i-- {
		tree.Insert(i, nil)
		input = append(input, i)
	}
	if !isBalance(tree.root) {
		t.Error("not balance")
	}
	doFindTest(t, tree, input)
}

func TestAvlInsert(t *testing.T) {
	input := genTestInput()
	tree := NewAVLTree()
	for _, v := range input {
		tree.Insert(int(v), nil)
	}
	if !isBalance(tree.root) {
		t.Errorf("values:%v\n", input)
		t.Fail()
	}
	doFindTest(t, tree, input)
}

func doAvlDelTest(t *testing.T, inputSize, size int) {
	input := genTestInputV2(inputSize, size)
	tree := NewAVLTree()
	for _, v := range input {
		tree.Insert(int(v), nil)
	}
	doFindTest(t, tree, input)
	num := len(input) / 2
	for i := 0; i < num; i++ {
		key := int(input[i])
		if _, ok := tree.Find(key); !ok {
			t.Errorf("before delete input:%v key:%d\n", input, key)
			return
		}
		tree.Delete(key)
		if _, ok := tree.Find(key); ok {
			t.Errorf("after delete input:%v key:%d\n", input, key)
			return
		}
		if !isBalance(tree.root) {
			t.Error("not balance")
		}
	}
	doFindTest(t, tree, input[num:])
}

func TestAvlDel(t *testing.T) {
	input := genTestInput()
	//input := genInputFromStr(inputForAvlDel)
	tree := NewAVLTree()
	for _, v := range input {
		tree.Insert(int(v), nil)
	}
	doFindTest(t, tree, input)
	num := len(input) / 2
	for i := 0; i < num; i++ {
		key := int(input[i])
		if _, ok := tree.Find(key); !ok {
			t.Errorf("before delete input:%v key:%d\n", input, key)
			return
		}
		tree.Delete(key)
		if _, ok := tree.Find(key); ok {
			t.Errorf("after delete input:%v key:%d\n", input, key)
			return
		}
		if !isBalance(tree.root) {
			t.Error("not balance")
		}
	}
	doFindTest(t, tree, input[num:])
}

func TestAvlDel2(t *testing.T) {
	for i := 0; i < 100; i++ {
		doAvlDelTest(t, 500, 500*(i+1))
	}
}

func BenchmarkAvlFind(b *testing.B) {
	input := genTestInputV2(100000, 1000000)
	tree := NewAVLTree()

	for _, v := range input {
		tree.Insert(v, nil)
	}
	size := len(input)
	fmt.Printf("size:%d\n", size)
	for i := 0; i < b.N; i++ {
		r := rand.Intn(1000000) % size
		if _, ok := tree.Find(input[r]); !ok {
			b.Fail()
			return
		}
	}
}
