let name = "Ogunewu Ezekiel Anuoluwapo";
let courses = ["HTML", "CSS", "JavaScript", "Python", "PHP", "C#", "Go"];
let display = "Name:\n" + name + "\n\nCourses:";
var num = 0;

var test = courses.length % 2;

for (var i = courses.length - 1; i >= 0; i--) {
	display += "\n"+ courses[i];
}

if (courses.length % 2 == 0) {
	display += "\n\nEven Numbers: \n 2";
	for (var i = 3; i <= 200; i++) {
		if (i % 2 == 0) {
			display += "," + i;
		}
	}
} else {
	display += "\n\nOdd Numbers: \n 1";
	for (var i = 2; i <= 200; i++) {
		if (i % 2 == 1) {
			display += "," + i;
		}
	}
}

console.log(display);
alert(display);