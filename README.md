---


---

<h1 id="dwg-transformer--dal">Dwg-transformer : DAL</h1>
<p>This module is a part of the dwg-transformer-microservice tool. It’s called the  <strong>DAL</strong>  because he is responsible for accessing the database (in our hybrid architecture for the moment there is only one DB because other microservices do not require any storing, they just need metadata to indicates which files they have to convert/deal with)</p>
<h1 id="configuration">Configuration</h1>
<p>Only two environment variables must be set before running this program :</p>
<p><strong>$DEV_PROD</strong>  : define the environment. Must be set to one of those values : “prod” / “dev”, depending on the configuration files that you rely on.</p>
<p><strong>$LOOKUP_ERRORS_SQL</strong>  : define the SQL to get the look up table of the check errors</p>
<ol>
<li>config.prod.json (for production)</li>
<li>config.dev.json (for dev)</li>
</ol>
<p>Here is an example of the configuration file.</p>
<pre><code>    {  
		  "Queue" : {  
		    "Rabbitmq" : {  
		      "ConnString" : "amqp://guest:guest@localhost:5672?connection_attempts=5&amp;retry_delay=5",  
			  "QueueNames" : [  "ConvertedDWG", "CheckedDWG", "ConvertDWG", "CheckDWG", "Dal_Res", "Dal_Req"],  
			  "Listennig" : [ "ConvertedDWG", "CheckedDWG", "Dal_Res"]  
			}  
		  },  
		  "Logs" : {  
		      "Main" : {  
		        "Path" : "controller.log",  
				"Level" : "DEBUG"  
		  }  
  }  
}

</code></pre>
<p><strong>Queue</strong>  : define which message broker will be used and how to connect to it. Here is an example with RabbitMQ connection string to localhost.<br>
—&gt;  <strong>ConnString</strong>  : the connection string to the queue<br>
—&gt;  <strong>QueueNames</strong>  : the queues that this service will open ( if not exists) an publish to them<br>
—&gt;  <strong>Listennig</strong>  : the queues that this service will subscribe and listen for messages</p>
<p><strong>Logs</strong>  : define differents logger foile for the service log. Currently, the service is using  <a href="https://github.com/sirupsen/logrus"><strong>logrus</strong></a>  library but it could change in future versions.<br>
—&gt;  <strong>Path</strong>  : define the log file location<br>
—&gt;  <strong>Level</strong>  : Log level for this file</p>
<h2 id="dependencies">Dependencies</h2>
<p>This service is dependent on at least 1 others services to run correctly, an do the job.</p>
<ol>
<li><strong>Message Broker</strong>  service : responsible for the communication between microservices</li>
</ol>

