// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// 2023-07-18 15:21:31.1196366 +0200 CEST m=+0.033039101
package identity


import (
	A "github.com/IBM/fp-go/internal/apply"
	T "github.com/IBM/fp-go/tuple"
)

// SequenceT1 converts 1 parameters of [T] into a [Tuple1].
func SequenceT1[T1 any](t1 T1) T.Tuple1[T1] {
  return A.SequenceT1(
    Map[T1, T.Tuple1[T1]],
    t1,
  )
}

// SequenceTuple1 converts a [Tuple1] of [T] into an [Tuple1].
func SequenceTuple1[T1 any](t T.Tuple1[T1]) T.Tuple1[T1] {
  return A.SequenceTuple1(
    Map[T1, T.Tuple1[T1]],
    t,
  )
}

// TraverseTuple1 converts a [Tuple1] of [A] via transformation functions transforming [A] to [A] into a [Tuple1].
func TraverseTuple1[F1 ~func(A1) T1, A1, T1 any](f1 F1) func (T.Tuple1[A1]) T.Tuple1[T1] {
  return func(t T.Tuple1[A1]) T.Tuple1[T1] {
    return A.TraverseTuple1(
      Map[T1, T.Tuple1[T1]],
      f1,
      t,
    )
  }
}

// SequenceT2 converts 2 parameters of [T] into a [Tuple2].
func SequenceT2[T1, T2 any](t1 T1, t2 T2) T.Tuple2[T1, T2] {
  return A.SequenceT2(
    Map[T1, func(T2) T.Tuple2[T1, T2]],
    Ap[T.Tuple2[T1, T2], T2],
    t1,
    t2,
  )
}

// SequenceTuple2 converts a [Tuple2] of [T] into an [Tuple2].
func SequenceTuple2[T1, T2 any](t T.Tuple2[T1, T2]) T.Tuple2[T1, T2] {
  return A.SequenceTuple2(
    Map[T1, func(T2) T.Tuple2[T1, T2]],
    Ap[T.Tuple2[T1, T2], T2],
    t,
  )
}

// TraverseTuple2 converts a [Tuple2] of [A] via transformation functions transforming [A] to [A] into a [Tuple2].
func TraverseTuple2[F1 ~func(A1) T1, F2 ~func(A2) T2, A1, T1, A2, T2 any](f1 F1, f2 F2) func (T.Tuple2[A1, A2]) T.Tuple2[T1, T2] {
  return func(t T.Tuple2[A1, A2]) T.Tuple2[T1, T2] {
    return A.TraverseTuple2(
      Map[T1, func(T2) T.Tuple2[T1, T2]],
      Ap[T.Tuple2[T1, T2], T2],
      f1,
      f2,
      t,
    )
  }
}

// SequenceT3 converts 3 parameters of [T] into a [Tuple3].
func SequenceT3[T1, T2, T3 any](t1 T1, t2 T2, t3 T3) T.Tuple3[T1, T2, T3] {
  return A.SequenceT3(
    Map[T1, func(T2) func(T3) T.Tuple3[T1, T2, T3]],
    Ap[func(T3) T.Tuple3[T1, T2, T3], T2],
    Ap[T.Tuple3[T1, T2, T3], T3],
    t1,
    t2,
    t3,
  )
}

// SequenceTuple3 converts a [Tuple3] of [T] into an [Tuple3].
func SequenceTuple3[T1, T2, T3 any](t T.Tuple3[T1, T2, T3]) T.Tuple3[T1, T2, T3] {
  return A.SequenceTuple3(
    Map[T1, func(T2) func(T3) T.Tuple3[T1, T2, T3]],
    Ap[func(T3) T.Tuple3[T1, T2, T3], T2],
    Ap[T.Tuple3[T1, T2, T3], T3],
    t,
  )
}

// TraverseTuple3 converts a [Tuple3] of [A] via transformation functions transforming [A] to [A] into a [Tuple3].
func TraverseTuple3[F1 ~func(A1) T1, F2 ~func(A2) T2, F3 ~func(A3) T3, A1, T1, A2, T2, A3, T3 any](f1 F1, f2 F2, f3 F3) func (T.Tuple3[A1, A2, A3]) T.Tuple3[T1, T2, T3] {
  return func(t T.Tuple3[A1, A2, A3]) T.Tuple3[T1, T2, T3] {
    return A.TraverseTuple3(
      Map[T1, func(T2) func(T3) T.Tuple3[T1, T2, T3]],
      Ap[func(T3) T.Tuple3[T1, T2, T3], T2],
      Ap[T.Tuple3[T1, T2, T3], T3],
      f1,
      f2,
      f3,
      t,
    )
  }
}

