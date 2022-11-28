<!--
order: 4
-->

# Events

The group module emits the following events:

## EventCreateGroup

| Type                             | Attribute Key | Attribute Value                  |
| -------------------------------- | ------------- | -------------------------------- |
| message                          | action        | /witvalchain.group.v1.Msg/CreateGroup |
| witvalchain.group.v1.EventCreateGroup | group_id      | {groupId}                        |

## EventUpdateGroup

| Type                             | Attribute Key | Attribute Value                                            |
| -------------------------------- | ------------- | ---------------------------------------------------------- |
| message                          | action        | /witvalchain.group.v1.Msg/UpdateGroup{Admin\|Metadata\|Members} |
| witvalchain.group.v1.EventUpdateGroup | group_id      | {groupId}                                                  |

## EventCreateGroupPolicy

| Type                                   | Attribute Key | Attribute Value                        |
| -------------------------------------- | ------------- | -------------------------------------- |
| message                                | action        | /witvalchain.group.v1.Msg/CreateGroupPolicy |
| witvalchain.group.v1.EventCreateGroupPolicy | address       | {groupPolicyAddress}                   |

## EventUpdateGroupPolicy

| Type                                   | Attribute Key | Attribute Value                                                         |
| -------------------------------------- | ------------- | ----------------------------------------------------------------------- |
| message                                | action        | /witvalchain.group.v1.Msg/UpdateGroupPolicy{Admin\|Metadata\|DecisionPolicy} |
| witvalchain.group.v1.EventUpdateGroupPolicy | address       | {groupPolicyAddress}                                                    |

## EventCreateProposal

| Type                                | Attribute Key | Attribute Value                     |
| ----------------------------------- | ------------- | ----------------------------------- |
| message                             | action        | /witvalchain.group.v1.Msg/CreateProposal |
| witvalchain.group.v1.EventCreateProposal | proposal_id   | {proposalId}                        |

## EventWithdrawProposal

| Type                                  | Attribute Key | Attribute Value                       |
| ------------------------------------- | ------------- | ------------------------------------- |
| message                               | action        | /witvalchain.group.v1.Msg/WithdrawProposal |
| witvalchain.group.v1.EventWithdrawProposal | proposal_id   | {proposalId}                          |

## EventVote

| Type                      | Attribute Key | Attribute Value           |
| ------------------------- | ------------- | ------------------------- |
| message                   | action        | /witvalchain.group.v1.Msg/Vote |
| witvalchain.group.v1.EventVote | proposal_id   | {proposalId}              |

## EventExec

| Type                      | Attribute Key | Attribute Value           |
| ------------------------- | ------------- | ------------------------- |
| message                   | action        | /witvalchain.group.v1.Msg/Exec |
| witvalchain.group.v1.EventExec | proposal_id   | {proposalId}              |
| witvalchain.group.v1.EventExec | logs          | {logs_string}             |

## EventLeaveGroup

| Type                            | Attribute Key | Attribute Value                 |
| ------------------------------- | ------------- | ------------------------------- |
| message                         | action        | /witvalchain.group.v1.Msg/LeaveGroup |
| witvalchain.group.v1.EventLeaveGroup | proposal_id   | {proposalId}                    |
| witvalchain.group.v1.EventLeaveGroup | address       | {address}                       |
