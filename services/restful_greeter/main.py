import logging
import connexion
from helpers import dotcase_resolver

if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO)
    app = connexion.App(__name__,
                        specification_dir='../../',
                        swagger_ui=True,
                        swagger_json=True # Default but left for demonstration purposes
                       )
    app.add_api('greeter.swagger.json',
                resolver=connexion.resolver.Resolver(dotcase_resolver))
    app.run()