// SequenceT4 converts 4 parameters of [T] into a [Tuple4].
func SequenceT4[T1, T2, T3, T4 any](t1 T1, t2 T2, t3 T3, t4 T4) T.Tuple4[T1, T2, T3, T4] {
  return A.SequenceT4(
    Map[T1, func(T2) func(T3) func(T4) T.Tuple4[T1, T2, T3, T4]],
    Ap[func(T3) func(T4) T.Tuple4[T1, T2, T3, T4], T2],
    Ap[func(T4) T.Tuple4[T1, T2, T3, T4], T3],
    Ap[T.Tuple4[T1, T2, T3, T4], T4],
    t1,
    t2,
    t3,
    t4,
  )
}

// SequenceTuple4 converts a [Tuple4] of [T] into an [Tuple4].
func SequenceTuple4[T1, T2, T3, T4 any](t T.Tuple4[T1, T2, T3, T4]) T.Tuple4[T1, T2, T3, T4] {
  return A.SequenceTuple4(
    Map[T1, func(T2) func(T3) func(T4) T.Tuple4[T1, T2, T3, T4]],
    Ap[func(T3) func(T4) T.Tuple4[T1, T2, T3, T4], T2],
    Ap[func(T4) T.Tuple4[T1, T2, T3, T4], T3],
    Ap[T.Tuple4[T1, T2, T3, T4], T4],
    t,
  )
}

// TraverseTuple4 converts a [Tuple4] of [A] via transformation functions transforming [A] to [A] into a [Tuple4].
func TraverseTuple4[F1 ~func(A1) T1, F2 ~func(A2) T2, F3 ~func(A3) T3, F4 ~func(A4) T4, A1, T1, A2, T2, A3, T3, A4, T4 any](f1 F1, f2 F2, f3 F3, f4 F4) func (T.Tuple4[A1, A2, A3, A4]) T.Tuple4[T1, T2, T3, T4] {
  return func(t T.Tuple4[A1, A2, A3, A4]) T.Tuple4[T1, T2, T3, T4] {
    return A.TraverseTuple4(
      Map[T1, func(T2) func(T3) func(T4) T.Tuple4[T1, T2, T3, T4]],
      Ap[func(T3) func(T4) T.Tuple4[T1, T2, T3, T4], T2],
      Ap[func(T4) T.Tuple4[T1, T2, T3, T4], T3],
      Ap[T.Tuple4[T1, T2, T3, T4], T4],
      f1,
      f2,
      f3,
      f4,
      t,
    )
  }
}

// SequenceT5 converts 5 parameters of [T] into a [Tuple5].
func SequenceT5[T1, T2, T3, T4, T5 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5) T.Tuple5[T1, T2, T3, T4, T5] {
  return A.SequenceT5(
    Map[T1, func(T2) func(T3) func(T4) func(T5) T.Tuple5[T1, T2, T3, T4, T5]],
    Ap[func(T3) func(T4) func(T5) T.Tuple5[T1, T2, T3, T4, T5], T2],
    Ap[func(T4) func(T5) T.Tuple5[T1, T2, T3, T4, T5], T3],
    Ap[func(T5) T.Tuple5[T1, T2, T3, T4, T5], T4],
    Ap[T.Tuple5[T1, T2, T3, T4, T5], T5],
    t1,
    t2,
    t3,
    t4,
    t5,
  )
}

// SequenceTuple5 converts a [Tuple5] of [T] into an [Tuple5].
func SequenceTuple5[T1, T2, T3, T4, T5 any](t T.Tuple5[T1, T2, T3, T4, T5]) T.Tuple5[T1, T2, T3, T4, T5] {
  return A.SequenceTuple5(
    Map[T1, func(T2) func(T3) func(T4) func(T5) T.Tuple5[T1, T2, T3, T4, T5]],
    Ap[func(T3) func(T4) func(T5) T.Tuple5[T1, T2, T3, T4, T5], T2],
    Ap[func(T4) func(T5) T.Tuple5[T1, T2, T3, T4, T5], T3],
    Ap[func(T5) T.Tuple5[T1, T2, T3, T4, T5], T4],
    Ap[T.Tuple5[T1, T2, T3, T4, T5], T5],
    t,
  )
}

