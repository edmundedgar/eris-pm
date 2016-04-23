contract c {
	uint[] maUIntz;
	int[] maIntz;
	bytes32[] maBytez;
	address[] maAddresses;
	
	function intStorageArray(uint b) constant returns (uint) {
        maUIntz.push(1);
        maUIntz.push(b);
        maUIntz.push(3);
        return maUIntz[1];
	}

	function intMemoryArray() constant returns (uint8[4]) {
		return [1, 2, 3, 4];
	}

	function intPushArray() constant returns (uint){
        return maUIntz.push(1);
	}

	function intDeleteArrray() {
        delete maUIntz;
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
}
