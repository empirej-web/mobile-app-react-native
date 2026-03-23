import logging
import re
import json
from datetime import datetime
from typing import Union, List, Dict

# Set up logging
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

def validate_email(email: str) -> bool:
    """
    Validate an email address.

    Args:
    email (str): The email address to validate.

    Returns:
    bool: True if the email is valid, False otherwise.
    """
    pattern = r"^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$"
    return bool(re.match(pattern, email))

def convert_date(date_str: str) -> datetime:
    """
    Convert a date string to a datetime object.

    Args:
    date_str (str): The date string to convert.

    Returns:
    datetime: The datetime object.
    """
    return datetime.strptime(date_str, "%Y-%m-%d")

def load_json(file_path: str) -> Dict:
    """
    Load a JSON file.

    Args:
    file_path (str): The path to the JSON file.

    Returns:
    Dict: The loaded JSON data.
    """
    with open(file_path, "r") as file:
        return json.load(file)

def save_json(data: Dict, file_path: str) -> None:
    """
    Save JSON data to a file.

    Args:
    data (Dict): The JSON data to save.
    file_path (str): The path to the file to save to.
    """
    with open(file_path, "w") as file:
        json.dump(data, file, indent=4)

def split_list(input_list: List, chunk_size: int) -> List:
    """
    Split a list into chunks of a specified size.

    Args:
    input_list (List): The list to split.
    chunk_size (int): The size of each chunk.

    Returns:
    List: A list of chunks.
    """
    return [input_list[i:i + chunk_size] for i in range(0, len(input_list), chunk_size)]

def is_valid_integer(num: Union[str, int]) -> bool:
    """
    Check if a value is a valid integer.

    Args:
    num (Union[str, int]): The value to check.

    Returns:
    bool: True if the value is a valid integer, False otherwise.
    """
    try:
        int(num)
        return True
    except ValueError:
        return False