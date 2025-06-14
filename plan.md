# Ideas for commands

## **Daily Standup Mode**
- `jira standup` - Shows each team member's active issues with status, blockers, and last updated
- `jira team [username]` - Deep dive into one person's full workload and overdue items
- `jira blockers` - Surfaces all issues marked as blocked across the team
- Quick assignment: `jira assign PROJ-123 @john` for missed tasks you spot during standup

## **Sprint Planning Assistant**
- `jira epics --ready` - Lists all epics in "Ready for Development" with story point estimates
- `jira epic EPIC-456 --breakdown` - Shows existing tasks under an epic, identifies gaps
- `jira create-tasks EPIC-456` - Interactive task creation wizard within an epic
- `jira plan-sprint` - Interactive sprint planning that shows team capacity vs epic priorities
- `jira balance` - Shows current sprint point distribution across team members

## **Smart Search & Memory**
- `jira find "user authentication"` - Fuzzy search across epic/task titles and descriptions
- `jira recent` - Your recently viewed/created items
- `jira bookmark PROJ-789` with `jira bookmarks` to save important epics you reference often
- Auto-complete for epic names when creating tasks: `jira create --epic "auth"` suggests "User Authentication Epic"

## **Sprint Context Awareness**
- Always show current sprint info in commands
- `jira capacity` - Team velocity vs current sprint commitment
- Flag when epics are at risk based on remaining points vs sprint time