// TraverseTuple5 converts a [Tuple5] of [A] via transformation functions transforming [A] to [A] into a [Tuple5].
func TraverseTuple5[F1 ~func(A1) T1, F2 ~func(A2) T2, F3 ~func(A3) T3, F4 ~func(A4) T4, F5 ~func(A5) T5, A1, T1, A2, T2, A3, T3, A4, T4, A5, T5 any](f1 F1, f2 F2, f3 F3, f4 F4, f5 F5) func (T.Tuple5[A1, A2, A3, A4, A5]) T.Tuple5[T1, T2, T3, T4, T5] {
  return func(t T.Tuple5[A1, A2, A3, A4, A5]) T.Tuple5[T1, T2, T3, T4, T5] {
    return A.TraverseTuple5(
      Map[T1, func(T2) func(T3) func(T4) func(T5) T.Tuple5[T1, T2, T3, T4, T5]],
      Ap[func(T3) func(T4) func(T5) T.Tuple5[T1, T2, T3, T4, T5], T2],
      Ap[func(T4) func(T5) T.Tuple5[T1, T2, T3, T4, T5], T3],
      Ap[func(T5) T.Tuple5[T1, T2, T3, T4, T5], T4],
      Ap[T.Tuple5[T1, T2, T3, T4, T5], T5],
      f1,
      f2,
      f3,
      f4,
      f5,
      t,
    )
  }
}

// SequenceT6 converts 6 parameters of [T] into a [Tuple6].
func SequenceT6[T1, T2, T3, T4, T5, T6 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6) T.Tuple6[T1, T2, T3, T4, T5, T6] {
  return A.SequenceT6(
    Map[T1, func(T2) func(T3) func(T4) func(T5) func(T6) T.Tuple6[T1, T2, T3, T4, T5, T6]],
    Ap[func(T3) func(T4) func(T5) func(T6) T.Tuple6[T1, T2, T3, T4, T5, T6], T2],
    Ap[func(T4) func(T5) func(T6) T.Tuple6[T1, T2, T3, T4, T5, T6], T3],
    Ap[func(T5) func(T6) T.Tuple6[T1, T2, T3, T4, T5, T6], T4],
    Ap[func(T6) T.Tuple6[T1, T2, T3, T4, T5, T6], T5],
    Ap[T.Tuple6[T1, T2, T3, T4, T5, T6], T6],
    t1,
    t2,
    t3,
    t4,
    t5,
    t6,
  )
}

// SequenceTuple6 converts a [Tuple6] of [T] into an [Tuple6].
func SequenceTuple6[T1, T2, T3, T4, T5, T6 any](t T.Tuple6[T1, T2, T3, T4, T5, T6]) T.Tuple6[T1, T2, T3, T4, T5, T6] {
  return A.SequenceTuple6(
    Map[T1, func(T2) func(T3) func(T4) func(T5) func(T6) T.Tuple6[T1, T2, T3, T4, T5, T6]],
    Ap[func(T3) func(T4) func(T5) func(T6) T.Tuple6[T1, T2, T3, T4, T5, T6], T2],
    Ap[func(T4) func(T5) func(T6) T.Tuple6[T1, T2, T3, T4, T5, T6], T3],
    Ap[func(T5) func(T6) T.Tuple6[T1, T2, T3, T4, T5, T6], T4],
    Ap[func(T6) T.Tuple6[T1, T2, T3, T4, T5, T6], T5],
    Ap[T.Tuple6[T1, T2, T3, T4, T5, T6], T6],
    t,
  )
}

