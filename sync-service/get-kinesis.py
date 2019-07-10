import logging
import base64

def lambda_handler(event, context):

    logger = logging.getLogger()
    logger.setLevel(logging.INFO)

    try:

        logger.info(event)
        payloads = []
        for record in event['Records']:
            payload = base64.b64decode(record["kinesis"]["data"])
            payloads.append(payload)
        logger.info(payloads)
        return payloads

    except Exception as e:
        logger.error(e)
        raise e
