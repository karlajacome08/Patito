programa test2;
vars
  a : entero;
  b : entero;

nula imprimeSuma(x : entero, y : entero) {
  {
    escribe(x + y);
  }
};

inicio {
  a = 3;
  b = 5;
  imprimeSuma(a, b);
}
fin
