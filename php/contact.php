<?php

$first_name = $last_name = $email = $message = "";

if (isset($_POST['email'])) {
	# code...
	$first_name = Test_Input($_POST['first_name']);
	$last_name = Test_Input($_POST['last_name']);
	$email = Test_Input($_POST['email']);
	$message = Test_Input($_POST['message']);

	$fname = $first_name . "_" . $last_name . ".txt";

	$file = fopen($fname, "a") or die("Unable to open file!");
	$txt = "Name: " . $first_name . " " . $last_name;
	fwrite($file, $txt);

	$txt = "Email: " . $email;
	fwrite($file, "\n". $txt);

	$txt = "Message: " . $message;
	fwrite($file, "\n".$txt);

	fclose($file);
}

function Test_Input($data){
    $data = trim($data);
    $data = stripslashes($data);
    $data = htmlspecialchars($data);
    return $data;
}
echo <<<_END
<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8" />
	<title>Contact Me</title>
	<link rel="stylesheet" type="text/css" href="style.css" />
	<meta name="description" content="First PHP Task for start.ng" />
</head>
<body>

	<div>
		<h2>Contact Form</h2>
		<form action="contact.php" method="POST">
			<input type="text" name="first_name" placeholder="Enter your first name" /><br />
			<input type="text" name="last_name" placeholder="Enter your last name" /><br />
			<input type="email" name="email" placeholder="Enter your email" required /><br />
			<br />

			<textarea name="message"></textarea><br />

			<button type="submit">Send Message</button>
		</form>
	</div>
</body>
</html>
_END
?>