* Search YTD

XX Connect to the Database
XX Create Item Table and User Table
XX Create cron job to push the Items to the UXM API
XX uxmpusher sendItems Fetch only Items which has sent=false only
XX Save Attributes to UXM API (So create func in uxmgo to save Attributes)

XX Create a CreationWatch Pkg
XX CreationWatch - Create command to pull the feeds and push to the Item Table by Normalizing
XX Later - CreationWatch Add Logger to save pull request and response for debugging
XX Finalize the creationwatch pull with Logger and everything

XX Add Logger to the UxmPusher to UXM API request and response for debugging
* Send regular emails if there are any errors


## Deployment
* Add crons to the cronjob

## Ask John

## Consider
* How to delete the items from the UXM or how do we know that e.g. creationwatch stopped selling particular item
* How to update in stock or how to know that if particular item went out of stock