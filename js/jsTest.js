// JSON test
var dict1 = {custom: "fang", default: "haiqun"}
console.log(dict1)
jstr = JSON.stringify(dict1)
jpar = JSON.parse(jstr)
console.log(jstr)
console.log(jpar)
// constructor
function Student(name){
    this.name = name;
    this.hello = function () {
    	console.log('hello ' + this.name + '!');
    }
}
var xiaoming = new Student('xiaoming');
xiaoming.hello()
// operator || test
var s = 1!==1 || 'fang'
console.log(s)
