<h1 align="center">📨Aula básica de Apache Kafka📩</h1>

<h3 align="center">✉️Aprendendo o básico de mensageria com Kafka e Golang✉️</h3>

<p>
    Projeto criado com base na <a href="https://www.youtube.com/watch?v=E2SWb4rn2dY" style="text-decoration: none;">live</a> live sobre Apache Kafka do esquenta da maratona FullStack FullCycle 3.0, idealizada pelo 
    <a href="https://github.com/wesleywillians" style="text-decoration: none;">Wesley Wilians</a>.
</p>

<h4>🛠Ferramentas e tecnologias: 🛠</h4>
<ul>
    <li>Apache Kafka</li>
    <li>Apache ZooKeeper</li>
    <li>Go</li>
    <li>librdkafka</li>
</ul>

<h4>🕹️Principais comandos: 🕹️</h4>
<ul>
    <li>
        kafka-topics --create --topic=exemplo --partitions=3 --bootstrap-server=localhost:9092
    </li>
    <li>
        kafka-console-producer --topic=exemplo --bootstrap-server=localhost:9092
    </li>
    <li>
        kafka-console-consumer --topic=exemplo --bootstrap-server=localhost:9092
    </li>
    <li>
        kafka-console-consumer --topic=exemplo --bootstrap-server=localhost:9092 --group=x
    </li>
    <li>
        kafka-consumer-groups --group=x --describe --bootstrap-server=localhost:9092
    </li>
</ul>