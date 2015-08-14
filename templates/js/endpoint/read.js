console.log('Loading function - {{ .Name }} read');

var doc = require('dynamodb-doc');
var dynamo = new doc.DynamoDB();

exports.handler = function(event, context) {
    //console.log('Received event:', JSON.stringify(event, null, 2));

    var operation = event.operation;
    delete event.operation;

    switch (operation) {
        case 'read':
            dynamo.getItem(event, context.done);
            break;
        default:
          {{template "endpoint_fail.js" . }}
    }
};
