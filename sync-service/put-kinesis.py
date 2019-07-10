import logging
import boto3

def lambda_handler(event, context):
    logger = logging.getLogger()
    logger.setLevel(logging.INFO)
    try:
        logger.info(event)
        response = boto3.client('kinesis').put_record(
            StreamName = "test",
            Data = "test",
            PartitionKey = "test",
        )
        logger.info(response)
        return response

    except Exception as e:
        logger.error(e)
        raise e
