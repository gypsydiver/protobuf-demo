""" utility functions for our restful greeter """
import connexion
import case_conversion

def dotcase_resolver(operation_id):
    """
    Transform operation_id to dotcase before the name is sent to the default
    Resolver algorithm
    """
    name = case_conversion.dotcase(operation_id)
    return connexion.utils.get_function_from_name(name)
