### PDR PostgreSQL Database Access Example

#### System Requirements

* Python 3.7 or newer installed
* Python Libraries (See: requirements.txt)
* Google Cloud SDK installed
* Google Cloud SQL Proxy SDK component installed


#### Authentication Requirements

* DRC PMI-Ops Account
* PostgreSQL database userid and password

For detailed instructions please look at the 'Instruction - How to connect to PDR PostgreSQL' Jupyter Notebook in this directory.

#### Notebooks

##### Instruction - How to connect to PDR PostgreSQL.ipynb

This notebook shows how to connect to the PDR PostgreSQL database and query records.

##### Participant Enrollment Overview (PEO) Report.ipynb

This notebook is used to calculate the PEO report via a Jupyter Notebook.  The purpose of the report is to give a high level overview of certain metrics for all enrollment participants. This report is broken down by Awardee and will calculate the number of participants by enrollment status, UBR % of Core Participants, UBR % of Core - PM Participants, and participant counts of Gender Identity, Racial Identity, and Age.

##### Social Determinants of Health.ipynb

The SDOH notebook is used to better understand completion rates overall for SDOH and by question, as well as better understanding of demographic-related differences in completion rates.

##### COPE Survey Report.ipynb

The COPE Survey notebook will allow users to easily analyze COPE/COVID minute survey eligibility using the newly created PDR flag.
