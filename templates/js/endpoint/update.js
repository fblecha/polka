console.log('Loading {{ .Name }} update endpoint function');

var doc = require('dynamodb-doc');
var dynamo = new doc.DynamoDB();

exports.handler = function(event, context) {
    //console.log('Received event:', JSON.stringify(event, null, 2));

    var operation = event.operation;
    delete event.operation;

    switch (operation) {
        case 'update':
            console.log("{{ .Name }} update called")
            //dynamo.updateItem(event, context.done);
            break;
        default:
          {{template "endpoint_fail.js" . }}
    }
};
