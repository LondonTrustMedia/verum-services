# Rough Design Ideas

So these are some rough thoughts on the design of various features, as I go through this.


## Bots

These are the separate 'bots' we plan on having:

* **NickServ**: Nickname and account management.
* **ChanServ**: Channel management.
* **IdServ**: Identity management.
* **OperServ**: Services management.
* **MemoServ**: Sending and storing memos between users.


## Identity management

1. Users can easily create an account.
2. Users can improve the verifiability of their account by adding GPG key and Bitcoin details.
3. To add these details, users request for them to be added, and then must complete a given test (i.e. provide a hash to prove they own that Bitcoin address, or send an email encrypted / authenticated with GPG to a given verification address with their nickname).
4. Show verified details in /WHOIS, and note that users should be /WHOISing other users often.

We could also tie the WHOIS display into the various Bitcoin/GPG trust indexes/sites. So for instance, show their *coin/GPG identity as well as how trusted that identity is. Essentially, to stop people from just verifying to get 'the tick' next to their name and then impersonating people.
