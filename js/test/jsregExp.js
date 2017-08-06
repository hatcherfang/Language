// match all
var re = /^\d{3}\s*\-\s*\d{3,8}$/;
var result = re.test('010   - 343536');
console.log(result);
// split match
r = 'a b  c'.split(/\s+/);
console.log(r);
r = 'a;,f;; b;  c'.split(/[\;\s\,]+/);
console.log(r);
// group match
var re = /^(\d{3})-(\d{3,8})$/;
r = re.exec('22-010-23232');
console.log(r)
// greedy match
var re = /^(\d+)(0*)$/;
r = re.exec('102300');
console.log(r);
//VS
var re = /^(\d+?)(0*)$/;
r = re.exec('102300');
console.log(r);
// global search
var s = 'javascript, JavaScript, VBScript, JScript and ECMAScript';
var re=/[a-zA-Z]+[Ss]cript/g;
r = re.exec(s);
console.log(r)
console.log(re.lastIndex)

r = re.exec(s);
console.log(r)
console.log(re.lastIndex)

r = re.exec(s);
console.log(r)
console.log(re.lastIndex)

r = re.exec(s);
console.log(r)
console.log(re.lastIndex)

