console.log('{{ .Name }}.health - Loading function');

var doc = require('dynamodb-doc');
var dynamo = new doc.DynamoDB();

exports.handler = function(event, context) {
    //console.log('Received event:', JSON.stringify(event, null, 2));
    var operation = event.operation;
    delete event.operation;

    switch (operation) {
        case 'ping':
            context.succeed('pong');
            break;
        default:
          {{template "endpoint_fail.js" . }}
    }
};
