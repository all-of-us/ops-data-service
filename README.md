### Ops data api
The ops data api allows data stewards to use the `/participantsummary` endpoint to fetch information about their respective awardees.
This project is designed to facilitate easier use of the api, including management and authentication to Google cloud sdk.

### Prerequisites
* A working Python installation that is >= version 3
	* Documentation on installing Python can be found at [python.org](https://www.python.org/downloads/) 

* The google cloud sdk
	* Documentation on installing the google cloud sdk can be found at [google cloud sdk](https://cloud.google.com/sdk/install) 	

* Git
	* You likely already have this, but docs can be found at [git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)

### Documentation
* Specific documentation on the ops-data api is maintained at [rdr ops-data api](https://github.com/all-of-us/raw-data-repository/blob/master/opsdataAPI.md)
* The documentation enumerates uses and available parameters
* You can edit the cells (particularly cells 7 and 8) to adjust paramters and output that meet your needs.


### Usage (command line)
* Clone this repo into your home directory using `git clone <url>`
    You can grab the url from the git clone button


<img src="assets/git_clone.png" width=40%>

This will create a directory called ops-data-service
* change directory into ops-data-service
     `cd ops-data-service` 

* Run the executable named after your system/architecture
	* darwin = MacOS

Running the script will confirm you have the required prerequisites and setup a python virtual environment

###Caveat
The script will attempt to run the Jupyter notebook for you in your browser. This is however, not always possible.
If the executable fails for any reason you can still run the notebook yourself. 
This particularly happens on Windows where permissions and required libraries can be persnickety.
In that case:
The script will attempt to exit with instructions, also listed here with examples

Activate the virtual environment (or create one if it doesnt exist)

on Unix systems:

	`source venv/bin/activate`
	
on Windows systems:

	`\venv\Scripts\activate`

Install requirements
	`pip install -r requirements.txt`

Start the Jupyter notebook

	`jupyter notebook ops_data_api.ipynb`

This is an interactive ipython kernel already setup to do the hard work.
The notebook itself is documented explaining what each cell does.
The first cell contains variables you will need to setup such as `awardee` & `service_account`.

Note that it uses the stable environment by default.
You can save the file as you would any other and not need to fill out this info again.

The notebook creates a keyfile to authenticate the service account to the api.
It also removes the key when it is done running due to Googles enforced ten key limit.

If you have not setup the gcloud tool (google cloud sdk) yet, you should do that before running the notebook.
You can start by typing `gcloud init` in a terminal after it has been installed as described above.

When you are done with the notebook you can type `control c` where jupyter is running to close it.

You can type `deactivate` to leave the virtual environment
