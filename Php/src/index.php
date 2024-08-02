<!--
	$servername = $_ENV["MYSQL_HOST"];
	$username = $_ENV["MYSQL_USER"];
	$password = $_ENV["MYSQL_PASS"];
	$dbname = $_ENV["MYSQL_DB"];

	// Create connection
	$conn = new mysqli($servername, $username, $password, $dbname);

	// Check connection
	if ($conn->connect_error) {
	    die("Connection failed: " . $conn->connect_error);
	}

	echo "Connected successfully";
-->

<!DOCTYPE html>
<html>
<head>
    <title>WebSocket Example</title>
</head>
<body>
    <div id="content">Waiting for updates...</div>
    <script>
        // Create a new WebSocket connection
        var ws = new WebSocket('ws://localhost:8080');

        // Event handler for when a message is received from the server
        ws.onmessage = function(event) {
            // Update the content div with the received message
            document.getElementById('content').innerText = event.data;
        };

        // Event handler for when the WebSocket connection is opened
        ws.onopen = function(event) {
            console.log('WebSocket connection opened');
        };

        // Event handler for when the WebSocket connection is closed
        ws.onclose = function(event) {
            console.log('WebSocket connection closed');
        };

        // Event handler for when there's an error with the WebSocket connection
        ws.onerror = function(event) {
            console.error('WebSocket error:', event);
        };
    </script>
</body>
</html>

