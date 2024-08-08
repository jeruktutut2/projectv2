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