// TraverseTuple6 converts a [Tuple6] of [A] via transformation functions transforming [A] to [A] into a [Tuple6].
func TraverseTuple6[F1 ~func(A1) T1, F2 ~func(A2) T2, F3 ~func(A3) T3, F4 ~func(A4) T4, F5 ~func(A5) T5, F6 ~func(A6) T6, A1, T1, A2, T2, A3, T3, A4, T4, A5, T5, A6, T6 any](f1 F1, f2 F2, f3 F3, f4 F4, f5 F5, f6 F6) func (T.Tuple6[A1, A2, A3, A4, A5, A6]) T.Tuple6[T1, T2, T3, T4, T5, T6] {
  return func(t T.Tuple6[A1, A2, A3, A4, A5, A6]) T.Tuple6[T1, T2, T3, T4, T5, T6] {
    return A.TraverseTuple6(
      Map[T1, func(T2) func(T3) func(T4) func(T5) func(T6) T.Tuple6[T1, T2, T3, T4, T5, T6]],
      Ap[func(T3) func(T4) func(T5) func(T6) T.Tuple6[T1, T2, T3, T4, T5, T6], T2],
      Ap[func(T4) func(T5) func(T6) T.Tuple6[T1, T2, T3, T4, T5, T6], T3],
      Ap[func(T5) func(T6) T.Tuple6[T1, T2, T3, T4, T5, T6], T4],
      Ap[func(T6) T.Tuple6[T1, T2, T3, T4, T5, T6], T5],
      Ap[T.Tuple6[T1, T2, T3, T4, T5, T6], T6],
      f1,
      f2,
      f3,
      f4,
      f5,
      f6,
      t,
    )
  }
}

// SequenceT7 converts 7 parameters of [T] into a [Tuple7].
func SequenceT7[T1, T2, T3, T4, T5, T6, T7 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7) T.Tuple7[T1, T2, T3, T4, T5, T6, T7] {
  return A.SequenceT7(
    Map[T1, func(T2) func(T3) func(T4) func(T5) func(T6) func(T7) T.Tuple7[T1, T2, T3, T4, T5, T6, T7]],
    Ap[func(T3) func(T4) func(T5) func(T6) func(T7) T.Tuple7[T1, T2, T3, T4, T5, T6, T7], T2],
    Ap[func(T4) func(T5) func(T6) func(T7) T.Tuple7[T1, T2, T3, T4, T5, T6, T7], T3],
    Ap[func(T5) func(T6) func(T7) T.Tuple7[T1, T2, T3, T4, T5, T6, T7], T4],
    Ap[func(T6) func(T7) T.Tuple7[T1, T2, T3, T4, T5, T6, T7], T5],
    Ap[func(T7) T.Tuple7[T1, T2, T3, T4, T5, T6, T7], T6],
    Ap[T.Tuple7[T1, T2, T3, T4, T5, T6, T7], T7],
    t1,
    t2,
    t3,
    t4,
    t5,
    t6,
    t7,
  )
}

// SequenceTuple7 converts a [Tuple7] of [T] into an [Tuple7].
func SequenceTuple7[T1, T2, T3, T4, T5, T6, T7 any](t T.Tuple7[T1, T2, T3, T4, T5, T6, T7]) T.Tuple7[T1, T2, T3, T4, T5, T6, T7] {
  return A.SequenceTuple7(
    Map[T1, func(T2) func(T3) func(T4) func(T5) func(T6) func(T7) T.Tuple7[T1, T2, T3, T4, T5, T6, T7]],
    Ap[func(T3) func(T4) func(T5) func(T6) func(T7) T.Tuple7[T1, T2, T3, T4, T5, T6, T7], T2],
    Ap[func(T4) func(T5) func(T6) func(T7) T.Tuple7[T1, T2, T3, T4, T5, T6, T7], T3],
    Ap[func(T5) func(T6) func(T7) T.Tuple7[T1, T2, T3, T4, T5, T6, T7], T4],
    Ap[func(T6) func(T7) T.Tuple7[T1, T2, T3, T4, T5, T6, T7], T5],
    Ap[func(T7) T.Tuple7[T1, T2, T3, T4, T5, T6, T7], T6],
    Ap[T.Tuple7[T1, T2, T3, T4, T5, T6, T7], T7],
    t,
  )
}

