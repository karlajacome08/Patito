programa fib_ret;
vars
  n : entero;
  r : entero;

entero fibIter(k : entero) {
  vars
    i : entero;
    a : entero;
    b : entero;
    temp : entero;
  {
    a = 0;
    b = 1;
    i = 0;

    mientras (i < k) haz {
      temp = a + b;
      a = b;
      b = temp;
      i = i + 1;
    };

    return(a);
  }
};

inicio {
  n = 12;
  r = fibIter(n);
  escribe(r);
}
fin
