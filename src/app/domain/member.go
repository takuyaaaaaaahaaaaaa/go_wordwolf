package domain

type Member struct {
	id          int
	groupId     int
	accountName string
}

type Members []Member

func NewMember(id, groupId int, accountName string) *Member {
	return &Member{
		id:          id,
		groupId:     groupId,
		accountName: accountName,
	}
}

func (m *Member) GetID() int {
	return m.id
}

func (m *Member) GetGroupId() int {
	return m.groupId
}

func (m *Member) GetAccountName() string {
	return m.accountName
}

// インターフェイス
type MemberAction interface {
	JoinGame(id int, groupId int) error
	LeaveGame(id int, groupId int) error
	Vote(id int, groupId int, voteForId int) error
	ConfirmTheme(id int, groupId int) error
}