// TraverseTuple7 converts a [Tuple7] of [A] via transformation functions transforming [A] to [A] into a [Tuple7].
func TraverseTuple7[F1 ~func(A1) T1, F2 ~func(A2) T2, F3 ~func(A3) T3, F4 ~func(A4) T4, F5 ~func(A5) T5, F6 ~func(A6) T6, F7 ~func(A7) T7, A1, T1, A2, T2, A3, T3, A4, T4, A5, T5, A6, T6, A7, T7 any](f1 F1, f2 F2, f3 F3, f4 F4, f5 F5, f6 F6, f7 F7) func (T.Tuple7[A1, A2, A3, A4, A5, A6, A7]) T.Tuple7[T1, T2, T3, T4, T5, T6, T7] {
  return func(t T.Tuple7[A1, A2, A3, A4, A5, A6, A7]) T.Tuple7[T1, T2, T3, T4, T5, T6, T7] {
    return A.TraverseTuple7(
      Map[T1, func(T2) func(T3) func(T4) func(T5) func(T6) func(T7) T.Tuple7[T1, T2, T3, T4, T5, T6, T7]],
      Ap[func(T3) func(T4) func(T5) func(T6) func(T7) T.Tuple7[T1, T2, T3, T4, T5, T6, T7], T2],
      Ap[func(T4) func(T5) func(T6) func(T7) T.Tuple7[T1, T2, T3, T4, T5, T6, T7], T3],
      Ap[func(T5) func(T6) func(T7) T.Tuple7[T1, T2, T3, T4, T5, T6, T7], T4],
      Ap[func(T6) func(T7) T.Tuple7[T1, T2, T3, T4, T5, T6, T7], T5],
      Ap[func(T7) T.Tuple7[T1, T2, T3, T4, T5, T6, T7], T6],
      Ap[T.Tuple7[T1, T2, T3, T4, T5, T6, T7], T7],
      f1,
      f2,
      f3,
      f4,
      f5,
      f6,
      f7,
      t,
    )
  }
}

// SequenceT8 converts 8 parameters of [T] into a [Tuple8].
func SequenceT8[T1, T2, T3, T4, T5, T6, T7, T8 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8) T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8] {
  return A.SequenceT8(
    Map[T1, func(T2) func(T3) func(T4) func(T5) func(T6) func(T7) func(T8) T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]],
    Ap[func(T3) func(T4) func(T5) func(T6) func(T7) func(T8) T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8], T2],
    Ap[func(T4) func(T5) func(T6) func(T7) func(T8) T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8], T3],
    Ap[func(T5) func(T6) func(T7) func(T8) T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8], T4],
    Ap[func(T6) func(T7) func(T8) T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8], T5],
    Ap[func(T7) func(T8) T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8], T6],
    Ap[func(T8) T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8], T7],
    Ap[T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8], T8],
    t1,
    t2,
    t3,
    t4,
    t5,
    t6,
    t7,
    t8,
  )
}

// SequenceTuple8 converts a [Tuple8] of [T] into an [Tuple8].
func SequenceTuple8[T1, T2, T3, T4, T5, T6, T7, T8 any](t T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]) T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8] {
  return A.SequenceTuple8(
    Map[T1, func(T2) func(T3) func(T4) func(T5) func(T6) func(T7) func(T8) T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]],
    Ap[func(T3) func(T4) func(T5) func(T6) func(T7) func(T8) T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8], T2],
    Ap[func(T4) func(T5) func(T6) func(T7) func(T8) T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8], T3],
    Ap[func(T5) func(T6) func(T7) func(T8) T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8], T4],
    Ap[func(T6) func(T7) func(T8) T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8], T5],
    Ap[func(T7) func(T8) T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8], T6],
    Ap[func(T8) T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8], T7],
    Ap[T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8], T8],
    t,
  )
}

// TraverseTuple8 converts a [Tuple8] of [A] via transformation functions transforming [A] to [A] into a [Tuple8].
func TraverseTuple8[F1 ~func(A1) T1, F2 ~func(A2) T2, F3 ~func(A3) T3, F4 ~func(A4) T4, F5 ~func(A5) T5, F6 ~func(A6) T6, F7 ~func(A7) T7, F8 ~func(A8) T8, A1, T1, A2, T2, A3, T3, A4, T4, A5, T5, A6, T6, A7, T7, A8, T8 any](f1 F1, f2 F2, f3 F3, f4 F4, f5 F5, f6 F6, f7 F7, f8 F8) func (T.Tuple8[A1, A2, A3, A4, A5, A6, A7, A8]) T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8] {
  return func(t T.Tuple8[A1, A2, A3, A4, A5, A6, A7, A8]) T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8] {
    return A.TraverseTuple8(
      Map[T1, func(T2) func(T3) func(T4) func(T5) func(T6) func(T7) func(T8) T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]],
      Ap[func(T3) func(T4) func(T5) func(T6) func(T7) func(T8) T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8], T2],
      Ap[func(T4) func(T5) func(T6) func(T7) func(T8) T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8], T3],
      Ap[func(T5) func(T6) func(T7) func(T8) T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8], T4],
      Ap[func(T6) func(T7) func(T8) T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8], T5],
      Ap[func(T7) func(T8) T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8], T6],
      Ap[func(T8) T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8], T7],
      Ap[T.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8], T8],
      f1,
      f2,
      f3,
      f4,
      f5,
      f6,
      f7,
      f8,
      t,
    )
  }
}

