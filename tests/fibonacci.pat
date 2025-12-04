programa fib;
vars
  n : entero;

nula fibSerie(limite : entero) {
  vars
    a : entero;
    b : entero;
    i : entero;
    temp : entero;
  {
    a = 0;
    b = 1;
    i = 0;

    mientras (i < limite) haz {
      escribe(a);
      temp = a + b;
      a = b;
      b = temp;
      i = i + 1;
    };
  }
};

inicio {
  n = 6;           
  fibSerie(n);     
}
fin
