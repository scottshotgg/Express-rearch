#include "/Users/sgg7269/Development/go/src/github.com/scottshotgg/express/lib/defer.cpp"
#include "/Users/sgg7269/Development/go/src/github.com/scottshotgg/express/lib/file.cpp"
#include "/Users/sgg7269/Development/go/src/github.com/scottshotgg/express/lib/std.cpp"
#include "/Users/sgg7269/Development/go/src/github.com/scottshotgg/express/lib/var.cpp"
#include <string>
defer onExitFuncs;

int printStuff(int k) {
  defer onReturnFuncs;
  {
    defer onLeaveFuncs;

    {
      int i = 0;
      while (i < k) {
        {
          defer onLeaveFuncs;
          onExitFuncs.deferStack.push([=](...) { Println("on exit", i); });
          onReturnFuncs.deferStack.push([=](...) { Println("on return", i); });
          onLeaveFuncs.deferStack.push([=](...) { Println("on leave", i); });
          onReturnFuncs.deferStack.push([=](...) { Println("defer", i); });
        }
        i += 1;
      }
    }
    return 0;
  }
}

var increment(var i) {
  defer onReturnFuncs;
  {
    defer onLeaveFuncs;
    var _hVZXpHISeH = {};
    _hVZXpHISeH["something"] = "else";
    return _hVZXpHISeH;
  }
}

int main() {
  var thingy = 7;

  Print("thingy =", thingy, "\n");

  Println(thingy);

  Println();
  thingy = 69.69;

  Print("thingy =", thingy, "\n");

  Println(thingy);

  Println();
  thingy = "woah woah woah";

  Print("thingy =", thingy, "\n");

  Println(thingy);

  Println();
  thingy = false;

  Print("thingy =", thingy, "\n");

  Println(thingy);

  Println();
  var thingyObject = {};
  thingyObject["im_just_a"] = "DEAD BOY";
}