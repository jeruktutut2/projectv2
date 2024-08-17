# Order

## python
install pyenv: brew install pyenv  
install python: pyenv install 3.12.4  
set python: pyenv global 3.12.4  
set pyenv python: eval "$(pyenv init -)"  
install mysql connector: pip install mysql-connector-python  
https://www.youtube.com/watch?v=rWC2iFlN3TM  
install virtualenv: 
https://www.youtube.com/watch?v=rWC2iFlN3TM  
https://packaging.python.org/en/latest/guides/installing-using-pip-and-virtual-environments/  
python3 -m venv .venv  
source .venv/bin/activate  
which python  
python3 -m pip install --upgrade pip  
python3 -m pip --version  

## install mysql connector
pip install mysql-connector-python  

## install pydantic
eval "$(pyenv init -)"  
source .venv/bin/activate  
pip install pydantic  

## install pytest
eval "$(pyenv init -)"  
source .venv/bin/activate  
pip install -U pytest  

## add pythonpath
export PYTHONPATH=$(pwd):$PYTHONPATH

## run test pytest
eval "$(pyenv init -)"
source .venv/bin/activate
export PYTHONPATH=$(pwd):$PYTHONPATH
<!-- pytest tests/features/create_order/services/create_order_test.py -->
pytest -s tests/integration_tests/features/create_order/services/test_create_order.py
pytest -s tests/integration_tests/features/get_order/services/test_get_order.py
pytest -s tests/integration_tests/features/update_order/services/test_update_order.py
pytest -s tests/integration_tests/features/delete_order/services/test_delete_order.py

<!-- failed doing unit test when wrapping method in mysqlutil and test it in test_success -->
pytest -s tests/unit_tests/features/create_order/services/test_create_order.py

pytest -s tests/api_tests/features/create_order/test_create_order.py
pytest -s tests/api_tests/features/get_order/test_get_order.py
pytest -s tests/api_tests/features/update_order/test_update_order.py
pytest -s tests/api_tests/features/delete_order/test_delete_order.py

## instal flask
eval "$(pyenv init -)"
source .venv/bin/activate
export PYTHONPATH=$(pwd):$PYTHONPATH
pip install Flask

## run application
eval "$(pyenv init -)"
source .venv/bin/activatepython app.py
export PYTHONPATH=$(pwd):$PYTHONPATH
python main.py

## run curl test
chmod +x test_create_order.sh
./test_create_order.sh

chmod +x test_get_order.sh
./test_get_order.sh

chmod +x test_update_order.sh
./test_update_order.sh

chmod +x test_delete_order.sh
./test_delete_order.sh