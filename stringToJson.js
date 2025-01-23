function stringToJson(str) {
  str = JSON.parse(str);

  console.log(str);
}

stringToJson('{"op": "skip", "count": 40}');
