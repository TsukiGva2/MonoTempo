<?php
require 'vendor/autoload.php';

use Ratchet\MessageComponentInterface;
use Ratchet\ConnectionInterface;
use PhpAmqpLib\Connection\AMQPStreamConnection;
use PhpAmqpLib\Message\AMQPMessage;

class WebSocketServer implements MessageComponentInterface {
    protected $clients;
    protected $channel;

    public function __construct() {
        $this->clients = new \SplObjectStorage;

        // Set up RabbitMQ connection
        $connection = new AMQPStreamConnection('localhost', 5672, 'guest', 'guest');
        $this->channel = $connection->channel();
        $this->channel->queue_declare('api.data', false, false, false, false);
        $this->channel->basic_consume('api.data', '', false, true, false, false, [$this, 'processMessage']);
    }

    public function onOpen(ConnectionInterface $conn) {
        $this->clients->attach($conn);
    }

    public function onMessage(ConnectionInterface $from, $msg) {
        // Handle incoming messages if needed
    }

    public function onClose(ConnectionInterface $conn) {
        $this->clients->detach($conn);
    }

    public function onError(ConnectionInterface $conn, \Exception $e) {
        $conn->close();
    }

    public function processMessage(AMQPMessage $msg) {
        $data = $msg->body;
        foreach ($this->clients as $client) {
            $client->send($data);
        }
    }
}

use Ratchet\Server\IoServer;
use Ratchet\Http\HttpServer;
use Ratchet\WebSocket\WsServer;

$server = IoServer::factory(
    new HttpServer(
        new WsServer(
            new WebSocketServer()
        )
    ),
    8080
);

$server->run();

