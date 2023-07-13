# H Docs

Additional documentation which gives details about our implementation.

## Features

Explanation of the main Hydra specific features.

### Fee Distribution

We've added a burn address to send there 50% of the fees. It is an EOA address we assume no one will be able to retrieve.
We don't enable the --burn-address of the usptream implementation because it turns on the EIP1559 which is something we don't want to apply.

## General findings

### Validators update

Events for change in the balance are handled after every block and validators state is saved in the DB
At the end of the epoch the new state (lasly update at block endOfEpoch - 1) is applied.
PROBLEM: changes in the validators made in the last block of an epoch are not applied in the next epoch but in the next + 1 epoch.
So our contract would have different state of the current validators compared to the state in the node. That leds to problems when commitEpoch is executed.
After temporary fixes we've made, a problem can occur when chanigng balance at th last block of an epoch.

// TODO: Modify our implementation of the contracts
