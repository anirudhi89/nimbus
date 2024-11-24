Someone registers, signs up, selects billing, etc. on Monitr.app

Somehow, a new website is created, hosted on ClientCompanyname.Monitr.app

Client connects their AWS acct to Monitr (input ARN?)

A dashboard is created for each selected service (for each EC2 instance, for each EKS cluster, etc)


Features:
bitemporal data for monitoring (at this point in time, how many users were using service X, etc), report generation, alerting (link to PagerDuty)

In the future: Okta integration