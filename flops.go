package main
import (
        "fmt"
        "time"
       )
func main() {
    fmt.Println(" ")
    fmt.Println("   FLOPS Go Program (double Precision), V0.9 25 Mar 2023")
    fmt.Println("                       Ported From                        ")
    fmt.Println("   FLOPS c Program (double Precision), V2.0 18 dec 1992")
    fmt.Println("         by Al Aburto      aburto@marlin.nosc.mil")
    fmt.Println(" ")

    var nulltime float64= 0.0
    var timearray[3] float64
    var tlimit float64 = 0.0 /* Threshold to determine Number of    */
                             /* Loops to run. Fixed at 15.0 seconds.*/

    var t[36] float64         /* Global array used to hold timing    */
                              /* results and other information.      */

    var sa float64 = 0.0
    var sb float64 = 0.0
    var sc float64 = 0.0
    var one float64 = 1.0
    var two float64 = 2.0
    var three float64 = 3.0
    var four float64 = 4.0
    var five float64 = 5.0
    var piref float64 = 0.0
    var piprg float64 = 0.0
    var scale float64 = 0.0
    var pierr float64 = 0.0

    var a0 float64 = 1.0
    var a1 float64 = -0.1666666666671334
    var a2 float64 = 0.833333333809067e-2
    var a3 float64 = 0.198412715551283e-3
    var a4 float64 = 0.27557589750762e-5
    var a5 float64 = 0.2507059876207e-7
    var a6 float64 = 0.164105986683e-9

//    var b0 = 1.0
    var b1 float64 = -0.4999999999982
    var b2 float64 = 0.4166666664651e-1
    var b3 float64 = -0.1388888805755e-2
    var b4 float64 = 0.24801428034e-4
    var b5 float64 = -0.2754213324e-6
    var b6 float64 = 0.20189405e-8
/*
    let c0 = 1.0
    let c1 = 0.99999999668
    let c2 = 0.49999995173
    let c3 = 0.16666704243
    let c4 = 0.4166685027e-1
    let c5 = 0.832672635e-2
    let c6 = 0.140836136e-2
    let c7 = 0.17358267e-3
    let c8 = 0.3931683e-4
*/
    var d1 float64 = 0.3999999946405e-1
    var d2 float64 = 0.96e-3
    var d3 float64 = 0.1233153e-5

    var e2 float64 = 0.48e-3
    var e3 float64 = 0.411051e-6

    var  s float64 = 0.0
    var  u float64 = 0.0
    var  v float64 = 0.0
    var  w float64 = 0.0
    var  x float64 = 0.0

    var loops int64 = 15625
    var nlimit int64 = 0


    var i int64 = 0
    var m int64 = 0
    var n int64 = 0

    /****************************************************/
    /* Set Variable Values.                             */
    /* t[1] references all timing results relative to   */
    /* one million loops.                               */
    /*                                                  */
    /* The program will execute from 31250 to 512000000 */
    /* loops based on a runtime of Module 1 of at least */
    /* tlimit = 15.0 seconds. That is, a runtime of 15  */
    /* seconds for Module 1 is used to determine the    */
    /* number of loops to execute.                      */
    /*                                                  */
    /* No more than nlimit = 512000000 loops are allowed*/
    /****************************************************/

    t[1] = 1.0e+06 / float64(loops)
    tlimit = 15.0
    nlimit = 512000000
/*            51609600 */

    piref = 3.14159265358979324
    one = 1.0
    two = 2.0
    three = 3.0
    four = 4.0
    five = 5.0
    scale = one

    fmt.Println("   Module     error        RunTime      MFLOPS")
    fmt.Println("                            (usec)")

    /*************************/
    /* Initialize the timer. */
    /*************************/


    /*******************************************************/
    /* Module 1.  Calculate integral of df(x)/f(x) defined */
    /*            below.  Result is ln(f(1)). There are 14 */
    /*            double precision operations per loop     */
    /*            ( 7 +, 0 -, 6 *, 1 / ) that are included */
    /*            in the timing.                           */
    /*            50.0% +, 00.0% -, 42.9% *, and 07.1% /   */
    /*******************************************************/

/*
MODULE   FADD   FSUB   FMUL   FDIV   TOTAL  Comment
  1        7      0      6      1      14   7.1%  FDIV's
*/
    n = loops
    sa = 0.0

//    start1 := time.Now()
    for sa < tlimit {
//fmt.Println(sa,tlimit)
        n = 2 * n
        x = one / float64(n)   /*********************/
        s = 0.0                /*  Loop 1.          */
        v = 0.0                /*********************/
        w = one

        loop1 := time.Now()
        for i =1; i <=n-1; i++  {
            v = v + w
            u = v * x
            s = s + (d1 + u * (d2 + u * d3)) / (w + u * (d1 + u * (e2 + u * e3)))
        }
        loop1t := time.Since(loop1)

     sa = float64(loop1t)*1.0e-9
        timearray[1]=sa

        if n == nlimit {
            break
        }
        //fmt.Println("n = %d  time = %d",n,sa)
    }

    scale = 1.0e6/float64(n)

    t[1] = scale

    /****************************************/
    /* Estimate nulltime ('for' loop time). */
    /****************************************/
    start := time.Now()
    for i=1; i<=n-1; i++  {
    }
    dnulltime := time.Since(start)

    timearray[1]=float64(dnulltime)*1.0e-9
    nulltime = t[1] * timearray[1]
    if nulltime < 0.0 {
        nulltime = 0.0
    }
//    Println("Null time {:?}", dnulltime)
    nul := nulltime
    t[2] = t[1] * sa - nulltime

    sa = (d1 + d2 + d3) / (one + d1 + e2 + e3)
    sb = d1

    t[3] = t[2] / 14.0                      /*********************/
    sa = x * (sa + sb + two * s) / two      /* Module 1 Results. */
    sb = one / sa                           /*********************/
    n = int64(float64(40000*int64(sb))/scale)
  sc = sb - 25.2
    t[4] = one / t[3]
    /********************/
    /*  DO NOT REMOVE   */
    /*  THIS PRINTOUT!  */
    /********************/
    res:=fmt.Sprintf("     1    %+.4e       %.4f     %.4f",sc,t[2],t[4])
    fmt.Println(res)
    m = n
    /*******************************************************/
    /* Module 2.  Calculate value of PI from Taylor Series */
    /*            expansion of atan(1.0).  There are 7     */
    /*            double precision operations per loop     */
    /*            ( 3 +, 2 -, 1 *, 1 / ) that are included */
    /*            in the timing.                           */
    /*            42.9% +, 28.6% -, 14.3% *, and 14.3% /   */
    /*******************************************************/

/*
   MODULE   FADD   FSUB   FMUL   FDIV   TOTAL  Comment
     2        3      2      1      1       7   difficult to vectorize.
*/

    s = -five                       /********************/
    sa = -one                       /* Loop 2.          */
                                    /********************/
    l2m2 := time.Now()
    for i=1; i<=m; i++ {
        s = -s
        sa = sa + s
    }
    elapl2m2 := time.Since(l2m2)
    timearray[1]=float64(elapl2m2)*1.0e-9

    t[5] = t[1] * timearray[1]
    if t[5] < 0.0 {
        t[5] = 0.0
 }

    sc = float64(m)

    u = sa        /*********************/
    v = 0.0       /* Loop 3.           */
    w = 0.0       /*********************/
    x = 0.0

    l3m2 := time.Now()
    for i=1; i<=m; i++ {
        s = -s
        sa = sa + s
        u = u + two
        x = x + (s - u)
        v = v - s * u
        w = w + s / u
    }
    elapl3m2 := time.Since(l3m2)
    timearray[1]=float64(elapl3m2)*1.0e-9
    t[6] = t[1] * timearray[1]

    t[7] = (t[6] - t[5]) / 7.0  /*********************/
    m = int64(sa * x / sc)       /*  PI Results       */
    sa = four * w / five        /*********************/
    sb = sa + five / v
    sc = 31.25
    piprg = sb - sc / (v * v * v)
    pierr = piprg - piref
    t[8] = one / t[7]
    /*********************/
    /*   DO NOT REMOVE   */
    /*   THIS PRINTOUT!  */
    /*********************/
    //   printf("     2   %13.4le  %10.4lf  %10.4lf\n",pierr,t[6]-t[5],t[8])
    res=fmt.Sprintf("     2    %+.4e       %.4f     %.4f     ",pierr,t[6]-t[5],t[8])
    fmt.Println(res)

  /*******************************************************/
    /* Module 3.  Calculate integral of sin(x) from 0.0 to */
    /*            PI/3.0 using Trapazoidal Method. Result  */
    /*            is 0.5. There are 17 double precision    */
    /*            operations per loop (6 +, 2 -, 9 *, 0 /) */
    /*            included in the timing.                  */
    /*            35.3% +, 11.8% -, 52.9% *, and 00.0% /   */
    /*******************************************************/

/*
   MODULE   FADD   FSUB   FMUL   FDIV   TOTAL  Comment
     3        6      2      9      0      17   0.0%  FDIV's
*/

    x = piref / (three * float64(m)) /*********************/
    s = 0.0                          /*  Loop 4.          */
    v = 0.0                          /*********************/

    //   dtime(TimeArray)
    start3 := time.Now()
    for i = 1; i <= m-1; i++ {
        v = v + one
        u = v * x
        w = u * u
        s = s + u * ((((((a6 * w - a5) * w + a4) * w - a3) * w + a2) * w + a1) * w + one)
    }
    //   dtime(TimeArray)
    dnulltime3 := time.Since(start3)
    timearray[1] = float64(dnulltime3)*1.0e-9
    t[9] = t[1] * timearray[1] - nulltime

    u = piref / three
    w = u * u
    sa = u * ((((((a6 * w - a5) * w + a4) * w - a3) * w + a2) * w + a1) * w + one)

    t[10] = t[9] / 17.0              /*********************/
    sa = x * (sa + two * s) / two    /* sin(x) Results.   */
    sb = 0.5                         /*********************/
    sc = sa - sb
    t[11] = one / t[10]
    /*********************/
    /*   DO NOT REMOVE   */
    /*   THIS PRINTOUT!  */
    /*********************/
    res=fmt.Sprintf("     3    %+.4e       %.4f     %.4f      ",sc,t[9],t[11])
    fmt.Println(res)


    /************************************************************/
    /* Module 4.  Calculate Integral of cos(x) from 0.0 to PI/3 */
    /*            using the Trapazoidal Method. Result is       */
    /*            sin(PI/3). There are 15 double precision      */
    /*            operations per loop (7 +, 0 -, 8 *, and 0 / ) */
    /*            included in the timing.                       */
    /*            50.0% +, 00.0% -, 50.0% *, 00.0% /            */
    /************************************************************/

/*
   MODULE   FADD   FSUB   FMUL   FDIV   TOTAL  Comment
     4        7      0      8      0      15   0.0%  FDIV's
*/
    a3 = -a3
    a5 = -a5
    x = piref / (three * float64(m)) /*********************/
    s = 0.0                          /*  Loop 5.          */
    v = 0.0                          /*********************/

    //   dtime(TimeArray)
    start4 := time.Now()
    for i=1; i<m; i++ {
        u = float64(i) * x
        w = u * u
    sa = w * (w * (w * (w * (w * (b6 * w + b5) + b4) + b3) + b2) + b1) + one

    t[13] = t[12] / 15.0                 /*******************/
    sa = x * (sa + one + two * s) / two  /* Module 4 Result */
    u = piref / three                    /*******************/
    w = u * u
    sb = u * ((((((a6 * w + a5) * w + a4) * w + a3) * w + a2) * w + a1) * w + a0)
    sc = sa - sb
    t[14] = one / t[13]
    /*********************/
    /*   DO NOT REMOVE   */
    /*   THIS PRINTOUT!  */
    /*********************/
    res=fmt.Sprintf("     4    %+.4e       %.4f     %.4f      ",sc,t[12],t[14])
    fmt.Println(res)

    /************************************************************/
    /* Module 5.  Calculate Integral of tan(x) from 0.0 to PI/3 */
    /*            using the Trapazoidal Method. Result is       */
    /*            ln(cos(PI/3)). There are 29 double precision  */
    /*            operations per loop (13 +, 0 -, 15 *, and 1 /)*/
    /*            included in the timing.                       */
    /*            46.7% +, 00.0% -, 50.0% *, and 03.3% /        */
    /************************************************************/

 /*
  MODULE   FADD   FSUB   FMUL   FDIV   TOTAL  Comment
     5       13      0     15      1      29   3.4%  FDIV's
*/

    x = piref / (three * float64(m)) /*********************/
    s = 0.0                          /*  Loop 6.          */
    v = 0.0                          /*********************/

    //   dtime(TimeArray)
    start5 := time.Now()
    for i=1; i<m; i++ {
      u = float64(i) * x
        w = u * u
        v = u * ((((((a6 * w + a5) * w + a4) * w + a3) * w + a2) * w + a1) * w + one)
        s = s + v / (w * (w * (w * (w * (w * (b6 * w + b5) + b4) + b3) + b2) + b1) + one)
    }
    //   dtime(TimeArray)
    dnulltime5 := time.Since(start5)
    timearray[1]=float64(dnulltime5)*1.0e-9
    t[15] = t[1] * timearray[1] - nulltime

    u = piref / three
    w = u * u
    sa = u * ((((((a6 * w + a5) * w + a4) * w + a3) * w + a2) * w + a1) * w + one)
    sb = w * (w * (w * (w * (w * (b6 * w + b5) + b4) + b3) + b2) + b1) + one
    sa = sa / sb

    t[16] = t[15] / 29.0           /*******************/
    sa = x * (sa + two * s)  / two /* Module 5 Result */
    sb = 0.6931471805599453        /*******************/
    sc = sa - sb
    t[17] = one / t[16]
    /*********************/
    /*   DO NOT REMOVE   */
    /*   THIS PRINTOUT!  */
    /*********************/
    res=fmt.Sprintf("     5    %+.4e       %.4f     %.4f     ",sc,t[15],t[17])
    fmt.Println(res)

    /************************************************************/
    /* Module 6.  Calculate Integral of sin(x)*cos(x) from 0.0  */
    /*            to PI/4 using the Trapazoidal Method. Result  */
    /*            is sin(PI/4)^2. There are 29 double precision */
    /*            operations per loop (13 +, 0 -, 16 *, and 0 /)*/
    /*            included in the timing.                       */
    /*            46.7% +, 00.0% -, 53.3% *, and 00.0% /        */
    /************************************************************/
/*
   MODULE   FADD   FSUB   FMUL   FDIV   TOTAL  Comment
     6       13      0     16      0      29   0.0%  FDIV's
*/

    x = piref / (four * float64(m)) /*********************/
    s = 0.0                         /*  Loop 7.          */
    v = 0.0                         /*********************/

    //   dtime(TimeArray)
    start6 := time.Now()
    for i=1; i<m; i++  {
        u = float64(i) * x
        w = u * u
        v = u * ((((((a6 * w + a5) * w + a4) * w + a3) * w + a2) * w + a1) * w + one)
        s = s + v * (w * (w * (w * (w * (w * (b6 * w + b5) + b4) + b3) + b2) + b1) + one)
    }
    // dtime(TimeArray)
     dnulltime6 := time.Since(start6)
     timearray[1]=float64(dnulltime6)*1.0e-9
    t[18] = t[1] * timearray[1] - nulltime

    u = piref / four
    w = u * u
    sa = u * ((((((a6 * w + a5) * w + a4) * w + a3) * w + a2) * w + a1) * w + one)
    sb = w * (w * (w * (w * (w * (b6 * w + b5) + b4) + b3) + b2) + b1) + one
    sa = sa * sb

    t[19] = t[18] / 29.0   /*******************/
    sa = x * (sa + two * s) / two /* Module 6 Result */
    sb = 0.25              /*******************/
    sc = sa - sb
    t[20] = one / t[19]
    /*********************/
    /*   DO NOT REMOVE   */
    /*   THIS PRINTOUT!  */
    /*********************/

    res=fmt.Sprintf("     6    %+.4e       %.4f     %.4f     ",sc,t[18],t[20])
    fmt.Println(res)

    /*******************************************************/
    /* Module 7.  Calculate value of the definite integral */
    /*            from 0 to sa of 1/(x+1), x/(x*x+1), and  */
    /*            x*x/(x*x*x+1) using the Trapizoidal Rule.*/
    /*            There are 12 double precision operations */
    /*            per loop ( 3 +, 3 -, 3 *, and 3 / ) that */
    /*            are included in the timing.              */
    /*            25.0% +, 25.0% -, 25.0% *, and 25.0% /   */
    /*******************************************************/
/*
   MODULE   FADD   FSUB   FMUL   FDIV   TOTAL  Comment
     7        3      3      3      3      12   25.0% FDIV's
*/
                           /*********************/
    s = 0.0               /* Loop 8.           */
    w = one               /*********************/
    sa = 102.3321513995275
    v = sa / float64(m)

    //   dtime(TimeArray)
    start7 := time.Now()
    for i=1; i<m; i++ {
        x = float64(i) * v
        u = x * x
        s = s - w / (x + w) - x / (u + w) - u / (x * u + w)
    }
    //   dtime(TimeArray)
    dnulltime7 := time.Since(start7)
    timearray[1]=float64(dnulltime7)*1.0e-9
    t[21] = t[1] * timearray[1] - nulltime
    /*********************/
    /* Module 7 Results  */
    /*********************/
    t[22] = t[21] / 12.0
    x = sa
    u = x * x
    sa = -w - w / (x + w) - x / (u + w) - u / (x * u + w)
    sa = 18.0 * v * (sa + two * s)

    m = -2000 * int64(sa)
    m = int64(float64(m) / scale)

    sc = sa + 500.2
    t[23] = one / t[22]
    /********************/
    /*  DO NOT REMOVE   */
    /*  THIS PRINTOUT!  */
    /********************/

    res=fmt.Sprintf("     7    %+.4e       %.4f     %.4f     ",sc,t[21],t[23])
    fmt.Println(res)


    /************************************************************/
    /* Module 8.  Calculate Integral of sin(x)*cos(x)*cos(x)    */
    /*            from 0 to PI/3 using the Trapazoidal Method.  */
    /*            Result is (1-cos(PI/3)^3)/3. There are 30     */
    /*            double precision operations per loop included */
    /*            in the timing:                                */
    /*               13 +,     0 -,    17 *          0 /        */
    /*            46.7% +, 00.0% -, 53.3% *, and 00.0% /        */
    /************************************************************/
/*
   MODULE   FADD   FSUB   FMUL   FDIV   TOTAL  Comment
     8       13      0     17      0      30   0.0%  FDIV's
*/

    x = piref / (three * float64(m)) /*********************/
    s = 0.0                          /*  Loop 9.          */
    v = 0.0                          /*********************/

    //   dtime(TimeArray)
    start8 := time.Now()
    for i=1; i<m; i++ {
        u = float64(i) * x
        w = u * u
        v = w * (w * (w * (w * (w * (b6 * w + b5) + b4) + b3) + b2) + b1) + one
        s = s + v * v * u * ((((((a6 * w + a5) * w + a4) * w + a3) * w + a2) * w + a1) * w + one)
    }
    // dtime(TimeArray)
    dnulltime8 := time.Since(start8)
    timearray[1]=float64(dnulltime8)*1.0e-9
    t[24] = t[1] * timearray[1] - nulltime

    u = piref / three
    w = u * u
    sa = u * ((((((a6 * w + a5) * w + a4) * w + a3) * w + a2) * w + a1) * w + one)
    sb = w * (w * (w * (w * (w * (b6 * w + b5) + b4) + b3) + b2) + b1) + one
    sa = sa * sb * sb

    t[25] = t[24] / 30.0          /*******************/
    sa = x * (sa + two * s) / two /* Module 8 Result */
    sb = 0.29166666666666667      /*******************/
    sc = sa - sb
    t[26] = one / t[25]
    /*********************/
    /*   DO NOT REMOVE   */
    /*   THIS PRINTOUT!  */
    /*********************/

    res=fmt.Sprintf("     8    %+.4e       %.4f     %.4f     ",sc,t[24],t[26])
    fmt.Println(res)

    /**************************************************/
    /* MFLOPS(1) output. This is the same weighting   */
    /* used for all previous versions of the flops.c  */
    /* used for all previous versions of the flops.c  */
    /* program. Includes Modules 2 and 3 only.        */
    /**************************************************/
    t[27] = (five * (t[6] - t[5]) + t[9]) / 52.0
    t[28] = one / t[27]

    /**************************************************/
    /* MFLOPS(2) output. This output does not include */
    /* Module 2, but it still does 9.2% FDIV's.       */
    /**************************************************/
    t[29] = t[2] + t[9] + t[12] + t[15] + t[18]
    t[29] = (t[29] + four * t[21]) / 152.0
    t[30] = one / t[29]

    /**************************************************/
    /* MFLOPS(3) output. This output does not include */
    /* Module 2, but it still does 3.4% FDIV's.       */
    /**************************************************/
    t[31] = t[2] + t[9] + t[12] + t[15] + t[18]
    t[31] = (t[31] + t[21] + t[24]) / 146.0
    t[32] = one / t[31]

    /**************************************************/
    /* MFLOPS(4) output. This output does not include */
    /* Module 2, and it does NO FDIV's.               */
    /**************************************************/
    t[33] = (t[9] + t[12] + t[18] + t[24]) / 91.0
    t[34] = one / t[33]

         fmt.Println(" ")
         fmt.Println("   Iterations      = ",m)
         res=fmt.Sprintf("   NullTime (usec) = %.f4",nul*1e6)
         fmt.Println(res)
         res=fmt.Sprintf("   MFLOPS(1)       = %.4f",t[28])
         fmt.Println(res)
         res=fmt.Sprintf("   MFLOPS(2)       = %.4f",t[30])
         fmt.Println(res)
         res=fmt.Sprintf("   MFLOPS(3)       = %.4f",t[32])
         fmt.Println(res)
         res=fmt.Sprintf("   MFLOPS(4)       = %.4f",t[34])
         fmt.Println(res)

    /*********** END ***************/





