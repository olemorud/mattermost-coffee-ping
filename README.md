
# Mattermost Coffee Bot

## Setup

These instructions have not been thoroughly tested, open an issue if something
does not work or is unclear.

### Setting up the Mattermost webhook

Open the Channel you want the bot integration for in Mattermost.

On the top-left there is a waffle menu icon: (looks kind of like: ☷)

Open **☷ -> Integrations -> Outgoing Webhooks -> Add Outgoing Webhook**

Add required and optional fields as you wish, but leave the Callback URLs field
empty for now. When saved, a card with information containing a _Token_ should
be displayed. This will be used later.

### Setting up Google Cloud

#### Setting up the project

1. [Set up Google Cloud ](https://cloud.google.com/docs/get-started) if you haven't already

2. [Create a Google Cloud project](https://cloud.google.com/resource-manager/docs/creating-managing-projects)

3. [Enable the Cloud Functions API](https://console.cloud.google.com/flows/enableapi?apiid=cloudfunctions,cloudbuild.googleapis.com&redirect=https://cloud.google.com/functions/quickstart&_ga=2.243466565.1565709090.1684240419-1621491083.1684240370)

4. [Check if billing is enabled for the project](https://cloud.google.com/billing/docs/how-to/verify-billing-enabled)

#### Generating credentials

Open **Google Cloud Dashboard -> APIs & Services -> Credentials**

Click on the App Engine default service account, named something like
`adjective-noun-1234@appspot.gserviceaccount.com`

Download credentials as a json file from
**Keys -> Add Key -> Create new key -> Key type: JSON**

#### Adding required secrets

CERN e-groups do not support API keys, and the bot has to authenticate with
username and password. Create a new e-groups account to mitigate exposure of
your own CERN user credentials. The account needs ownership of the e-group.
The link to the e-group is https://e-groups.cern.ch/e-groups/Egroup.do?egroupId=10542497

The password and username **should not** be stored as a secret within a GitHub
repository, but inside Google Cloud.

Open **Google Cloud -> Security -> Secret Manager** ([link](https://console.cloud.google.com/security/secret-manager)) and add secret keys named:

 - `EGROUPS_USERNAME`
 - `EGROUPS_PASSWORD`

with appropriate values.

### GitHub repository setup

#### Adding secrets

In this repository, open **Settings -> Secrets and variables -> Actions** and
add the following repository secrets:

 - `CLOUD_CREDENTIALS` The Cloud credentials downloaded in
    [Generating credentials](#generating-credentials)

 - `COLON_SEPARATED_TOKEN_WHITELIST` Token found in
    [the first step](#setting-up-the-mattermost-webhook). As the name indicates,
    several webhooks can be used at the same time.

#### Deploying bot

Open **Actions -> Deploy Google Cloud Function** and press **Run workflow**.
Deploying a function usually takes a around 5 minutes.

### Verify and final fixes

Open **Google Cloud -> Cloud Functions** [link](https://console.cloud.google.com/functions/list)

A function named `coffee` should now be visible in the list of cloud functions.
The cloud function still needs to allow all traffic. Open
**coffee -> Edit -> Runtime, build, connections and security settings -> Connections **
and check _"Allow all traffic"_ under _Ingress settings_. Egress can remain untouched.

Copy the trigger URL on the same page. Find the webhook you added in [setting up the mattermost webhook](#setting-up-the-mattermost-webhook). Edit it and add the trigger URL
to the "Callback URLs" field.

Finally test that the bot is working by typing the trigger word in the Channel
you've enabled it for. Happy coffee drinking
