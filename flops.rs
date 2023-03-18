use std::time::{Duration, Instant};
fn main() {
    println!("\n");
    println!("   FLOPS Rust Program (double Precision), V0.9 16 Mar 2023");
    println!("                       Ported From                        ");
    println!("   FLOPS c Program (double Precision), V2.0 18 dec 1992");
    println!("         by Al Aburto      aburto@marlin.nosc.mil\n\n");

    let mut nulltime = 0.0;
    let mut timearray: [f64; 3] = [0.0; 3];
    let mut tlimit = 0.0; /* Threshold to determine Number of    */
    /* Loops to run. Fixed at 15.0 seconds.*/

    let mut T: [f64; 36] = [0.0; 36]; /* Global array used to hold timing    */
    /* results and other information.      */

    let mut sa = 0.0;
    let mut sb = 0.0;
    let mut sc = 0.0;
    let mut one = 1.0;
    let mut two = 2.0;
    let mut three = 3.0;
    let mut four = 4.0;
    let mut five = 5.0;
    let mut piref = 0.0;
    let mut piprg = 0.0;
    let mut scale = 0.0;
    let mut pierr = 0.0;

    let a0 = 1.0;
    let a1 = -0.1666666666671334;
    let a2 = 0.833333333809067e-2;
    let mut a3 = 0.198412715551283e-3;
    let a4 = 0.27557589750762e-5;
    let mut a5 = 0.2507059876207e-7;
    let a6 = 0.164105986683e-9;

    let b0 = 1.0;
    let b1 = -0.4999999999982;
    let b2 = 0.4166666664651e-1;
    let b3 = -0.1388888805755e-2;
    let b4 = 0.24801428034e-4;
    let b5 = -0.2754213324e-6;
    let b6 = 0.20189405e-8;

    let c0 = 1.0;
    let c1 = 0.99999999668;
    let c2 = 0.49999995173;
    let c3 = 0.16666704243;
    let c4 = 0.4166685027e-1;
    let c5 = 0.832672635e-2;
    let c6 = 0.140836136e-2;
    let c7 = 0.17358267e-3;
    let c8 = 0.3931683e-4;

    let d1 = 0.3999999946405e-1;
    let d2 = 0.96e-3;
    let d3 = 0.1233153e-5;

    let e2 = 0.48e-3;
    let e3 = 0.411051e-6;

    let mut s = 0.0;
    let mut u = 0.0;
    let mut v = 0.0;
    let mut w = 0.0;
    let mut x = 0.0;

    let loops = 15625;
    let mut nlimit = 0;
    

    let mut i = 0;
    let mut m = 0;
    let mut n = 0;

    /****************************************************/
    /* Set Variable Values.                             */
    /* T[1] references all timing results relative to   */
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

    T[1] = 1.0e+06 / (loops as f64);
    tlimit = 15.0;
    nlimit = 512000000;

    piref = 3.14159265358979324;
    one = 1.0;
    two = 2.0;
    three = 3.0;
    four = 4.0;
    five = 5.0;
    scale = one;

    println!("   Module     error        RunTime      MFLOPS\n");
    println!("                            (usec)\n");

    /*************************/
    /* Initialize the timer. */
    /*************************/

    //   dtime(timearray);
    //   dtime(timearray);

    /*******************************************************/
    /* Module 1.  Calculate integral of df(x)/f(x) defined */
    /*            below.  Result is ln(f(1)). There are 14 */
    /*            double precision operations per loop     */
    /*            ( 7 +, 0 -, 6 *, 1 / ) that are included */
    /*            in the timing.                           */
    /*            50.0% +, 00.0% -, 42.9% *, and 07.1% /   */
    /*******************************************************/
    n = loops;
    sa = 0.0;

    while sa < tlimit {
        n = 2 * n;
        x = one / (n as f64); /*********************/
        s = 0.0; /*  Loop 1.          */
        v = 0.0; /*********************/
        w = one;

        //     dtime(timearray);
        for i in 1..n //- 1
        //= 1 ; i <= n-1 ; i++ )
        {
            v = v + w;
            u = v * x;
            s = s + (d1 + u * (d2 + u * d3)) / (w + u * (d1 + u * (e2 + u * e3)));
        }
        //      dtime(timearray);
        sa = timearray[1];

        if n == nlimit {
            break;
        };
        /* printf(" %10ld  %12.5lf\n",n,sa); */
    }

    scale = 1.0e+06 / (n as f64);
    T[1] = scale;

    /****************************************/
    /* Estimate nulltime ('for' loop time). */
    /****************************************/
    let start = Instant::now();
    for i in 1..n {}
    let dnulltime = start.elapsed();
    println!("Time elapsed in expensive_function() is: {:?}", dnulltime);
    nulltime = T[1] * timearray[1];
    if nulltime < 0.0 {
        nulltime = 0.0;
    }
    println!("{}", nulltime);
    T[2] = T[1] * sa - nulltime;

    sa = (d1 + d2 + d3) / (one + d1 + e2 + e3);
    sb = d1;

    T[3] = T[2] / 14.0; /*********************/
    sa = x * (sa + sb + two * s) / two; /* Module 1 Results. */
    sb = one / sa; /*********************/
    n = (((40000 * sb as i64) as f64) / scale) as i64;
    sc = sb - 25.2;
    T[4] = one / T[3];
    /********************/
    /*  DO NOT REMOVE   */
    /*  THIS PRINTOUT!  */
    /********************/
    //   println!("     1   %13.4le  %10.4lf  %10.4lf\n",sc,T[2],T[4]);
    println!("     1    {:+.4e} {} {}",sc,T[2],T[4]);
    let dnulltime = start.elapsed();
    println!("Time elapsed in expensive_function() is: {:?}", dnulltime);

    m = n;
    /*******************************************************/
    /* Module 2.  Calculate value of PI from Taylor Series */
    /*            expansion of atan(1.0).  There are 7     */
    /*            double precision operations per loop     */
    /*            ( 3 +, 2 -, 1 *, 1 / ) that are included */
    /*            in the timing.                           */
    /*            42.9% +, 28.6% -, 14.3% *, and 14.3% /   */
    /*******************************************************/


    let start2 = Instant::now();
    s = -five;                      /********************/
    sa = -one;                      /* Loop 2.          */
                                    /********************/
    //   dtime(TimeArray);
    for i in 1..m+1 {
        s = -s;
        sa = sa + s;
    }
    //   dtime(TimeArray);
    T[5] = T[1] * timearray[1];
    if T[5] < 0.0 {
        T[5] = 0.0;
    }

    sc = m as f64;

    u = sa; /*********************/
    v = 0.0; /* Loop 3.           */
    w = 0.0; /*********************/
    x = 0.0;

    //   dtime(TimeArray);
    for i in 1..m+1 {
        s = -s;
        sa = sa + s;
        u = u + two;
        x = x + (s - u);
        v = v - s * u;
        w = w + s / u;
    }
    //   dtime(TimeArray);
    T[6] = T[1] * timearray[1];

    T[7] = (T[6] - T[5]) / 7.0; /*********************/
    m = (sa * x / sc) as i64; /*  PI Results       */
    sa = four * w / five; /*********************/
    sb = sa + five / v;
    sc = 31.25;
    piprg = sb - sc / (v * v * v);
    pierr = piprg - piref;
    T[8] = one / T[7];
    /*********************/
    /*   DO NOT REMOVE   */
    /*   THIS PRINTOUT!  */
    /*********************/
    //   printf("     2   %13.4le  %10.4lf  %10.4lf\n",pierr,T[6]-T[5],T[8]);
    println!("     2    {:+.4e} {} {}",pierr,T[2],T[4]);
    let dnulltime2 = start2.elapsed();
    println!("Time elapsed in expensive_function() is: {:?}", dnulltime2);

    /*******************************************************/
    /* Module 3.  Calculate integral of sin(x) from 0.0 to */
    /*            PI/3.0 using Trapazoidal Method. Result  */
    /*            is 0.5. There are 17 double precision    */
    /*            operations per loop (6 +, 2 -, 9 *, 0 /) */
    /*            included in the timing.                  */
    /*            35.3% +, 11.8% -, 52.9% *, and 00.0% /   */
    /*******************************************************/

    x = piref / (three * (m as f64)); /*********************/
    s = 0.0; /*  Loop 4.          */
    v = 0.0; /*********************/

    //   dtime(TimeArray);
    for i in 1..m //- 1 {
        {
        v = v + one;
        u = v * x;
        w = u * u;
        s = s + u * ((((((a6 * w - a5) * w + a4) * w - a3) * w + a2) * w + a1) * w + one);
    }
    //   dtime(TimeArray);
    T[9] = T[1] * timearray[1] - nulltime;

    u = piref / three;
    w = u * u;
    sa = u * ((((((a6 * w - a5) * w + a4) * w - a3) * w + a2) * w + a1) * w + one);

    T[10] = T[9] / 17.0; /*********************/
    sa = x * (sa + two * s) / two; /* sin(x) Results.   */
    sb = 0.5; /*********************/
    sc = sa - sb;
    T[11] = one / T[10];
    /*********************/
    /*   DO NOT REMOVE   */
    /*   THIS PRINTOUT!  */
    /*********************/
    //   printf("     3   %13.4le  %10.4lf  %10.4lf\n",sc,T[9],T[11]);
    println!("     3    {:+.4e} {} {}",sc,T[2],T[4]);

    /************************************************************/
    /* Module 4.  Calculate Integral of cos(x) from 0.0 to PI/3 */
    /*            using the Trapazoidal Method. Result is       */
    /*            sin(PI/3). There are 15 double precision      */
    /*            operations per loop (7 +, 0 -, 8 *, and 0 / ) */
    /*            included in the timing.                       */
    /*            50.0% +, 00.0% -, 50.0% *, 00.0% /            */
    /************************************************************/
    a3 = -a3;
    a5 = -a5;
    x = piref / (three * (m as f64)); /*********************/
    s = 0.0; /*  Loop 5.          */
    v = 0.0; /*********************/

    //   dtime(TimeArray);
    for i in 1..m //- 1 {
        {
        u = (i as f64) * x;
        w = u * u;
        s = s + w * (w * (w * (w * (w * (b6 * w + b5) + b4) + b3) + b2) + b1) + one;
    }
    //   dtime(TimeArray);
    T[12] = T[1] * timearray[1] - nulltime;

    u = piref / three;
    w = u * u;
    sa = w * (w * (w * (w * (w * (b6 * w + b5) + b4) + b3) + b2) + b1) + one;

    T[13] = T[12] / 15.0; /*******************/
    sa = x * (sa + one + two * s) / two; /* Module 4 Result */
    u = piref / three; /*******************/
    w = u * u;
    sb = u * ((((((a6 * w + a5) * w + a4) * w + a3) * w + a2) * w + a1) * w + a0);
    sc = sa - sb;
    T[14] = one / T[13];
    /*********************/
    /*   DO NOT REMOVE   */
    /*   THIS PRINTOUT!  */
    /*********************/
    //   printf("     4   %13.4le  %10.4lf  %10.4lf\n",sc,T[12],T[14]);
    println!("     4    {:+.4e} {} {}",sc,T[2],T[4]);

    /************************************************************/
    /* Module 5.  Calculate Integral of tan(x) from 0.0 to PI/3 */
    /*            using the Trapazoidal Method. Result is       */
    /*            ln(cos(PI/3)). There are 29 double precision  */
    /*            operations per loop (13 +, 0 -, 15 *, and 1 /)*/
    /*            included in the timing.                       */
    /*            46.7% +, 00.0% -, 50.0% *, and 03.3% /        */
    /************************************************************/

    x = piref / (three * (m as f64)); /*********************/
    s = 0.0; /*  Loop 6.          */
    v = 0.0; /*********************/

    //   dtime(TimeArray);
    for i in 1..m { //- 1 {
        u = (i as f64) * x;
        w = u * u;
        v = u * ((((((a6 * w + a5) * w + a4) * w + a3) * w + a2) * w + a1) * w + one);
        s = s + v / (w * (w * (w * (w * (w * (b6 * w + b5) + b4) + b3) + b2) + b1) + one);
    }
    //   dtime(TimeArray);
    T[15] = T[1] * timearray[1] - nulltime;

    u = piref / three;
    w = u * u;
    sa = u * ((((((a6 * w + a5) * w + a4) * w + a3) * w + a2) * w + a1) * w + one);
    sb = w * (w * (w * (w * (w * (b6 * w + b5) + b4) + b3) + b2) + b1) + one;
    sa = sa / sb;

    T[16] = T[15] / 29.0; /*******************/
    sa = x * (sa + two * s) / two; /* Module 5 Result */
    sb = 0.6931471805599453; /*******************/
    sc = sa - sb;
    T[17] = one / T[16];
    /*********************/
    /*   DO NOT REMOVE   */
    /*   THIS PRINTOUT!  */
    /*********************/
    //   printf("     5   %13.4le  %10.4lf  %10.4lf\n",sc,T[15],T[17]);
    println!("     5    {:+.4e} {} {}",sc,T[2],T[4]);

    /************************************************************/
    /* Module 6.  Calculate Integral of sin(x)*cos(x) from 0.0  */
    /*            to PI/4 using the Trapazoidal Method. Result  */
    /*            is sin(PI/4)^2. There are 29 double precision */
    /*            operations per loop (13 +, 0 -, 16 *, and 0 /)*/
    /*            included in the timing.                       */
    /*            46.7% +, 00.0% -, 53.3% *, and 00.0% /        */
    /************************************************************/

    x = piref / (four * (m as f64)); /*********************/
    s = 0.0; /*  Loop 7.          */
    v = 0.0; /*********************/

    //   dtime(TimeArray);
    for i in 1..m { //- 1 {
        u = (i as f64) * x;
        w = u * u;
        v = u * ((((((a6 * w + a5) * w + a4) * w + a3) * w + a2) * w + a1) * w + one);
        s = s + v * (w * (w * (w * (w * (w * (b6 * w + b5) + b4) + b3) + b2) + b1) + one);
    }
    // dtime(TimeArray);
    T[18] = T[1] * timearray[1] - nulltime;

    u = piref / four;
    w = u * u;
    sa = u * ((((((a6 * w + a5) * w + a4) * w + a3) * w + a2) * w + a1) * w + one);
    sb = w * (w * (w * (w * (w * (b6 * w + b5) + b4) + b3) + b2) + b1) + one;
    sa = sa * sb;

    T[19] = T[18] / 29.0; /*******************/
    sa = x * (sa + two * s) / two; /* Module 6 Result */
    sb = 0.25; /*******************/
    sc = sa - sb;
    T[20] = one / T[19];
    /*********************/
    /*   DO NOT REMOVE   */
    /*   THIS PRINTOUT!  */
    /*********************/
    //   printf("     6   %13.4le  %10.4lf  %10.4lf\n",sc,T[18],T[20]);
    println!("     6    {:+.4e} {} {}",sc,T[2],T[4]);

    /*******************************************************/
    /* Module 7.  Calculate value of the definite integral */
    /*            from 0 to sa of 1/(x+1), x/(x*x+1), and  */
    /*            x*x/(x*x*x+1) using the Trapizoidal Rule.*/
    /*            There are 12 double precision operations */
    /*            per loop ( 3 +, 3 -, 3 *, and 3 / ) that */
    /*            are included in the timing.              */
    /*            25.0% +, 25.0% -, 25.0% *, and 25.0% /   */
    /*******************************************************/

    /*********************/
    s = 0.0; /* Loop 8.           */
    w = one; /*********************/
    sa = 102.3321513995275;
    v = sa / (m as f64);

    //   dtime(TimeArray);
    for i in 1..m { //- 1 {
        x = (i as f64) * v;
        u = x * x;
        s = s - w / (x + w) - x / (u + w) - u / (x * u + w);
    }
    //   dtime(TimeArray);
    T[21] = T[1] * timearray[1] - nulltime;
    /*********************/
    /* Module 7 Results  */
    /*********************/
    T[22] = T[21] / 12.0;
    x = sa;
    u = x * x;
    sa = -w - w / (x + w) - x / (u + w) - u / (x * u + w);
    sa = 18.0 * v * (sa + two * s);

    m = -2000 * (sa as i64);
    m = ((m as f64) / scale) as i64;

    sc = sa + 500.2;
    T[23] = one / T[22];
    /********************/
    /*  DO NOT REMOVE   */
    /*  THIS PRINTOUT!  */
    /********************/
    //   printf("     7   %13.4le  %10.4lf  %10.4lf\n",sc,T[21],T[23]);
    println!("     7    {:+.4e} {} {}",sc,T[2],T[4]);

    /************************************************************/
    /* Module 8.  Calculate Integral of sin(x)*cos(x)*cos(x)    */
    /*            from 0 to PI/3 using the Trapazoidal Method.  */
    /*            Result is (1-cos(PI/3)^3)/3. There are 30     */
    /*            double precision operations per loop included */
    /*            in the timing:                                */
    /*               13 +,     0 -,    17 *          0 /        */
    /*            46.7% +, 00.0% -, 53.3% *, and 00.0% /        */
    /************************************************************/

    x = piref / (three * (m as f64)); /*********************/
    s = 0.0; /*  Loop 9.          */
    v = 0.0; /*********************/

    //   dtime(TimeArray);
    for i in 1..m {//- 1 {
        u = (i as f64) * x;
        w = u * u;
        v = w * (w * (w * (w * (w * (b6 * w + b5) + b4) + b3) + b2) + b1) + one;
        s = s + v * v * u * ((((((a6 * w + a5) * w + a4) * w + a3) * w + a2) * w + a1) * w + one);
    }
    // dtime(TimeArray);
    T[24] = T[1] * timearray[1] - nulltime;

    u = piref / three;
    w = u * u;
    sa = u * ((((((a6 * w + a5) * w + a4) * w + a3) * w + a2) * w + a1) * w + one);
    sb = w * (w * (w * (w * (w * (b6 * w + b5) + b4) + b3) + b2) + b1) + one;
    sa = sa * sb * sb;

    T[25] = T[24] / 30.0; /*******************/
    sa = x * (sa + two * s) / two; /* Module 8 Result */
    sb = 0.29166666666666667; /*******************/
    sc = sa - sb;
    T[26] = one / T[25];
    /*********************/
    /*   DO NOT REMOVE   */
    /*   THIS PRINTOUT!  */
    /*********************/
    //   printf("     8   %13.4le  %10.4lf  %10.4lf\n",sc,T[24],T[26]);
    println!("     8    {:+.4e} {} {}",sc,T[2],T[4]);

    /**************************************************/
    /* MFLOPS(1) output. This is the same weighting   */
    /* used for all previous versions of the flops.c  */
    /* program. Includes Modules 2 and 3 only.        */
    /**************************************************/
    T[27] = (five * (T[6] - T[5]) + T[9]) / 52.0;
    T[28] = one / T[27];

    /**************************************************/
    /* MFLOPS(2) output. This output does not include */
    /* Module 2, but it still does 9.2% FDIV's.       */
    /**************************************************/
    T[29] = T[2] + T[9] + T[12] + T[15] + T[18];
    T[29] = (T[29] + four * T[21]) / 152.0;
    T[30] = one / T[29];

    /**************************************************/
    /* MFLOPS(3) output. This output does not include */
    /* Module 2, but it still does 3.4% FDIV's.       */
    /**************************************************/
    T[31] = T[2] + T[9] + T[12] + T[15] + T[18];
    T[31] = (T[31] + T[21] + T[24]) / 146.0;
    T[32] = one / T[31];

    /**************************************************/
    /* MFLOPS(4) output. This output does not include */
    /* Module 2, and it does NO FDIV's.               */
    /**************************************************/
    T[33] = (T[9] + T[12] + T[18] + T[24]) / 91.0;
    T[34] = one / T[33];

    //   printf("\n");
    //   printf("   Iterations      = %10ld\n",m);
    //   printf("   NullTime (usec) = %10.4lf\n",nulltime);
    //   printf("   MFLOPS(1)       = %10.4lf\n",T[28]);
    //   printf("   MFLOPS(2)       = %10.4lf\n",T[30]);
    //   printf("   MFLOPS(3)       = %10.4lf\n",T[32]);
    //   printf("   MFLOPS(4)       = %10.4lf\n\n",T[34]);

    /*********** END ***************/
}
