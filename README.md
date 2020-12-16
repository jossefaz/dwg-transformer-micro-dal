---


---

<h1 id="dwg-transformer--dal">Dwg-transformer : DAL</h1>
<p>This module is a part of the dwg-transformer-microservice tool. It’s called the <strong>DAL</strong> because he is responsible for accessing the database (in our hybrid architecture for the moment there is only one DB because other microservices do not require any storing, they just need metadata to indicates which files they have to convert/deal with)</p>
<h1 id="configuration">Configuration</h1>
<p>Some environment variables must be set before running this program :</p>
<p><strong>$DEV_PROD</strong> :  define the environment. Must be set to one of those values : “prod” / “dev”, depending on the configuration files that you rely on.</p>
<p><strong>$DB</strong>  : define the which database to look for the data (currently the only db configured in the DAL service is “mysql but it is extremly simple to add more)</p>
<p><strong>$SCHEMA</strong>  : define the schema’s name in which the DAL will look up for data (“dwg_transformer” is the schema defined in the init.sql of the DAL service for the current version)</p>
<p><strong>$CAD_STATUS</strong>  : define the table name which stores the dwg files metadata to check</p>
<p><strong>$CAD_ERRORS</strong>  : define the table name which stores different errors that occured on dwg files checking</p>
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
<p><strong>Queue</strong> : define which message broker will be used and how to connect to it. Here is an example with RabbitMQ connection string to localhost.<br>
—&gt; <strong>ConnString</strong> : the connection string to the queue<br>
—&gt; <strong>QueueNames</strong> : the queues that this service will open ( if not exists) an publish to them<br>
—&gt; <strong>Listennig</strong> : the queues that this service will subscribe and listen for messages</p>
<p><strong>Logs</strong> : define differents logger foile for the service log. Currently, the service is using <a href="https://github.com/sirupsen/logrus"><strong>logrus</strong></a> library but it could change in future versions.<br>
—&gt; <strong>Path</strong> : define the log file location<br>
—&gt; <strong>Level</strong> : Log level for this file</p>
<h2 id="dependencies">Dependencies</h2>
<p>This service, as the orchestrator, is dependent on at least 4 others services to run correctly, an do the job.</p>
<ol>
<li><strong>DAL</strong> service : implement access to the dwg database.</li>
<li><strong>transformer</strong> service : responsible to convert the DWG files to DXF or others format</li>
<li><strong>worker</strong> service : will perform validity check on the converted file</li>
<li><strong>Message Broker</strong> service  : responsible for the communication between microservices</li>
</ol>
<h2 id="flow">Flow</h2>
<p>The controller workflow is very simple :</p>
<div class="mermaid"><svg xmlns="http://www.w3.org/2000/svg" id="mermaid-svg-J2VvU4TwwuOFrJ7r" height="100%" width="100%" style="max-width:750px;" viewBox="-150 -10 750 415"><g></g><g><line id="actor3" x1="75" y1="5" x2="75" y2="404" class="actor-line" stroke-width="0.5px" stroke="#999"></line><rect x="0" y="0" fill="#eaeaea" stroke="#666" width="150" height="65" rx="3" ry="3" class="actor"></rect><text x="75" y="32.5" dominant-baseline="central" alignment-baseline="central" class="actor" style="text-anchor: middle;"><tspan x="75" dy="0">Controller</tspan></text></g><g><line id="actor4" x1="275" y1="5" x2="275" y2="404" class="actor-line" stroke-width="0.5px" stroke="#999"></line><rect x="200" y="0" fill="#eaeaea" stroke="#666" width="150" height="65" rx="3" ry="3" class="actor"></rect><text x="275" y="32.5" dominant-baseline="central" alignment-baseline="central" class="actor" style="text-anchor: middle;"><tspan x="275" dy="0">Message Broker</tspan></text></g><g><line id="actor5" x1="475" y1="5" x2="475" y2="404" class="actor-line" stroke-width="0.5px" stroke="#999"></line><rect x="400" y="0" fill="#eaeaea" stroke="#666" width="150" height="65" rx="3" ry="3" class="actor"></rect><text x="475" y="32.5" dominant-baseline="central" alignment-baseline="central" class="actor" style="text-anchor: middle;"><tspan x="475" dy="0">DAL</tspan></text></g><defs><marker id="arrowhead" refX="5" refY="2" markerWidth="6" markerHeight="4" orient="auto"><path d="M 0,0 V 4 L6,2 Z"></path></marker></defs><defs><marker id="crosshead" markerWidth="15" markerHeight="8" orient="auto" refX="16" refY="4"><path fill="black" stroke="#000000" stroke-width="1px" d="M 9,2 V 6 L16,4 Z" style="stroke-dasharray: 0, 0;"></path><path fill="none" stroke="#000000" stroke-width="1px" d="M 0,1 L 6,7 M 6,1 L 0,7" style="stroke-dasharray: 0, 0;"></path></marker></defs><g><text x="175" y="93" class="messageText" style="text-anchor: middle;">DAL request</text><line x1="75" y1="100" x2="275" y2="100" class="messageLine0" stroke-width="2" stroke="black" marker-end="url(#arrowhead)" style="fill: none;"></line></g><g><text x="375" y="128" class="messageText" style="text-anchor: middle;">Is there new data ?</text><line x1="275" y1="135" x2="475" y2="135" class="messageLine1" stroke-width="2" stroke="black" marker-end="url(#arrowhead)" style="stroke-dasharray: 3, 3; fill: none;"></line></g><g><text x="375" y="163" class="messageText" style="text-anchor: middle;">Data payload (if any)</text><line x1="475" y1="170" x2="275" y2="170" class="messageLine1" stroke-width="2" stroke="black" marker-end="url(#arrowhead)" style="stroke-dasharray: 3, 3; fill: none;"></line></g><g><text x="175" y="198" class="messageText" style="text-anchor: middle;">Data payload (if any)</text><line x1="275" y1="205" x2="75" y2="205" class="messageLine1" stroke-width="2" stroke="black" marker-end="url(#arrowhead)" style="stroke-dasharray: 3, 3; fill: none;"></line></g><g><rect x="-100" y="215" fill="#EDF2AE" stroke="#666" width="150" height="104" rx="0" ry="0" class="note"></rect><text x="-104" y="239" fill="black" class="noteText"><tspan x="-84" fill="black">From now the </tspan></text><text x="-104" y="253" fill="black" class="noteText"><tspan x="-84" fill="black"> controller will send </tspan></text><text x="-104" y="267" fill="black" class="noteText"><tspan x="-84" fill="black">to the transformer</tspan></text><text x="-104" y="281" fill="black" class="noteText"><tspan x="-84" fill="black">a message with </tspan></text><text x="-104" y="295" fill="black" class="noteText"><tspan x="-84" fill="black"> the path of the files </tspan></text><text x="-104" y="309" fill="black" class="noteText"><tspan x="-84" fill="black"> to convert</tspan></text></g><g><rect x="0" y="339" fill="#eaeaea" stroke="#666" width="150" height="65" rx="3" ry="3" class="actor"></rect><text x="75" y="371.5" dominant-baseline="central" alignment-baseline="central" class="actor" style="text-anchor: middle;"><tspan x="75" dy="0">Controller</tspan></text></g><g><rect x="200" y="339" fill="#eaeaea" stroke="#666" width="150" height="65" rx="3" ry="3" class="actor"></rect><text x="275" y="371.5" dominant-baseline="central" alignment-baseline="central" class="actor" style="text-anchor: middle;"><tspan x="275" dy="0">Message Broker</tspan></text></g><g><rect x="400" y="339" fill="#eaeaea" stroke="#666" width="150" height="65" rx="3" ry="3" class="actor"></rect><text x="475" y="371.5" dominant-baseline="central" alignment-baseline="central" class="actor" style="text-anchor: middle;"><tspan x="475" dy="0">DAL</tspan></text></g></svg></div>

