#include "/Users/sgg7269/Development/go/src/github.com/scottshotgg/express/lib/defer.cpp"
#include "/Users/sgg7269/Development/go/src/github.com/scottshotgg/express/lib/file.cpp"
#include "/Users/sgg7269/Development/go/src/github.com/scottshotgg/express/lib/std.cpp"
#include "/Users/sgg7269/Development/go/src/github.com/scottshotgg/express/lib/var.cpp"
#include <string>
defer onExitFuncs;
std::map<std::string, var> structMap;
var genStruct(std::string structName) {
  var structValue = structMap[structName];
  return structValue;
}

int main() {

  var another = {};
  another["woah"] = 0;

  var yeah = {};
  yeah["woah"] = 66;

  var Thing = {};
  Thing["fieldA"] = 0;
  Thing["stringField"] = "";
  Thing["false_field"] = false;
  Thing["anotherFielderino"] = 0.000000;
  var thing_vqLFlSnmfB = {};
  thing_vqLFlSnmfB["woah"] = 0;
  Thing["thing"] = thing_vqLFlSnmfB;

  var something = {};
  something["fieldA"] = 0;
  something["stringField"] = "";
  something["false_field"] = false;
  something["anotherFielderino"] = 0.000000;
  var thing_WZIfYsYMHq = {};
  thing_WZIfYsYMHq["woah"] = 0;
  something["thing"] = thing_WZIfYsYMHq;

  var yeah2 = {};
  yeah2["woah"] = 8;

  var we = {};

  var yeah3 = {};
  yeah3["woah"] = 8;
}
