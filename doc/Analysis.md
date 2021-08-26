# Goal to achive
Main goal of __golock__ project is to create service/utility which can be introduced into any system as lightweight as possible and allow to safely store keys passwords connection strings and all other not intended to share data.

# Objectives to fulfil
Multiple access layers should be provided.
Every client will have its own __locker__ which they can use as they want but there needs to be __key__ which they own and only they can use it. Even if someone captures key there is additional level of protection based on who is using this __key__. Therefor any client needs to have identity.
Actors involved:
* Administrators 
* Clients
* Lockers
* Keys
## Administrators
 __Administrator__ can add __key__ to __locker__ and give it to __client__
 
 __Administrator__ cannot open __locker__ so only operations on __locker__ allowed by __administrator__ are:
  * create new __locker__
  * create __key__ for __locker__ for __client__
  * invalidate __key__ for __locker__
  * remove __locker__
  ## Client
  __Client__ is end user for __locker__. There might be:
  * multiple __clients__ per __locker__ 
  * multiple __lockers__ per __client__

  __Client__ asks __Administrator__ for :
  * new __locker__
  * access to existing __locker__ : TODO what is procedure for access to existing - human acceptance ?

  Client sends to __Administrator__ its  and recivies :
  * __locker__ address
  * __key__

  ## Locker
  __Locker__ is a repository of Key Values in encrypted form. Decryption can be made using:
  * __client__ identity
  * __key__