// SequenceT9 converts 9 parameters of [T] into a [Tuple9].
func SequenceT9[T1, T2, T3, T4, T5, T6, T7, T8, T9 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8, t9 T9) T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9] {
  return A.SequenceT9(
    Map[T1, func(T2) func(T3) func(T4) func(T5) func(T6) func(T7) func(T8) func(T9) T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]],
    Ap[func(T3) func(T4) func(T5) func(T6) func(T7) func(T8) func(T9) T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9], T2],
    Ap[func(T4) func(T5) func(T6) func(T7) func(T8) func(T9) T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9], T3],
    Ap[func(T5) func(T6) func(T7) func(T8) func(T9) T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9], T4],
    Ap[func(T6) func(T7) func(T8) func(T9) T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9], T5],
    Ap[func(T7) func(T8) func(T9) T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9], T6],
    Ap[func(T8) func(T9) T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9], T7],
    Ap[func(T9) T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9], T8],
    Ap[T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9], T9],
    t1,
    t2,
    t3,
    t4,
    t5,
    t6,
    t7,
    t8,
    t9,
  )
}

// SequenceTuple9 converts a [Tuple9] of [T] into an [Tuple9].
func SequenceTuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9 any](t T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9] {
  return A.SequenceTuple9(
    Map[T1, func(T2) func(T3) func(T4) func(T5) func(T6) func(T7) func(T8) func(T9) T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]],
    Ap[func(T3) func(T4) func(T5) func(T6) func(T7) func(T8) func(T9) T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9], T2],
    Ap[func(T4) func(T5) func(T6) func(T7) func(T8) func(T9) T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9], T3],
    Ap[func(T5) func(T6) func(T7) func(T8) func(T9) T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9], T4],
    Ap[func(T6) func(T7) func(T8) func(T9) T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9], T5],
    Ap[func(T7) func(T8) func(T9) T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9], T6],
    Ap[func(T8) func(T9) T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9], T7],
    Ap[func(T9) T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9], T8],
    Ap[T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9], T9],
    t,
  )
}

// TraverseTuple9 converts a [Tuple9] of [A] via transformation functions transforming [A] to [A] into a [Tuple9].
func TraverseTuple9[F1 ~func(A1) T1, F2 ~func(A2) T2, F3 ~func(A3) T3, F4 ~func(A4) T4, F5 ~func(A5) T5, F6 ~func(A6) T6, F7 ~func(A7) T7, F8 ~func(A8) T8, F9 ~func(A9) T9, A1, T1, A2, T2, A3, T3, A4, T4, A5, T5, A6, T6, A7, T7, A8, T8, A9, T9 any](f1 F1, f2 F2, f3 F3, f4 F4, f5 F5, f6 F6, f7 F7, f8 F8, f9 F9) func (T.Tuple9[A1, A2, A3, A4, A5, A6, A7, A8, A9]) T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9] {
  return func(t T.Tuple9[A1, A2, A3, A4, A5, A6, A7, A8, A9]) T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9] {
    return A.TraverseTuple9(
      Map[T1, func(T2) func(T3) func(T4) func(T5) func(T6) func(T7) func(T8) func(T9) T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]],
      Ap[func(T3) func(T4) func(T5) func(T6) func(T7) func(T8) func(T9) T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9], T2],
      Ap[func(T4) func(T5) func(T6) func(T7) func(T8) func(T9) T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9], T3],
      Ap[func(T5) func(T6) func(T7) func(T8) func(T9) T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9], T4],
      Ap[func(T6) func(T7) func(T8) func(T9) T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9], T5],
      Ap[func(T7) func(T8) func(T9) T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9], T6],
      Ap[func(T8) func(T9) T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9], T7],
      Ap[func(T9) T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9], T8],
      Ap[T.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9], T9],
      f1,
      f2,
      f3,
      f4,
      f5,
      f6,
      f7,
      f8,
      f9,
      t,
    )
  }
}

