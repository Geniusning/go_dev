


function CArray(){
    this.dataStore = [];
    this.pos = 0;
    this.numElements = numElements;
    this.insert = insert;
    this.toString = toString;
    this.clear = clear;
    this.setData = setData;
    this.swap = swap;
    this.gaps = [5,3,1];
    this.shellSort =shellSort;
    this.insertSort = insertSort;
    for (let i = 0; i < numElements; i++) {
       this.dataStore[i] = i     
    }
}

function setData() {
    for (let i = 0; i < this.numElements; i++) {
        this.dataStore[i] = Math.floor(Math.random()*(this.numElements+1))     
    }
}
function clear(){
    for (let index = 0; index < this.dataStore.length; index++) {
        this.dataStore[index] = 0
    }
}
function insert(element) {
    this.dataStore[this.pos++] = element
}
function toString() {
    var restr = "";
    for (let index = 0; index < this.dataStore.length; index++) {
        restr += this.dataStore[index]+" "
        if(index>0&index%10==0){
            restr += "\n"
        }
    }
    return restr 
}
function swap(arr,index1,index2) {
    let temp = arr[index1];
    arr[index1] = arr[index2];
    arr[index2] = temp;
}

function insertSort() { //插入排序
    var temp,inner;
    for (let outer = 1; outer < this.dataStore.length; outer++) {
        temp = this.dataStore[outer];
        inner = outer;
        while (inner>0 && this.dataStore[inner-1]>=temp) {
            this.dataStore[inner] = this.dataStore[inner-1]
            --inner
        }
        this.dataStore[inner] = temp
    }
}

function shellSort() {
    for(var g= 0;g<this.gaps.length;g++){
        for(var i = this.gaps[g];i< this.dataStore.length;i++){
            var temp = this.dataStore[i]
            for(var j = i;j>=this.gaps[g] && this.dataStore[j-this.gaps[g]]>temp;j-=this.gaps[g]){
                this.dataStore[j] = this.dataStore[j-this.gaps[g]]
            }
            this.dataStore[j] = temp;
        }
    }
}//希尔排序


function quickSort(list) {
    if(list.length<1){
        return list
    }
    var less = [];
    var greater = [];
    var baseValue = list[0];
    for (let i = 0; i < list.length; i++) {
        let element = list[i];
        if (element<baseValue){
            less.push(element)
        }else {
            greater.push(element)
        }
    }
    return quickSort(less).concat([baseValue],quickSort(greater))
}

function qsort(list) {
    if(list.length==0){
        return []
    }
    var left = [];
    var right = [];
    var pivot = list[0];
    for (let i = 0; i < list.length; i++) {
        const element = list[i];
        if(element<pivot){
            left.push(element)
        }else {
            right.push(element)
        }
    }
    return qsort(left).concat(pivot,qsort(right))
}


