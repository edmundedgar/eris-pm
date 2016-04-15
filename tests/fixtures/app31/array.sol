contract c {
	struct MyStruct {
		uint myInt;
		string myString;
		bytes32 myBytes;
	}
	uint[] maIntz;
	string[] maStringz;
	bytes32[] maBytez;
	MyStruct[] maStructz;
	
	function intStorageArray(uint b) constant returns (uint) {
        maIntz.push(1);
        maIntz.push(b);
        maIntz.push(3);
        return maIntz[1];
	}

	function intMemoryArray() constant returns (uint8[4]) {
		return [1, 2, 3, 4];
	}

	function intPushArray() constant returns (uint){
        return maIntz.push(1);
	}

	function intDeleteArrray() {
        delete maIntz;
	}

	function stringStorageArray() constant returns (string) {
        string[5] memory a = stringMemoryArray();
        maStringz = a;
        return maStringz[2];
	}

	function stringMemoryArray() internal returns (string[5]) {
		return ["hello", "marmots", "how", "are", "you?"];
	}

	function stringPushArray() returns (uint) {
        maStringz.push("foo");
        maStringz.push("bar");
        return maStringz.push("diddles");
	}

	function stringDeleteArrray() {
        delete maStringz;
	}

	function bytesStorageArray() constant returns (bytes32) {
        bytes32[5] memory a = bytesMemoryArray();
	    for (var i = 0; i < 3; i++)
            maBytez.push(a[i]);
        return maBytez[2];

	}

	function bytesMemoryArray() internal constant returns (bytes32[5]){
		bytes32[5] memory b;
		b[0] = "hello";
		b[1] = "marmots";
		b[2] = "how";
		b[3] = "are";
		b[4] = "you";
		return b;
	}

	function bytesPushArray() returns (uint) {
        return maBytez.push("marmot");
	}

	function bytesDeleteArrray() {
        delete maBytez;
	}

	function structStorageArray() constant returns (uint) {
        MyStruct[3] memory a = structMemoryArray();
        for (var i = 0; i < 3; i++)
            maStructz.push(a[i]);
        return maStructz[2].myInt;
	}

	function structMemoryArray() internal constant returns (MyStruct[3]){
        return [MyStruct(1, "this", "is"), MyStruct(2, "a", "really"), MyStruct(3, "long", "test")];
	}

	function structPushArray() returns (uint) {
        return maStructz.push(MyStruct(5, "bring", "it"));
	}

	function structDeleteArrray() {
        delete maStructz;
	}