// SequenceT10 converts 10 parameters of [T] into a [Tuple10].
func SequenceT10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8, t9 T9, t10 T10) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10] {
  return A.SequenceT10(
    Map[T1, func(T2) func(T3) func(T4) func(T5) func(T6) func(T7) func(T8) func(T9) func(T10) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]],
    Ap[func(T3) func(T4) func(T5) func(T6) func(T7) func(T8) func(T9) func(T10) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], T2],
    Ap[func(T4) func(T5) func(T6) func(T7) func(T8) func(T9) func(T10) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], T3],
    Ap[func(T5) func(T6) func(T7) func(T8) func(T9) func(T10) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], T4],
    Ap[func(T6) func(T7) func(T8) func(T9) func(T10) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], T5],
    Ap[func(T7) func(T8) func(T9) func(T10) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], T6],
    Ap[func(T8) func(T9) func(T10) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], T7],
    Ap[func(T9) func(T10) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], T8],
    Ap[func(T10) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], T9],
    Ap[T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], T10],
    t1,
    t2,
    t3,
    t4,
    t5,
    t6,
    t7,
    t8,
    t9,
    t10,
  )
}

// SequenceTuple10 converts a [Tuple10] of [T] into an [Tuple10].
func SequenceTuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10 any](t T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10] {
  return A.SequenceTuple10(
    Map[T1, func(T2) func(T3) func(T4) func(T5) func(T6) func(T7) func(T8) func(T9) func(T10) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]],
    Ap[func(T3) func(T4) func(T5) func(T6) func(T7) func(T8) func(T9) func(T10) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], T2],
    Ap[func(T4) func(T5) func(T6) func(T7) func(T8) func(T9) func(T10) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], T3],
    Ap[func(T5) func(T6) func(T7) func(T8) func(T9) func(T10) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], T4],
    Ap[func(T6) func(T7) func(T8) func(T9) func(T10) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], T5],
    Ap[func(T7) func(T8) func(T9) func(T10) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], T6],
    Ap[func(T8) func(T9) func(T10) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], T7],
    Ap[func(T9) func(T10) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], T8],
    Ap[func(T10) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], T9],
    Ap[T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], T10],
    t,
  )
}

// TraverseTuple10 converts a [Tuple10] of [A] via transformation functions transforming [A] to [A] into a [Tuple10].
func TraverseTuple10[F1 ~func(A1) T1, F2 ~func(A2) T2, F3 ~func(A3) T3, F4 ~func(A4) T4, F5 ~func(A5) T5, F6 ~func(A6) T6, F7 ~func(A7) T7, F8 ~func(A8) T8, F9 ~func(A9) T9, F10 ~func(A10) T10, A1, T1, A2, T2, A3, T3, A4, T4, A5, T5, A6, T6, A7, T7, A8, T8, A9, T9, A10, T10 any](f1 F1, f2 F2, f3 F3, f4 F4, f5 F5, f6 F6, f7 F7, f8 F8, f9 F9, f10 F10) func (T.Tuple10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10]) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10] {
  return func(t T.Tuple10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10]) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10] {
    return A.TraverseTuple10(
      Map[T1, func(T2) func(T3) func(T4) func(T5) func(T6) func(T7) func(T8) func(T9) func(T10) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]],
      Ap[func(T3) func(T4) func(T5) func(T6) func(T7) func(T8) func(T9) func(T10) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], T2],
      Ap[func(T4) func(T5) func(T6) func(T7) func(T8) func(T9) func(T10) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], T3],
      Ap[func(T5) func(T6) func(T7) func(T8) func(T9) func(T10) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], T4],
      Ap[func(T6) func(T7) func(T8) func(T9) func(T10) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], T5],
      Ap[func(T7) func(T8) func(T9) func(T10) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], T6],
      Ap[func(T8) func(T9) func(T10) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], T7],
      Ap[func(T9) func(T10) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], T8],
      Ap[func(T10) T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], T9],
      Ap[T.Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], T10],
      f1,
      f2,
      f3,
      f4,
      f5,
      f6,
      f7,
      f8,
      f9,
      f10,
      t,
    )
  }
}
