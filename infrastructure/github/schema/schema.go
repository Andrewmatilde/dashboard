package schema

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type Query struct {
	CodeOfConduct                            *CodeOfConduct                     "json:\"codeOfConduct\" graphql:\"codeOfConduct\""
	CodesOfConduct                           []*CodeOfConduct                   "json:\"codesOfConduct\" graphql:\"codesOfConduct\""
	Enterprise                               *Enterprise                        "json:\"enterprise\" graphql:\"enterprise\""
	EnterpriseAdministratorInvitation        *EnterpriseAdministratorInvitation "json:\"enterpriseAdministratorInvitation\" graphql:\"enterpriseAdministratorInvitation\""
	EnterpriseAdministratorInvitationByToken *EnterpriseAdministratorInvitation "json:\"enterpriseAdministratorInvitationByToken\" graphql:\"enterpriseAdministratorInvitationByToken\""
	License                                  *License                           "json:\"license\" graphql:\"license\""
	Licenses                                 []*License                         "json:\"licenses\" graphql:\"licenses\""
	MarketplaceCategories                    []*MarketplaceCategory             "json:\"marketplaceCategories\" graphql:\"marketplaceCategories\""
	MarketplaceCategory                      *MarketplaceCategory               "json:\"marketplaceCategory\" graphql:\"marketplaceCategory\""
	MarketplaceListing                       *MarketplaceListing                "json:\"marketplaceListing\" graphql:\"marketplaceListing\""
	MarketplaceListings                      MarketplaceListingConnection       "json:\"marketplaceListings\" graphql:\"marketplaceListings\""
	Meta                                     GitHubMetadata                     "json:\"meta\" graphql:\"meta\""
	Node                                     Node                               "json:\"node\" graphql:\"node\""
	Nodes                                    []Node                             "json:\"nodes\" graphql:\"nodes\""
	Organization                             *Organization                      "json:\"organization\" graphql:\"organization\""
	RateLimit                                *RateLimit                         "json:\"rateLimit\" graphql:\"rateLimit\""
	Relay                                    *Query                             "json:\"relay\" graphql:\"relay\""
	Repository                               *Repository                        "json:\"repository\" graphql:\"repository\""
	RepositoryOwner                          RepositoryOwner                    "json:\"repositoryOwner\" graphql:\"repositoryOwner\""
	Resource                                 UniformResourceLocatable           "json:\"resource\" graphql:\"resource\""
	Search                                   SearchResultItemConnection         "json:\"search\" graphql:\"search\""
	SecurityAdvisories                       SecurityAdvisoryConnection         "json:\"securityAdvisories\" graphql:\"securityAdvisories\""
	SecurityAdvisory                         *SecurityAdvisory                  "json:\"securityAdvisory\" graphql:\"securityAdvisory\""
	SecurityVulnerabilities                  SecurityVulnerabilityConnection    "json:\"securityVulnerabilities\" graphql:\"securityVulnerabilities\""
	SponsorsListing                          *SponsorsListing                   "json:\"sponsorsListing\" graphql:\"sponsorsListing\""
	Topic                                    *Topic                             "json:\"topic\" graphql:\"topic\""
	User                                     *User                              "json:\"user\" graphql:\"user\""
	Viewer                                   User                               "json:\"viewer\" graphql:\"viewer\""
}

type Mutation struct {
	AcceptEnterpriseAdministratorInvitation                     *AcceptEnterpriseAdministratorInvitationPayload                     "json:\"acceptEnterpriseAdministratorInvitation\" graphql:\"acceptEnterpriseAdministratorInvitation\""
	AcceptTopicSuggestion                                       *AcceptTopicSuggestionPayload                                       "json:\"acceptTopicSuggestion\" graphql:\"acceptTopicSuggestion\""
	AddAssigneesToAssignable                                    *AddAssigneesToAssignablePayload                                    "json:\"addAssigneesToAssignable\" graphql:\"addAssigneesToAssignable\""
	AddComment                                                  *AddCommentPayload                                                  "json:\"addComment\" graphql:\"addComment\""
	AddLabelsToLabelable                                        *AddLabelsToLabelablePayload                                        "json:\"addLabelsToLabelable\" graphql:\"addLabelsToLabelable\""
	AddProjectCard                                              *AddProjectCardPayload                                              "json:\"addProjectCard\" graphql:\"addProjectCard\""
	AddProjectColumn                                            *AddProjectColumnPayload                                            "json:\"addProjectColumn\" graphql:\"addProjectColumn\""
	AddPullRequestReview                                        *AddPullRequestReviewPayload                                        "json:\"addPullRequestReview\" graphql:\"addPullRequestReview\""
	AddPullRequestReviewComment                                 *AddPullRequestReviewCommentPayload                                 "json:\"addPullRequestReviewComment\" graphql:\"addPullRequestReviewComment\""
	AddPullRequestReviewThread                                  *AddPullRequestReviewThreadPayload                                  "json:\"addPullRequestReviewThread\" graphql:\"addPullRequestReviewThread\""
	AddReaction                                                 *AddReactionPayload                                                 "json:\"addReaction\" graphql:\"addReaction\""
	AddStar                                                     *AddStarPayload                                                     "json:\"addStar\" graphql:\"addStar\""
	ArchiveRepository                                           *ArchiveRepositoryPayload                                           "json:\"archiveRepository\" graphql:\"archiveRepository\""
	CancelEnterpriseAdminInvitation                             *CancelEnterpriseAdminInvitationPayload                             "json:\"cancelEnterpriseAdminInvitation\" graphql:\"cancelEnterpriseAdminInvitation\""
	ChangeUserStatus                                            *ChangeUserStatusPayload                                            "json:\"changeUserStatus\" graphql:\"changeUserStatus\""
	ClearLabelsFromLabelable                                    *ClearLabelsFromLabelablePayload                                    "json:\"clearLabelsFromLabelable\" graphql:\"clearLabelsFromLabelable\""
	CloneProject                                                *CloneProjectPayload                                                "json:\"cloneProject\" graphql:\"cloneProject\""
	CloneTemplateRepository                                     *CloneTemplateRepositoryPayload                                     "json:\"cloneTemplateRepository\" graphql:\"cloneTemplateRepository\""
	CloseIssue                                                  *CloseIssuePayload                                                  "json:\"closeIssue\" graphql:\"closeIssue\""
	ClosePullRequest                                            *ClosePullRequestPayload                                            "json:\"closePullRequest\" graphql:\"closePullRequest\""
	ConvertProjectCardNoteToIssue                               *ConvertProjectCardNoteToIssuePayload                               "json:\"convertProjectCardNoteToIssue\" graphql:\"convertProjectCardNoteToIssue\""
	CreateBranchProtectionRule                                  *CreateBranchProtectionRulePayload                                  "json:\"createBranchProtectionRule\" graphql:\"createBranchProtectionRule\""
	CreateCheckRun                                              *CreateCheckRunPayload                                              "json:\"createCheckRun\" graphql:\"createCheckRun\""
	CreateCheckSuite                                            *CreateCheckSuitePayload                                            "json:\"createCheckSuite\" graphql:\"createCheckSuite\""
	CreateEnterpriseOrganization                                *CreateEnterpriseOrganizationPayload                                "json:\"createEnterpriseOrganization\" graphql:\"createEnterpriseOrganization\""
	CreateIPAllowListEntry                                      *CreateIPAllowListEntryPayload                                      "json:\"createIpAllowListEntry\" graphql:\"createIpAllowListEntry\""
	CreateIssue                                                 *CreateIssuePayload                                                 "json:\"createIssue\" graphql:\"createIssue\""
	CreateProject                                               *CreateProjectPayload                                               "json:\"createProject\" graphql:\"createProject\""
	CreatePullRequest                                           *CreatePullRequestPayload                                           "json:\"createPullRequest\" graphql:\"createPullRequest\""
	CreateRef                                                   *CreateRefPayload                                                   "json:\"createRef\" graphql:\"createRef\""
	CreateRepository                                            *CreateRepositoryPayload                                            "json:\"createRepository\" graphql:\"createRepository\""
	CreateTeamDiscussion                                        *CreateTeamDiscussionPayload                                        "json:\"createTeamDiscussion\" graphql:\"createTeamDiscussion\""
	CreateTeamDiscussionComment                                 *CreateTeamDiscussionCommentPayload                                 "json:\"createTeamDiscussionComment\" graphql:\"createTeamDiscussionComment\""
	DeclineTopicSuggestion                                      *DeclineTopicSuggestionPayload                                      "json:\"declineTopicSuggestion\" graphql:\"declineTopicSuggestion\""
	DeleteBranchProtectionRule                                  *DeleteBranchProtectionRulePayload                                  "json:\"deleteBranchProtectionRule\" graphql:\"deleteBranchProtectionRule\""
	DeleteDeployment                                            *DeleteDeploymentPayload                                            "json:\"deleteDeployment\" graphql:\"deleteDeployment\""
	DeleteIPAllowListEntry                                      *DeleteIPAllowListEntryPayload                                      "json:\"deleteIpAllowListEntry\" graphql:\"deleteIpAllowListEntry\""
	DeleteIssue                                                 *DeleteIssuePayload                                                 "json:\"deleteIssue\" graphql:\"deleteIssue\""
	DeleteIssueComment                                          *DeleteIssueCommentPayload                                          "json:\"deleteIssueComment\" graphql:\"deleteIssueComment\""
	DeleteProject                                               *DeleteProjectPayload                                               "json:\"deleteProject\" graphql:\"deleteProject\""
	DeleteProjectCard                                           *DeleteProjectCardPayload                                           "json:\"deleteProjectCard\" graphql:\"deleteProjectCard\""
	DeleteProjectColumn                                         *DeleteProjectColumnPayload                                         "json:\"deleteProjectColumn\" graphql:\"deleteProjectColumn\""
	DeletePullRequestReview                                     *DeletePullRequestReviewPayload                                     "json:\"deletePullRequestReview\" graphql:\"deletePullRequestReview\""
	DeletePullRequestReviewComment                              *DeletePullRequestReviewCommentPayload                              "json:\"deletePullRequestReviewComment\" graphql:\"deletePullRequestReviewComment\""
	DeleteRef                                                   *DeleteRefPayload                                                   "json:\"deleteRef\" graphql:\"deleteRef\""
	DeleteTeamDiscussion                                        *DeleteTeamDiscussionPayload                                        "json:\"deleteTeamDiscussion\" graphql:\"deleteTeamDiscussion\""
	DeleteTeamDiscussionComment                                 *DeleteTeamDiscussionCommentPayload                                 "json:\"deleteTeamDiscussionComment\" graphql:\"deleteTeamDiscussionComment\""
	DismissPullRequestReview                                    *DismissPullRequestReviewPayload                                    "json:\"dismissPullRequestReview\" graphql:\"dismissPullRequestReview\""
	FollowUser                                                  *FollowUserPayload                                                  "json:\"followUser\" graphql:\"followUser\""
	InviteEnterpriseAdmin                                       *InviteEnterpriseAdminPayload                                       "json:\"inviteEnterpriseAdmin\" graphql:\"inviteEnterpriseAdmin\""
	LinkRepositoryToProject                                     *LinkRepositoryToProjectPayload                                     "json:\"linkRepositoryToProject\" graphql:\"linkRepositoryToProject\""
	LockLockable                                                *LockLockablePayload                                                "json:\"lockLockable\" graphql:\"lockLockable\""
	MarkFileAsViewed                                            *MarkFileAsViewedPayload                                            "json:\"markFileAsViewed\" graphql:\"markFileAsViewed\""
	MarkPullRequestReadyForReview                               *MarkPullRequestReadyForReviewPayload                               "json:\"markPullRequestReadyForReview\" graphql:\"markPullRequestReadyForReview\""
	MergeBranch                                                 *MergeBranchPayload                                                 "json:\"mergeBranch\" graphql:\"mergeBranch\""
	MergePullRequest                                            *MergePullRequestPayload                                            "json:\"mergePullRequest\" graphql:\"mergePullRequest\""
	MinimizeComment                                             *MinimizeCommentPayload                                             "json:\"minimizeComment\" graphql:\"minimizeComment\""
	MoveProjectCard                                             *MoveProjectCardPayload                                             "json:\"moveProjectCard\" graphql:\"moveProjectCard\""
	MoveProjectColumn                                           *MoveProjectColumnPayload                                           "json:\"moveProjectColumn\" graphql:\"moveProjectColumn\""
	RegenerateEnterpriseIdentityProviderRecoveryCodes           *RegenerateEnterpriseIdentityProviderRecoveryCodesPayload           "json:\"regenerateEnterpriseIdentityProviderRecoveryCodes\" graphql:\"regenerateEnterpriseIdentityProviderRecoveryCodes\""
	RemoveAssigneesFromAssignable                               *RemoveAssigneesFromAssignablePayload                               "json:\"removeAssigneesFromAssignable\" graphql:\"removeAssigneesFromAssignable\""
	RemoveEnterpriseAdmin                                       *RemoveEnterpriseAdminPayload                                       "json:\"removeEnterpriseAdmin\" graphql:\"removeEnterpriseAdmin\""
	RemoveEnterpriseIdentityProvider                            *RemoveEnterpriseIdentityProviderPayload                            "json:\"removeEnterpriseIdentityProvider\" graphql:\"removeEnterpriseIdentityProvider\""
	RemoveEnterpriseOrganization                                *RemoveEnterpriseOrganizationPayload                                "json:\"removeEnterpriseOrganization\" graphql:\"removeEnterpriseOrganization\""
	RemoveLabelsFromLabelable                                   *RemoveLabelsFromLabelablePayload                                   "json:\"removeLabelsFromLabelable\" graphql:\"removeLabelsFromLabelable\""
	RemoveOutsideCollaborator                                   *RemoveOutsideCollaboratorPayload                                   "json:\"removeOutsideCollaborator\" graphql:\"removeOutsideCollaborator\""
	RemoveReaction                                              *RemoveReactionPayload                                              "json:\"removeReaction\" graphql:\"removeReaction\""
	RemoveStar                                                  *RemoveStarPayload                                                  "json:\"removeStar\" graphql:\"removeStar\""
	ReopenIssue                                                 *ReopenIssuePayload                                                 "json:\"reopenIssue\" graphql:\"reopenIssue\""
	ReopenPullRequest                                           *ReopenPullRequestPayload                                           "json:\"reopenPullRequest\" graphql:\"reopenPullRequest\""
	RequestReviews                                              *RequestReviewsPayload                                              "json:\"requestReviews\" graphql:\"requestReviews\""
	RerequestCheckSuite                                         *RerequestCheckSuitePayload                                         "json:\"rerequestCheckSuite\" graphql:\"rerequestCheckSuite\""
	ResolveReviewThread                                         *ResolveReviewThreadPayload                                         "json:\"resolveReviewThread\" graphql:\"resolveReviewThread\""
	SetEnterpriseIdentityProvider                               *SetEnterpriseIdentityProviderPayload                               "json:\"setEnterpriseIdentityProvider\" graphql:\"setEnterpriseIdentityProvider\""
	SubmitPullRequestReview                                     *SubmitPullRequestReviewPayload                                     "json:\"submitPullRequestReview\" graphql:\"submitPullRequestReview\""
	TransferIssue                                               *TransferIssuePayload                                               "json:\"transferIssue\" graphql:\"transferIssue\""
	UnarchiveRepository                                         *UnarchiveRepositoryPayload                                         "json:\"unarchiveRepository\" graphql:\"unarchiveRepository\""
	UnfollowUser                                                *UnfollowUserPayload                                                "json:\"unfollowUser\" graphql:\"unfollowUser\""
	UnlinkRepositoryFromProject                                 *UnlinkRepositoryFromProjectPayload                                 "json:\"unlinkRepositoryFromProject\" graphql:\"unlinkRepositoryFromProject\""
	UnlockLockable                                              *UnlockLockablePayload                                              "json:\"unlockLockable\" graphql:\"unlockLockable\""
	UnmarkFileAsViewed                                          *UnmarkFileAsViewedPayload                                          "json:\"unmarkFileAsViewed\" graphql:\"unmarkFileAsViewed\""
	UnmarkIssueAsDuplicate                                      *UnmarkIssueAsDuplicatePayload                                      "json:\"unmarkIssueAsDuplicate\" graphql:\"unmarkIssueAsDuplicate\""
	UnminimizeComment                                           *UnminimizeCommentPayload                                           "json:\"unminimizeComment\" graphql:\"unminimizeComment\""
	UnresolveReviewThread                                       *UnresolveReviewThreadPayload                                       "json:\"unresolveReviewThread\" graphql:\"unresolveReviewThread\""
	UpdateBranchProtectionRule                                  *UpdateBranchProtectionRulePayload                                  "json:\"updateBranchProtectionRule\" graphql:\"updateBranchProtectionRule\""
	UpdateCheckRun                                              *UpdateCheckRunPayload                                              "json:\"updateCheckRun\" graphql:\"updateCheckRun\""
	UpdateCheckSuitePreferences                                 *UpdateCheckSuitePreferencesPayload                                 "json:\"updateCheckSuitePreferences\" graphql:\"updateCheckSuitePreferences\""
	UpdateEnterpriseActionExecutionCapabilitySetting            *UpdateEnterpriseActionExecutionCapabilitySettingPayload            "json:\"updateEnterpriseActionExecutionCapabilitySetting\" graphql:\"updateEnterpriseActionExecutionCapabilitySetting\""
	UpdateEnterpriseAdministratorRole                           *UpdateEnterpriseAdministratorRolePayload                           "json:\"updateEnterpriseAdministratorRole\" graphql:\"updateEnterpriseAdministratorRole\""
	UpdateEnterpriseAllowPrivateRepositoryForkingSetting        *UpdateEnterpriseAllowPrivateRepositoryForkingSettingPayload        "json:\"updateEnterpriseAllowPrivateRepositoryForkingSetting\" graphql:\"updateEnterpriseAllowPrivateRepositoryForkingSetting\""
	UpdateEnterpriseDefaultRepositoryPermissionSetting          *UpdateEnterpriseDefaultRepositoryPermissionSettingPayload          "json:\"updateEnterpriseDefaultRepositoryPermissionSetting\" graphql:\"updateEnterpriseDefaultRepositoryPermissionSetting\""
	UpdateEnterpriseMembersCanChangeRepositoryVisibilitySetting *UpdateEnterpriseMembersCanChangeRepositoryVisibilitySettingPayload "json:\"updateEnterpriseMembersCanChangeRepositoryVisibilitySetting\" graphql:\"updateEnterpriseMembersCanChangeRepositoryVisibilitySetting\""
	UpdateEnterpriseMembersCanCreateRepositoriesSetting         *UpdateEnterpriseMembersCanCreateRepositoriesSettingPayload         "json:\"updateEnterpriseMembersCanCreateRepositoriesSetting\" graphql:\"updateEnterpriseMembersCanCreateRepositoriesSetting\""
	UpdateEnterpriseMembersCanDeleteIssuesSetting               *UpdateEnterpriseMembersCanDeleteIssuesSettingPayload               "json:\"updateEnterpriseMembersCanDeleteIssuesSetting\" graphql:\"updateEnterpriseMembersCanDeleteIssuesSetting\""
	UpdateEnterpriseMembersCanDeleteRepositoriesSetting         *UpdateEnterpriseMembersCanDeleteRepositoriesSettingPayload         "json:\"updateEnterpriseMembersCanDeleteRepositoriesSetting\" graphql:\"updateEnterpriseMembersCanDeleteRepositoriesSetting\""
	UpdateEnterpriseMembersCanInviteCollaboratorsSetting        *UpdateEnterpriseMembersCanInviteCollaboratorsSettingPayload        "json:\"updateEnterpriseMembersCanInviteCollaboratorsSetting\" graphql:\"updateEnterpriseMembersCanInviteCollaboratorsSetting\""
	UpdateEnterpriseMembersCanMakePurchasesSetting              *UpdateEnterpriseMembersCanMakePurchasesSettingPayload              "json:\"updateEnterpriseMembersCanMakePurchasesSetting\" graphql:\"updateEnterpriseMembersCanMakePurchasesSetting\""
	UpdateEnterpriseMembersCanUpdateProtectedBranchesSetting    *UpdateEnterpriseMembersCanUpdateProtectedBranchesSettingPayload    "json:\"updateEnterpriseMembersCanUpdateProtectedBranchesSetting\" graphql:\"updateEnterpriseMembersCanUpdateProtectedBranchesSetting\""
	UpdateEnterpriseMembersCanViewDependencyInsightsSetting     *UpdateEnterpriseMembersCanViewDependencyInsightsSettingPayload     "json:\"updateEnterpriseMembersCanViewDependencyInsightsSetting\" graphql:\"updateEnterpriseMembersCanViewDependencyInsightsSetting\""
	UpdateEnterpriseOrganizationProjectsSetting                 *UpdateEnterpriseOrganizationProjectsSettingPayload                 "json:\"updateEnterpriseOrganizationProjectsSetting\" graphql:\"updateEnterpriseOrganizationProjectsSetting\""
	UpdateEnterpriseProfile                                     *UpdateEnterpriseProfilePayload                                     "json:\"updateEnterpriseProfile\" graphql:\"updateEnterpriseProfile\""
	UpdateEnterpriseRepositoryProjectsSetting                   *UpdateEnterpriseRepositoryProjectsSettingPayload                   "json:\"updateEnterpriseRepositoryProjectsSetting\" graphql:\"updateEnterpriseRepositoryProjectsSetting\""
	UpdateEnterpriseTeamDiscussionsSetting                      *UpdateEnterpriseTeamDiscussionsSettingPayload                      "json:\"updateEnterpriseTeamDiscussionsSetting\" graphql:\"updateEnterpriseTeamDiscussionsSetting\""
	UpdateEnterpriseTwoFactorAuthenticationRequiredSetting      *UpdateEnterpriseTwoFactorAuthenticationRequiredSettingPayload      "json:\"updateEnterpriseTwoFactorAuthenticationRequiredSetting\" graphql:\"updateEnterpriseTwoFactorAuthenticationRequiredSetting\""
	UpdateIPAllowListEnabledSetting                             *UpdateIPAllowListEnabledSettingPayload                             "json:\"updateIpAllowListEnabledSetting\" graphql:\"updateIpAllowListEnabledSetting\""
	UpdateIPAllowListEntry                                      *UpdateIPAllowListEntryPayload                                      "json:\"updateIpAllowListEntry\" graphql:\"updateIpAllowListEntry\""
	UpdateIssue                                                 *UpdateIssuePayload                                                 "json:\"updateIssue\" graphql:\"updateIssue\""
	UpdateIssueComment                                          *UpdateIssueCommentPayload                                          "json:\"updateIssueComment\" graphql:\"updateIssueComment\""
	UpdateProject                                               *UpdateProjectPayload                                               "json:\"updateProject\" graphql:\"updateProject\""
	UpdateProjectCard                                           *UpdateProjectCardPayload                                           "json:\"updateProjectCard\" graphql:\"updateProjectCard\""
	UpdateProjectColumn                                         *UpdateProjectColumnPayload                                         "json:\"updateProjectColumn\" graphql:\"updateProjectColumn\""
	UpdatePullRequest                                           *UpdatePullRequestPayload                                           "json:\"updatePullRequest\" graphql:\"updatePullRequest\""
	UpdatePullRequestReview                                     *UpdatePullRequestReviewPayload                                     "json:\"updatePullRequestReview\" graphql:\"updatePullRequestReview\""
	UpdatePullRequestReviewComment                              *UpdatePullRequestReviewCommentPayload                              "json:\"updatePullRequestReviewComment\" graphql:\"updatePullRequestReviewComment\""
	UpdateRef                                                   *UpdateRefPayload                                                   "json:\"updateRef\" graphql:\"updateRef\""
	UpdateRepository                                            *UpdateRepositoryPayload                                            "json:\"updateRepository\" graphql:\"updateRepository\""
	UpdateSubscription                                          *UpdateSubscriptionPayload                                          "json:\"updateSubscription\" graphql:\"updateSubscription\""
	UpdateTeamDiscussion                                        *UpdateTeamDiscussionPayload                                        "json:\"updateTeamDiscussion\" graphql:\"updateTeamDiscussion\""
	UpdateTeamDiscussionComment                                 *UpdateTeamDiscussionCommentPayload                                 "json:\"updateTeamDiscussionComment\" graphql:\"updateTeamDiscussionComment\""
	UpdateTopics                                                *UpdateTopicsPayload                                                "json:\"updateTopics\" graphql:\"updateTopics\""
}

// Represents an object which can take actions on GitHub. Typically a User or Bot.
type Actor interface {
	IsActor()
}

// An object that can have users assigned to it.
type Assignable interface {
	IsAssignable()
}

// Types that can be assigned to issues.
type Assignee interface {
	IsAssignee()
}

// An entry in the audit log.
type AuditEntry interface {
	IsAuditEntry()
}

// Types that can initiate an audit log event.
type AuditEntryActor interface {
	IsAuditEntryActor()
}

// An object that can be closed
type Closable interface {
	IsClosable()
}

// The object which triggered a `ClosedEvent`.
type Closer interface {
	IsCloser()
}

// Represents a comment.
type Comment interface {
	IsComment()
}

// Represents a contribution a user made on GitHub, such as opening an issue.
type Contribution interface {
	IsContribution()
}

// Represents either a issue the viewer can access or a restricted contribution.
type CreatedIssueOrRestrictedContribution interface {
	IsCreatedIssueOrRestrictedContribution()
}

// Represents either a pull request the viewer can access or a restricted contribution.
type CreatedPullRequestOrRestrictedContribution interface {
	IsCreatedPullRequestOrRestrictedContribution()
}

// Represents either a repository the viewer can access or a restricted contribution.
type CreatedRepositoryOrRestrictedContribution interface {
	IsCreatedRepositoryOrRestrictedContribution()
}

// Entities that can be deleted.
type Deletable interface {
	IsDeletable()
}

// Metadata for an audit entry containing enterprise account information.
type EnterpriseAuditEntryData interface {
	IsEnterpriseAuditEntryData()
}

// An object that is a member of an enterprise.
type EnterpriseMember interface {
	IsEnterpriseMember()
}

// Represents a Git object.
type GitObject interface {
	IsGitObject()
}

// Information about a signature (GPG or S/MIME) on a Commit or Tag.
type GitSignature interface {
	IsGitSignature()
}

// An individual line of a hovercard
type HovercardContext interface {
	IsHovercardContext()
}

// Types that can own an IP allow list.
type IPAllowListOwner interface {
	IsIPAllowListOwner()
}

// Used for return value of Repository.issueOrPullRequest.
type IssueOrPullRequest interface {
	IsIssueOrPullRequest()
}

// An item in an issue timeline
type IssueTimelineItem interface {
	IsIssueTimelineItem()
}

// An item in an issue timeline
type IssueTimelineItems interface {
	IsIssueTimelineItems()
}

// An object that can have labels assigned to it.
type Labelable interface {
	IsLabelable()
}

// An object that can be locked.
type Lockable interface {
	IsLockable()
}

// Entities that have members who can set status messages.
type MemberStatusable interface {
	IsMemberStatusable()
}

// Types that can be inside a Milestone.
type MilestoneItem interface {
	IsMilestoneItem()
}

// Entities that can be minimized.
type Minimizable interface {
	IsMinimizable()
}

// An object with an ID.
type Node interface {
	IsNode()
}

// Metadata for an audit entry with action oauth_application.*
type OauthApplicationAuditEntryData interface {
	IsOauthApplicationAuditEntryData()
}

// Types of memberships that can be restored for an Organization member.
type OrgRestoreMemberAuditEntryMembership interface {
	IsOrgRestoreMemberAuditEntryMembership()
}

// An audit entry in an organization audit log.
type OrganizationAuditEntry interface {
	IsOrganizationAuditEntry()
}

// Metadata for an audit entry with action org.*
type OrganizationAuditEntryData interface {
	IsOrganizationAuditEntryData()
}

// Represents an owner of a package.
type PackageOwner interface {
	IsPackageOwner()
}

// Types that can grant permissions on a repository to a user
type PermissionGranter interface {
	IsPermissionGranter()
}

// Types that can be pinned to a profile page.
type PinnableItem interface {
	IsPinnableItem()
}

// Represents any entity on GitHub that has a profile page.
type ProfileOwner interface {
	IsProfileOwner()
}

// Types that can be inside Project Cards.
type ProjectCardItem interface {
	IsProjectCardItem()
}

// Represents an owner of a Project.
type ProjectOwner interface {
	IsProjectOwner()
}

// An item in an pull request timeline
type PullRequestTimelineItem interface {
	IsPullRequestTimelineItem()
}

// An item in a pull request timeline
type PullRequestTimelineItems interface {
	IsPullRequestTimelineItems()
}

// Types that can be an actor.
type PushAllowanceActor interface {
	IsPushAllowanceActor()
}

// Represents a subject that can be reacted on.
type Reactable interface {
	IsReactable()
}

// Any referencable object
type ReferencedSubject interface {
	IsReferencedSubject()
}

// An object which has a renamable title
type RenamedTitleSubject interface {
	IsRenamedTitleSubject()
}

// Metadata for an audit entry with action repo.*
type RepositoryAuditEntryData interface {
	IsRepositoryAuditEntryData()
}

// A subset of repository info.
type RepositoryInfo interface {
	IsRepositoryInfo()
}

// Represents a object that belongs to a repository.
type RepositoryNode interface {
	IsRepositoryNode()
}

// Represents an owner of a Repository.
type RepositoryOwner interface {
	IsRepositoryOwner()
}

// Types that can be requested reviewers.
type RequestedReviewer interface {
	IsRequestedReviewer()
}

// Types that can be an actor.
type ReviewDismissalAllowanceActor interface {
	IsReviewDismissalAllowanceActor()
}

// The results of a search.
type SearchResultItem interface {
	IsSearchResultItem()
}

// Entites that can sponsor others via GitHub Sponsors
type Sponsor interface {
	IsSponsor()
}

// Entities that can be sponsored through GitHub Sponsors
type Sponsorable interface {
	IsSponsorable()
}

// Things that can be starred.
type Starrable interface {
	IsStarrable()
}

// Types that can be inside a StatusCheckRollup context.
type StatusCheckRollupContext interface {
	IsStatusCheckRollupContext()
}

// Entities that can be subscribed to for web and email notifications.
type Subscribable interface {
	IsSubscribable()
}

// Metadata for an audit entry with action team.*
type TeamAuditEntryData interface {
	IsTeamAuditEntryData()
}

// Metadata for an audit entry with a topic.
type TopicAuditEntryData interface {
	IsTopicAuditEntryData()
}

// Represents a type that can be retrieved by a URL.
type UniformResourceLocatable interface {
	IsUniformResourceLocatable()
}

// Entities that can be updated.
type Updatable interface {
	IsUpdatable()
}

// Comments that can be updated.
type UpdatableComment interface {
	IsUpdatableComment()
}

// Autogenerated input type of AcceptEnterpriseAdministratorInvitation
type AcceptEnterpriseAdministratorInvitationInput struct {
	// The id of the invitation being accepted
	InvitationID string `json:"invitationId"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of AcceptEnterpriseAdministratorInvitation
type AcceptEnterpriseAdministratorInvitationPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The invitation that was accepted.
	Invitation *EnterpriseAdministratorInvitation `json:"invitation"`
	// A message confirming the result of accepting an administrator invitation.
	Message *string `json:"message"`
}

// Autogenerated input type of AcceptTopicSuggestion
type AcceptTopicSuggestionInput struct {
	// The Node ID of the repository.
	RepositoryID string `json:"repositoryId"`
	// The name of the suggested topic.
	Name string `json:"name"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of AcceptTopicSuggestion
type AcceptTopicSuggestionPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The accepted topic.
	Topic *Topic `json:"topic"`
}

// Location information for an actor
type ActorLocation struct {
	// City
	City *string `json:"city"`
	// Country name
	Country *string `json:"country"`
	// Country code
	CountryCode *string `json:"countryCode"`
	// Region name
	Region *string `json:"region"`
	// Region or state code
	RegionCode *string `json:"regionCode"`
}

// Autogenerated input type of AddAssigneesToAssignable
type AddAssigneesToAssignableInput struct {
	// The id of the assignable object to add assignees to.
	AssignableID string `json:"assignableId"`
	// The id of users to add as assignees.
	AssigneeIds []string `json:"assigneeIds"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of AddAssigneesToAssignable
type AddAssigneesToAssignablePayload struct {
	// The item that was assigned.
	Assignable Assignable `json:"assignable"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated input type of AddComment
type AddCommentInput struct {
	// The Node ID of the subject to modify.
	SubjectID string `json:"subjectId"`
	// The contents of the comment.
	Body string `json:"body"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of AddComment
type AddCommentPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The edge from the subject's comment connection.
	CommentEdge *IssueCommentEdge `json:"commentEdge"`
	// The subject
	Subject Node `json:"subject"`
	// The edge from the subject's timeline connection.
	TimelineEdge *IssueTimelineItemEdge `json:"timelineEdge"`
}

// Autogenerated input type of AddLabelsToLabelable
type AddLabelsToLabelableInput struct {
	// The id of the labelable object to add labels to.
	LabelableID string `json:"labelableId"`
	// The ids of the labels to add.
	LabelIds []string `json:"labelIds"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of AddLabelsToLabelable
type AddLabelsToLabelablePayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The item that was labeled.
	Labelable Labelable `json:"labelable"`
}

// Autogenerated input type of AddProjectCard
type AddProjectCardInput struct {
	// The Node ID of the ProjectColumn.
	ProjectColumnID string `json:"projectColumnId"`
	// The content of the card. Must be a member of the ProjectCardItem union
	ContentID *string `json:"contentId"`
	// The note on the card.
	Note *string `json:"note"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of AddProjectCard
type AddProjectCardPayload struct {
	// The edge from the ProjectColumn's card connection.
	CardEdge *ProjectCardEdge `json:"cardEdge"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The ProjectColumn
	ProjectColumn *ProjectColumn `json:"projectColumn"`
}

// Autogenerated input type of AddProjectColumn
type AddProjectColumnInput struct {
	// The Node ID of the project.
	ProjectID string `json:"projectId"`
	// The name of the column.
	Name string `json:"name"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of AddProjectColumn
type AddProjectColumnPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The edge from the project's column connection.
	ColumnEdge *ProjectColumnEdge `json:"columnEdge"`
	// The project
	Project *Project `json:"project"`
}

// Autogenerated input type of AddPullRequestReviewComment
type AddPullRequestReviewCommentInput struct {
	// The node ID of the pull request reviewing
	PullRequestID *string `json:"pullRequestId"`
	// The Node ID of the review to modify.
	PullRequestReviewID *string `json:"pullRequestReviewId"`
	// The SHA of the commit to comment on.
	CommitOid *string `json:"commitOID"`
	// The text of the comment.
	Body string `json:"body"`
	// The relative path of the file to comment on.
	Path *string `json:"path"`
	// The line index in the diff to comment on.
	Position *int `json:"position"`
	// The comment id to reply to.
	InReplyTo *string `json:"inReplyTo"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of AddPullRequestReviewComment
type AddPullRequestReviewCommentPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The newly created comment.
	Comment *PullRequestReviewComment `json:"comment"`
	// The edge from the review's comment connection.
	CommentEdge *PullRequestReviewCommentEdge `json:"commentEdge"`
}

// Autogenerated input type of AddPullRequestReview
type AddPullRequestReviewInput struct {
	// The Node ID of the pull request to modify.
	PullRequestID string `json:"pullRequestId"`
	// The commit OID the review pertains to.
	CommitOid *string `json:"commitOID"`
	// The contents of the review body comment.
	Body *string `json:"body"`
	// The event to perform on the pull request review.
	Event *PullRequestReviewEvent `json:"event"`
	// The review line comments.
	Comments []*DraftPullRequestReviewComment `json:"comments"`
	// The review line comment threads.
	Threads []*DraftPullRequestReviewThread `json:"threads"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of AddPullRequestReview
type AddPullRequestReviewPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The newly created pull request review.
	PullRequestReview *PullRequestReview `json:"pullRequestReview"`
	// The edge from the pull request's review connection.
	ReviewEdge *PullRequestReviewEdge `json:"reviewEdge"`
}

// Autogenerated input type of AddPullRequestReviewThread
type AddPullRequestReviewThreadInput struct {
	// Path to the file being commented on.
	Path string `json:"path"`
	// Body of the thread's first comment.
	Body string `json:"body"`
	// The node ID of the pull request reviewing
	PullRequestID *string `json:"pullRequestId"`
	// The Node ID of the review to modify.
	PullRequestReviewID *string `json:"pullRequestReviewId"`
	// The line of the blob to which the thread refers. The end of the line range for multi-line comments.
	Line int `json:"line"`
	// The side of the diff on which the line resides. For multi-line comments, this is the side for the end of the line range.
	Side *DiffSide `json:"side"`
	// The first line of the range to which the comment refers.
	StartLine *int `json:"startLine"`
	// The side of the diff on which the start line resides.
	StartSide *DiffSide `json:"startSide"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of AddPullRequestReviewThread
type AddPullRequestReviewThreadPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The newly created thread.
	Thread *PullRequestReviewThread `json:"thread"`
}

// Autogenerated input type of AddReaction
type AddReactionInput struct {
	// The Node ID of the subject to modify.
	SubjectID string `json:"subjectId"`
	// The name of the emoji to react with.
	Content ReactionContent `json:"content"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of AddReaction
type AddReactionPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The reaction object.
	Reaction *Reaction `json:"reaction"`
	// The reactable subject.
	Subject Reactable `json:"subject"`
}

// Autogenerated input type of AddStar
type AddStarInput struct {
	// The Starrable ID to star.
	StarrableID string `json:"starrableId"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of AddStar
type AddStarPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The starrable.
	Starrable Starrable `json:"starrable"`
}

// Represents a 'added_to_project' event on a given issue or pull request.
type AddedToProjectEvent struct {
	// Identifies the actor who performed the event.
	Actor Actor `json:"actor"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	// Identifies the primary key from the database.
	DatabaseID *int   `json:"databaseId"`
	ID         string `json:"id"`
}

func (AddedToProjectEvent) IsNode()                     {}
func (AddedToProjectEvent) IsPullRequestTimelineItems() {}
func (AddedToProjectEvent) IsIssueTimelineItems()       {}

// A GitHub App.
type App struct {
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	// Identifies the primary key from the database.
	DatabaseID *int `json:"databaseId"`
	// The description of the app.
	Description *string `json:"description"`
	ID          string  `json:"id"`
	// The hex color code, without the leading '#', for the logo background.
	LogoBackgroundColor string `json:"logoBackgroundColor"`
	// A URL pointing to the app's logo.
	LogoURL string `json:"logoUrl"`
	// The name of the app.
	Name string `json:"name"`
	// A slug based on the name of the app for use in URLs.
	Slug string `json:"slug"`
	// Identifies the date and time when the object was last updated.
	UpdatedAt string `json:"updatedAt"`
	// The URL to the app's homepage.
	URL string `json:"url"`
}

func (App) IsSearchResultItem()   {}
func (App) IsNode()               {}
func (App) IsPushAllowanceActor() {}

// Autogenerated input type of ArchiveRepository
type ArchiveRepositoryInput struct {
	// The ID of the repository to mark as archived.
	RepositoryID string `json:"repositoryId"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of ArchiveRepository
type ArchiveRepositoryPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The repository that was marked as archived.
	Repository *Repository `json:"repository"`
}

// Represents an 'assigned' event on any assignable object.
type AssignedEvent struct {
	// Identifies the actor who performed the event.
	Actor Actor `json:"actor"`
	// Identifies the assignable associated with the event.
	Assignable Assignable `json:"assignable"`
	// Identifies the user or mannequin that was assigned.
	Assignee Assignee `json:"assignee"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// Identifies the user who was assigned.
	User *User `json:"user"`
}

func (AssignedEvent) IsPullRequestTimelineItems() {}
func (AssignedEvent) IsIssueTimelineItems()       {}
func (AssignedEvent) IsNode()                     {}
func (AssignedEvent) IsPullRequestTimelineItem()  {}
func (AssignedEvent) IsIssueTimelineItem()        {}

// Ordering options for Audit Log connections.
type AuditLogOrder struct {
	// The field to order Audit Logs by.
	Field *AuditLogOrderField `json:"field"`
	// The ordering direction.
	Direction *OrderDirection `json:"direction"`
}

// Represents a 'automatic_base_change_failed' event on a given pull request.
type AutomaticBaseChangeFailedEvent struct {
	// Identifies the actor who performed the event.
	Actor Actor `json:"actor"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// The new base for this PR
	NewBase string `json:"newBase"`
	// The old base for this PR
	OldBase string `json:"oldBase"`
	// PullRequest referenced by event.
	PullRequest *PullRequest `json:"pullRequest"`
}

func (AutomaticBaseChangeFailedEvent) IsPullRequestTimelineItems() {}
func (AutomaticBaseChangeFailedEvent) IsNode()                     {}

// Represents a 'automatic_base_change_succeeded' event on a given pull request.
type AutomaticBaseChangeSucceededEvent struct {
	// Identifies the actor who performed the event.
	Actor Actor `json:"actor"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// The new base for this PR
	NewBase string `json:"newBase"`
	// The old base for this PR
	OldBase string `json:"oldBase"`
	// PullRequest referenced by event.
	PullRequest *PullRequest `json:"pullRequest"`
}

func (AutomaticBaseChangeSucceededEvent) IsNode()                     {}
func (AutomaticBaseChangeSucceededEvent) IsPullRequestTimelineItems() {}

// Represents a 'base_ref_changed' event on a given issue or pull request.
type BaseRefChangedEvent struct {
	// Identifies the actor who performed the event.
	Actor Actor `json:"actor"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	// Identifies the name of the base ref for the pull request after it was changed.
	CurrentRefName string `json:"currentRefName"`
	// Identifies the primary key from the database.
	DatabaseID *int   `json:"databaseId"`
	ID         string `json:"id"`
	// Identifies the name of the base ref for the pull request before it was changed.
	PreviousRefName string `json:"previousRefName"`
	// PullRequest referenced by event.
	PullRequest *PullRequest `json:"pullRequest"`
}

func (BaseRefChangedEvent) IsPullRequestTimelineItems() {}
func (BaseRefChangedEvent) IsNode()                     {}

// Represents a 'base_ref_deleted' event on a given pull request.
type BaseRefDeletedEvent struct {
	// Identifies the actor who performed the event.
	Actor Actor `json:"actor"`
	// Identifies the name of the Ref associated with the `base_ref_deleted` event.
	BaseRefName *string `json:"baseRefName"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// PullRequest referenced by event.
	PullRequest *PullRequest `json:"pullRequest"`
}

func (BaseRefDeletedEvent) IsPullRequestTimelineItems() {}
func (BaseRefDeletedEvent) IsNode()                     {}
func (BaseRefDeletedEvent) IsPullRequestTimelineItem()  {}

// Represents a 'base_ref_force_pushed' event on a given pull request.
type BaseRefForcePushedEvent struct {
	// Identifies the actor who performed the event.
	Actor Actor `json:"actor"`
	// Identifies the after commit SHA for the 'base_ref_force_pushed' event.
	AfterCommit *Commit `json:"afterCommit"`
	// Identifies the before commit SHA for the 'base_ref_force_pushed' event.
	BeforeCommit *Commit `json:"beforeCommit"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// PullRequest referenced by event.
	PullRequest *PullRequest `json:"pullRequest"`
	// Identifies the fully qualified ref name for the 'base_ref_force_pushed' event.
	Ref *Ref `json:"ref"`
}

func (BaseRefForcePushedEvent) IsPullRequestTimelineItems() {}
func (BaseRefForcePushedEvent) IsPullRequestTimelineItem()  {}
func (BaseRefForcePushedEvent) IsNode()                     {}

// Represents a Git blame.
type Blame struct {
	// The list of ranges from a Git blame.
	Ranges []*BlameRange `json:"ranges"`
}

// Represents a range of information from a Git blame.
type BlameRange struct {
	// Identifies the recency of the change, from 1 (new) to 10 (old). This is calculated as a 2-quantile and determines the length of distance between the median age of all the changes in the file and the recency of the current range's change.
	Age int `json:"age"`
	// Identifies the line author
	Commit *Commit `json:"commit"`
	// The ending line for the range
	EndingLine int `json:"endingLine"`
	// The starting line for the range
	StartingLine int `json:"startingLine"`
}

// Represents a Git blob.
type Blob struct {
	// An abbreviated version of the Git object ID
	AbbreviatedOid string `json:"abbreviatedOid"`
	// Byte size of Blob object
	ByteSize int `json:"byteSize"`
	// The HTTP path for this Git object
	CommitResourcePath string `json:"commitResourcePath"`
	// The HTTP URL for this Git object
	CommitURL string `json:"commitUrl"`
	ID        string `json:"id"`
	// Indicates whether the Blob is binary or text. Returns null if unable to determine the encoding.
	IsBinary *bool `json:"isBinary"`
	// Indicates whether the contents is truncated
	IsTruncated bool `json:"isTruncated"`
	// The Git object ID
	Oid string `json:"oid"`
	// The Repository the Git object belongs to
	Repository *Repository `json:"repository"`
	// UTF8 text data or null if the Blob is binary
	Text *string `json:"text"`
}

func (Blob) IsNode()      {}
func (Blob) IsGitObject() {}

// A special type of user which takes actions on behalf of GitHub Apps.
type Bot struct {
	// A URL pointing to the GitHub App's public avatar.
	AvatarURL string `json:"avatarUrl"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	// Identifies the primary key from the database.
	DatabaseID *int   `json:"databaseId"`
	ID         string `json:"id"`
	// The username of the actor.
	Login string `json:"login"`
	// The HTTP path for this bot
	ResourcePath string `json:"resourcePath"`
	// Identifies the date and time when the object was last updated.
	UpdatedAt string `json:"updatedAt"`
	// The HTTP URL for this bot
	URL string `json:"url"`
}

func (Bot) IsNode()                     {}
func (Bot) IsActor()                    {}
func (Bot) IsUniformResourceLocatable() {}
func (Bot) IsAuditEntryActor()          {}
func (Bot) IsAssignee()                 {}

// A branch protection rule.
type BranchProtectionRule struct {
	// A list of conflicts matching branches protection rule and other branch protection rules
	BranchProtectionRuleConflicts *BranchProtectionRuleConflictConnection `json:"branchProtectionRuleConflicts"`
	// The actor who created this branch protection rule.
	Creator Actor `json:"creator"`
	// Identifies the primary key from the database.
	DatabaseID *int `json:"databaseId"`
	// Will new commits pushed to matching branches dismiss pull request review approvals.
	DismissesStaleReviews bool   `json:"dismissesStaleReviews"`
	ID                    string `json:"id"`
	// Can admins overwrite branch protection.
	IsAdminEnforced bool `json:"isAdminEnforced"`
	// Repository refs that are protected by this rule
	MatchingRefs *RefConnection `json:"matchingRefs"`
	// Identifies the protection rule pattern.
	Pattern string `json:"pattern"`
	// A list push allowances for this branch protection rule.
	PushAllowances *PushAllowanceConnection `json:"pushAllowances"`
	// The repository associated with this branch protection rule.
	Repository *Repository `json:"repository"`
	// Number of approving reviews required to update matching branches.
	RequiredApprovingReviewCount *int `json:"requiredApprovingReviewCount"`
	// List of required status check contexts that must pass for commits to be accepted to matching branches.
	RequiredStatusCheckContexts []*string `json:"requiredStatusCheckContexts"`
	// Are approving reviews required to update matching branches.
	RequiresApprovingReviews bool `json:"requiresApprovingReviews"`
	// Are reviews from code owners required to update matching branches.
	RequiresCodeOwnerReviews bool `json:"requiresCodeOwnerReviews"`
	// Are commits required to be signed.
	RequiresCommitSignatures bool `json:"requiresCommitSignatures"`
	// Are status checks required to update matching branches.
	RequiresStatusChecks bool `json:"requiresStatusChecks"`
	// Are branches required to be up to date before merging.
	RequiresStrictStatusChecks bool `json:"requiresStrictStatusChecks"`
	// Is pushing to matching branches restricted.
	RestrictsPushes bool `json:"restrictsPushes"`
	// Is dismissal of pull request reviews restricted.
	RestrictsReviewDismissals bool `json:"restrictsReviewDismissals"`
	// A list review dismissal allowances for this branch protection rule.
	ReviewDismissalAllowances *ReviewDismissalAllowanceConnection `json:"reviewDismissalAllowances"`
}

func (BranchProtectionRule) IsNode() {}

// A conflict between two branch protection rules.
type BranchProtectionRuleConflict struct {
	// Identifies the branch protection rule.
	BranchProtectionRule *BranchProtectionRule `json:"branchProtectionRule"`
	// Identifies the conflicting branch protection rule.
	ConflictingBranchProtectionRule *BranchProtectionRule `json:"conflictingBranchProtectionRule"`
	// Identifies the branch ref that has conflicting rules
	Ref *Ref `json:"ref"`
}

// The connection type for BranchProtectionRuleConflict.
type BranchProtectionRuleConflictConnection struct {
	// A list of edges.
	Edges []*BranchProtectionRuleConflictEdge `json:"edges"`
	// A list of nodes.
	Nodes []*BranchProtectionRuleConflict `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type BranchProtectionRuleConflictEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *BranchProtectionRuleConflict `json:"node"`
}

// The connection type for BranchProtectionRule.
type BranchProtectionRuleConnection struct {
	// A list of edges.
	Edges []*BranchProtectionRuleEdge `json:"edges"`
	// A list of nodes.
	Nodes []*BranchProtectionRule `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type BranchProtectionRuleEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *BranchProtectionRule `json:"node"`
}

// Autogenerated input type of CancelEnterpriseAdminInvitation
type CancelEnterpriseAdminInvitationInput struct {
	// The Node ID of the pending enterprise administrator invitation.
	InvitationID string `json:"invitationId"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of CancelEnterpriseAdminInvitation
type CancelEnterpriseAdminInvitationPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The invitation that was canceled.
	Invitation *EnterpriseAdministratorInvitation `json:"invitation"`
	// A message confirming the result of canceling an administrator invitation.
	Message *string `json:"message"`
}

// Autogenerated input type of ChangeUserStatus
type ChangeUserStatusInput struct {
	// The emoji to represent your status. Can either be a native Unicode emoji or an emoji name with colons, e.g., :grinning:.
	Emoji *string `json:"emoji"`
	// A short description of your current status.
	Message *string `json:"message"`
	// The ID of the organization whose members will be allowed to see the status. If omitted, the status will be publicly visible.
	OrganizationID *string `json:"organizationId"`
	// Whether this status should indicate you are not fully available on GitHub, e.g., you are away.
	LimitedAvailability *bool `json:"limitedAvailability"`
	// If set, the user status will not be shown after this date.
	ExpiresAt *string `json:"expiresAt"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of ChangeUserStatus
type ChangeUserStatusPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// Your updated status.
	Status *UserStatus `json:"status"`
}

// A single check annotation.
type CheckAnnotation struct {
	// The annotation's severity level.
	AnnotationLevel *CheckAnnotationLevel `json:"annotationLevel"`
	// The path to the file that this annotation was made on.
	BlobURL string `json:"blobUrl"`
	// Identifies the primary key from the database.
	DatabaseID *int `json:"databaseId"`
	// The position of this annotation.
	Location *CheckAnnotationSpan `json:"location"`
	// The annotation's message.
	Message string `json:"message"`
	// The path that this annotation was made on.
	Path string `json:"path"`
	// Additional information about the annotation.
	RawDetails *string `json:"rawDetails"`
	// The annotation's title
	Title *string `json:"title"`
}

// The connection type for CheckAnnotation.
type CheckAnnotationConnection struct {
	// A list of edges.
	Edges []*CheckAnnotationEdge `json:"edges"`
	// A list of nodes.
	Nodes []*CheckAnnotation `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// Information from a check run analysis to specific lines of code.
type CheckAnnotationData struct {
	// The path of the file to add an annotation to.
	Path string `json:"path"`
	// The location of the annotation
	Location *CheckAnnotationRange `json:"location"`
	// Represents an annotation's information level
	AnnotationLevel CheckAnnotationLevel `json:"annotationLevel"`
	// A short description of the feedback for these lines of code.
	Message string `json:"message"`
	// The title that represents the annotation.
	Title *string `json:"title"`
	// Details about this annotation.
	RawDetails *string `json:"rawDetails"`
}

// An edge in a connection.
type CheckAnnotationEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *CheckAnnotation `json:"node"`
}

// A character position in a check annotation.
type CheckAnnotationPosition struct {
	// Column number (1 indexed).
	Column *int `json:"column"`
	// Line number (1 indexed).
	Line int `json:"line"`
}

// Information from a check run analysis to specific lines of code.
type CheckAnnotationRange struct {
	// The starting line of the range.
	StartLine int `json:"startLine"`
	// The starting column of the range.
	StartColumn *int `json:"startColumn"`
	// The ending line of the range.
	EndLine int `json:"endLine"`
	// The ending column of the range.
	EndColumn *int `json:"endColumn"`
}

// An inclusive pair of positions for a check annotation.
type CheckAnnotationSpan struct {
	// End position (inclusive).
	End *CheckAnnotationPosition `json:"end"`
	// Start position (inclusive).
	Start *CheckAnnotationPosition `json:"start"`
}

// A check run.
type CheckRun struct {
	// The check run's annotations
	Annotations *CheckAnnotationConnection `json:"annotations"`
	// The check suite that this run is a part of.
	CheckSuite *CheckSuite `json:"checkSuite"`
	// Identifies the date and time when the check run was completed.
	CompletedAt *string `json:"completedAt"`
	// The conclusion of the check run.
	Conclusion *CheckConclusionState `json:"conclusion"`
	// Identifies the primary key from the database.
	DatabaseID *int `json:"databaseId"`
	// The URL from which to find full details of the check run on the integrator's site.
	DetailsURL *string `json:"detailsUrl"`
	// A reference for the check run on the integrator's system.
	ExternalID *string `json:"externalId"`
	ID         string  `json:"id"`
	// The name of the check for this check run.
	Name string `json:"name"`
	// The permalink to the check run summary.
	Permalink string `json:"permalink"`
	// The repository associated with this check run.
	Repository *Repository `json:"repository"`
	// The HTTP path for this check run.
	ResourcePath string `json:"resourcePath"`
	// Identifies the date and time when the check run was started.
	StartedAt *string `json:"startedAt"`
	// The current status of the check run.
	Status CheckStatusState `json:"status"`
	// A string representing the check run's summary
	Summary *string `json:"summary"`
	// A string representing the check run's text
	Text *string `json:"text"`
	// A string representing the check run
	Title *string `json:"title"`
	// The HTTP URL for this check run.
	URL string `json:"url"`
}

func (CheckRun) IsStatusCheckRollupContext() {}
func (CheckRun) IsNode()                     {}
func (CheckRun) IsUniformResourceLocatable() {}

// Possible further actions the integrator can perform.
type CheckRunAction struct {
	// The text to be displayed on a button in the web UI.
	Label string `json:"label"`
	// A short explanation of what this action would do.
	Description string `json:"description"`
	// A reference for the action on the integrator's system.
	Identifier string `json:"identifier"`
}

// The connection type for CheckRun.
type CheckRunConnection struct {
	// A list of edges.
	Edges []*CheckRunEdge `json:"edges"`
	// A list of nodes.
	Nodes []*CheckRun `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type CheckRunEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *CheckRun `json:"node"`
}

// The filters that are available when fetching check runs.
type CheckRunFilter struct {
	// Filters the check runs by this type.
	CheckType *CheckRunType `json:"checkType"`
	// Filters the check runs created by this application ID.
	AppID *int `json:"appId"`
	// Filters the check runs by this name.
	CheckName *string `json:"checkName"`
	// Filters the check runs by this status.
	Status *CheckStatusState `json:"status"`
}

// Descriptive details about the check run.
type CheckRunOutput struct {
	// A title to provide for this check run.
	Title string `json:"title"`
	// The summary of the check run (supports Commonmark).
	Summary string `json:"summary"`
	// The details of the check run (supports Commonmark).
	Text *string `json:"text"`
	// The annotations that are made as part of the check run.
	Annotations []*CheckAnnotationData `json:"annotations"`
	// Images attached to the check run output displayed in the GitHub pull request UI.
	Images []*CheckRunOutputImage `json:"images"`
}

// Images attached to the check run output displayed in the GitHub pull request UI.
type CheckRunOutputImage struct {
	// The alternative text for the image.
	Alt string `json:"alt"`
	// The full URL of the image.
	ImageURL string `json:"imageUrl"`
	// A short image description.
	Caption *string `json:"caption"`
}

// A check suite.
type CheckSuite struct {
	// The GitHub App which created this check suite.
	App *App `json:"app"`
	// The name of the branch for this check suite.
	Branch *Ref `json:"branch"`
	// The check runs associated with a check suite.
	CheckRuns *CheckRunConnection `json:"checkRuns"`
	// The commit for this check suite
	Commit *Commit `json:"commit"`
	// The conclusion of this check suite.
	Conclusion *CheckConclusionState `json:"conclusion"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	// Identifies the primary key from the database.
	DatabaseID *int   `json:"databaseId"`
	ID         string `json:"id"`
	// A list of open pull requests matching the check suite.
	MatchingPullRequests *PullRequestConnection `json:"matchingPullRequests"`
	// The push that triggered this check suite.
	Push *Push `json:"push"`
	// The repository associated with this check suite.
	Repository *Repository `json:"repository"`
	// The HTTP path for this check suite
	ResourcePath string `json:"resourcePath"`
	// The status of this check suite.
	Status CheckStatusState `json:"status"`
	// Identifies the date and time when the object was last updated.
	UpdatedAt string `json:"updatedAt"`
	// The HTTP URL for this check suite
	URL string `json:"url"`
}

func (CheckSuite) IsNode() {}

// The auto-trigger preferences that are available for check suites.
type CheckSuiteAutoTriggerPreference struct {
	// The node ID of the application that owns the check suite.
	AppID string `json:"appId"`
	// Set to `true` to enable automatic creation of CheckSuite events upon pushes to the repository.
	Setting bool `json:"setting"`
}

// The connection type for CheckSuite.
type CheckSuiteConnection struct {
	// A list of edges.
	Edges []*CheckSuiteEdge `json:"edges"`
	// A list of nodes.
	Nodes []*CheckSuite `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type CheckSuiteEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *CheckSuite `json:"node"`
}

// The filters that are available when fetching check suites.
type CheckSuiteFilter struct {
	// Filters the check suites created by this application ID.
	AppID *int `json:"appId"`
	// Filters the check suites by this name.
	CheckName *string `json:"checkName"`
}

// Autogenerated input type of ClearLabelsFromLabelable
type ClearLabelsFromLabelableInput struct {
	// The id of the labelable object to clear the labels from.
	LabelableID string `json:"labelableId"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of ClearLabelsFromLabelable
type ClearLabelsFromLabelablePayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The item that was unlabeled.
	Labelable Labelable `json:"labelable"`
}

// Autogenerated input type of CloneProject
type CloneProjectInput struct {
	// The owner ID to create the project under.
	TargetOwnerID string `json:"targetOwnerId"`
	// The source project to clone.
	SourceID string `json:"sourceId"`
	// Whether or not to clone the source project's workflows.
	IncludeWorkflows bool `json:"includeWorkflows"`
	// The name of the project.
	Name string `json:"name"`
	// The description of the project.
	Body *string `json:"body"`
	// The visibility of the project, defaults to false (private).
	Public *bool `json:"public"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of CloneProject
type CloneProjectPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The id of the JobStatus for populating cloned fields.
	JobStatusID *string `json:"jobStatusId"`
	// The new cloned project.
	Project *Project `json:"project"`
}

// Autogenerated input type of CloneTemplateRepository
type CloneTemplateRepositoryInput struct {
	// The Node ID of the template repository.
	RepositoryID string `json:"repositoryId"`
	// The name of the new repository.
	Name string `json:"name"`
	// The ID of the owner for the new repository.
	OwnerID string `json:"ownerId"`
	// A short description of the new repository.
	Description *string `json:"description"`
	// Indicates the repository's visibility level.
	Visibility RepositoryVisibility `json:"visibility"`
	// Whether to copy all branches from the template to the new repository. Defaults to copying only the default branch of the template.
	IncludeAllBranches *bool `json:"includeAllBranches"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of CloneTemplateRepository
type CloneTemplateRepositoryPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The new repository.
	Repository *Repository `json:"repository"`
}

// Autogenerated input type of CloseIssue
type CloseIssueInput struct {
	// ID of the issue to be closed.
	IssueID string `json:"issueId"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of CloseIssue
type CloseIssuePayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The issue that was closed.
	Issue *Issue `json:"issue"`
}

// Autogenerated input type of ClosePullRequest
type ClosePullRequestInput struct {
	// ID of the pull request to be closed.
	PullRequestID string `json:"pullRequestId"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of ClosePullRequest
type ClosePullRequestPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The pull request that was closed.
	PullRequest *PullRequest `json:"pullRequest"`
}

// Represents a 'closed' event on any `Closable`.
type ClosedEvent struct {
	// Identifies the actor who performed the event.
	Actor Actor `json:"actor"`
	// Object that was closed.
	Closable Closable `json:"closable"`
	// Object which triggered the creation of this event.
	Closer Closer `json:"closer"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// The HTTP path for this closed event.
	ResourcePath string `json:"resourcePath"`
	// The HTTP URL for this closed event.
	URL string `json:"url"`
}

func (ClosedEvent) IsPullRequestTimelineItems() {}
func (ClosedEvent) IsIssueTimelineItems()       {}
func (ClosedEvent) IsPullRequestTimelineItem()  {}
func (ClosedEvent) IsIssueTimelineItem()        {}
func (ClosedEvent) IsNode()                     {}
func (ClosedEvent) IsUniformResourceLocatable() {}

// The Code of Conduct for a repository
type CodeOfConduct struct {
	// The body of the Code of Conduct
	Body *string `json:"body"`
	ID   string  `json:"id"`
	// The key for the Code of Conduct
	Key string `json:"key"`
	// The formal name of the Code of Conduct
	Name string `json:"name"`
	// The HTTP path for this Code of Conduct
	ResourcePath *string `json:"resourcePath"`
	// The HTTP URL for this Code of Conduct
	URL *string `json:"url"`
}

func (CodeOfConduct) IsNode() {}

// Represents a 'comment_deleted' event on a given issue or pull request.
type CommentDeletedEvent struct {
	// Identifies the actor who performed the event.
	Actor Actor `json:"actor"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	// Identifies the primary key from the database.
	DatabaseID *int `json:"databaseId"`
	// The user who authored the deleted comment.
	DeletedCommentAuthor Actor  `json:"deletedCommentAuthor"`
	ID                   string `json:"id"`
}

func (CommentDeletedEvent) IsPullRequestTimelineItems() {}
func (CommentDeletedEvent) IsIssueTimelineItems()       {}
func (CommentDeletedEvent) IsNode()                     {}

// Represents a Git commit.
type Commit struct {
	// An abbreviated version of the Git object ID
	AbbreviatedOid string `json:"abbreviatedOid"`
	// The number of additions in this commit.
	Additions int `json:"additions"`
	// The pull requests associated with a commit
	AssociatedPullRequests *PullRequestConnection `json:"associatedPullRequests"`
	// Authorship details of the commit.
	Author *GitActor `json:"author"`
	// Check if the committer and the author match.
	AuthoredByCommitter bool `json:"authoredByCommitter"`
	// The datetime when this commit was authored.
	AuthoredDate string `json:"authoredDate"`
	// The list of authors for this commit based on the git author and the Co-authored-by
	// message trailer. The git author will always be first.
	//
	Authors *GitActorConnection `json:"authors"`
	// Fetches `git blame` information.
	Blame *Blame `json:"blame"`
	// The number of changed files in this commit.
	ChangedFiles int `json:"changedFiles"`
	// The check suites associated with a commit.
	CheckSuites *CheckSuiteConnection `json:"checkSuites"`
	// Comments made on the commit.
	Comments *CommitCommentConnection `json:"comments"`
	// The HTTP path for this Git object
	CommitResourcePath string `json:"commitResourcePath"`
	// The HTTP URL for this Git object
	CommitURL string `json:"commitUrl"`
	// The datetime when this commit was committed.
	CommittedDate string `json:"committedDate"`
	// Check if commited via GitHub web UI.
	CommittedViaWeb bool `json:"committedViaWeb"`
	// Committership details of the commit.
	Committer *GitActor `json:"committer"`
	// The number of deletions in this commit.
	Deletions int `json:"deletions"`
	// The deployments associated with a commit.
	Deployments *DeploymentConnection `json:"deployments"`
	// The tree entry representing the file located at the given path.
	File *TreeEntry `json:"file"`
	// The linear commit history starting from (and including) this commit, in the same order as `git log`.
	History *CommitHistoryConnection `json:"history"`
	ID      string                   `json:"id"`
	// The Git commit message
	Message string `json:"message"`
	// The Git commit message body
	MessageBody string `json:"messageBody"`
	// The commit message body rendered to HTML.
	MessageBodyHTML string `json:"messageBodyHTML"`
	// The Git commit message headline
	MessageHeadline string `json:"messageHeadline"`
	// The commit message headline rendered to HTML.
	MessageHeadlineHTML string `json:"messageHeadlineHTML"`
	// The Git object ID
	Oid string `json:"oid"`
	// The organization this commit was made on behalf of.
	OnBehalfOf *Organization `json:"onBehalfOf"`
	// The parents of a commit.
	Parents *CommitConnection `json:"parents"`
	// The datetime when this commit was pushed.
	PushedDate *string `json:"pushedDate"`
	// The Repository this commit belongs to
	Repository *Repository `json:"repository"`
	// The HTTP path for this commit
	ResourcePath string `json:"resourcePath"`
	// Commit signing information, if present.
	Signature GitSignature `json:"signature"`
	// Status information for this commit
	Status *Status `json:"status"`
	// Check and Status rollup information for this commit.
	StatusCheckRollup *StatusCheckRollup `json:"statusCheckRollup"`
	// Returns a list of all submodules in this repository as of this Commit parsed from the .gitmodules file.
	Submodules *SubmoduleConnection `json:"submodules"`
	// Returns a URL to download a tarball archive for a repository.
	// Note: For private repositories, these links are temporary and expire after five minutes.
	TarballURL string `json:"tarballUrl"`
	// Commit's root Tree
	Tree *Tree `json:"tree"`
	// The HTTP path for the tree of this commit
	TreeResourcePath string `json:"treeResourcePath"`
	// The HTTP URL for the tree of this commit
	TreeURL string `json:"treeUrl"`
	// The HTTP URL for this commit
	URL string `json:"url"`
	// Check if the viewer is able to change their subscription status for the repository.
	ViewerCanSubscribe bool `json:"viewerCanSubscribe"`
	// Identifies if the viewer is watching, not watching, or ignoring the subscribable entity.
	ViewerSubscription *SubscriptionState `json:"viewerSubscription"`
	// Returns a URL to download a zipball archive for a repository.
	// Note: For private repositories, these links are temporary and expire after five minutes.
	ZipballURL string `json:"zipballUrl"`
}

func (Commit) IsNode()                     {}
func (Commit) IsGitObject()                {}
func (Commit) IsSubscribable()             {}
func (Commit) IsUniformResourceLocatable() {}
func (Commit) IsPullRequestTimelineItem()  {}
func (Commit) IsIssueTimelineItem()        {}
func (Commit) IsCloser()                   {}

// Specifies an author for filtering Git commits.
type CommitAuthor struct {
	// ID of a User to filter by. If non-null, only commits authored by this user will be returned. This field takes precedence over emails.
	ID *string `json:"id"`
	// Email addresses to filter by. Commits authored by any of the specified email addresses will be returned.
	Emails []string `json:"emails"`
}

// Represents a comment on a given Commit.
type CommitComment struct {
	// The actor who authored the comment.
	Author Actor `json:"author"`
	// Author's association with the subject of the comment.
	AuthorAssociation CommentAuthorAssociation `json:"authorAssociation"`
	// Identifies the comment body.
	Body string `json:"body"`
	// The body rendered to HTML.
	BodyHTML string `json:"bodyHTML"`
	// The body rendered to text.
	BodyText string `json:"bodyText"`
	// Identifies the commit associated with the comment, if the commit exists.
	Commit *Commit `json:"commit"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	// Check if this comment was created via an email reply.
	CreatedViaEmail bool `json:"createdViaEmail"`
	// Identifies the primary key from the database.
	DatabaseID *int `json:"databaseId"`
	// The actor who edited the comment.
	Editor Actor  `json:"editor"`
	ID     string `json:"id"`
	// Check if this comment was edited and includes an edit with the creation data
	IncludesCreatedEdit bool `json:"includesCreatedEdit"`
	// Returns whether or not a comment has been minimized.
	IsMinimized bool `json:"isMinimized"`
	// The moment the editor made the last edit
	LastEditedAt *string `json:"lastEditedAt"`
	// Returns why the comment was minimized.
	MinimizedReason *string `json:"minimizedReason"`
	// Identifies the file path associated with the comment.
	Path *string `json:"path"`
	// Identifies the line position associated with the comment.
	Position *int `json:"position"`
	// Identifies when the comment was published at.
	PublishedAt *string `json:"publishedAt"`
	// A list of reactions grouped by content left on the subject.
	ReactionGroups []*ReactionGroup `json:"reactionGroups"`
	// A list of Reactions left on the Issue.
	Reactions *ReactionConnection `json:"reactions"`
	// The repository associated with this node.
	Repository *Repository `json:"repository"`
	// The HTTP path permalink for this commit comment.
	ResourcePath string `json:"resourcePath"`
	// Identifies the date and time when the object was last updated.
	UpdatedAt string `json:"updatedAt"`
	// The HTTP URL permalink for this commit comment.
	URL string `json:"url"`
	// A list of edits to this content.
	UserContentEdits *UserContentEditConnection `json:"userContentEdits"`
	// Check if the current viewer can delete this object.
	ViewerCanDelete bool `json:"viewerCanDelete"`
	// Check if the current viewer can minimize this object.
	ViewerCanMinimize bool `json:"viewerCanMinimize"`
	// Can user react to this subject
	ViewerCanReact bool `json:"viewerCanReact"`
	// Check if the current viewer can update this object.
	ViewerCanUpdate bool `json:"viewerCanUpdate"`
	// Reasons why the current viewer can not update this comment.
	ViewerCannotUpdateReasons []CommentCannotUpdateReason `json:"viewerCannotUpdateReasons"`
	// Did the viewer author this comment.
	ViewerDidAuthor bool `json:"viewerDidAuthor"`
}

func (CommitComment) IsNode()             {}
func (CommitComment) IsComment()          {}
func (CommitComment) IsDeletable()        {}
func (CommitComment) IsMinimizable()      {}
func (CommitComment) IsUpdatable()        {}
func (CommitComment) IsUpdatableComment() {}
func (CommitComment) IsReactable()        {}
func (CommitComment) IsRepositoryNode()   {}

// The connection type for CommitComment.
type CommitCommentConnection struct {
	// A list of edges.
	Edges []*CommitCommentEdge `json:"edges"`
	// A list of nodes.
	Nodes []*CommitComment `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type CommitCommentEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *CommitComment `json:"node"`
}

// A thread of comments on a commit.
type CommitCommentThread struct {
	// The comments that exist in this thread.
	Comments *CommitCommentConnection `json:"comments"`
	// The commit the comments were made on.
	Commit *Commit `json:"commit"`
	ID     string  `json:"id"`
	// The file the comments were made on.
	Path *string `json:"path"`
	// The position in the diff for the commit that the comment was made on.
	Position *int `json:"position"`
	// The repository associated with this node.
	Repository *Repository `json:"repository"`
}

func (CommitCommentThread) IsPullRequestTimelineItem() {}
func (CommitCommentThread) IsNode()                    {}
func (CommitCommentThread) IsRepositoryNode()          {}

// The connection type for Commit.
type CommitConnection struct {
	// A list of edges.
	Edges []*CommitEdge `json:"edges"`
	// A list of nodes.
	Nodes []*Commit `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// Ordering options for commit contribution connections.
type CommitContributionOrder struct {
	// The field by which to order commit contributions.
	Field CommitContributionOrderField `json:"field"`
	// The ordering direction.
	Direction OrderDirection `json:"direction"`
}

// This aggregates commits made by a user within one repository.
type CommitContributionsByRepository struct {
	// The commit contributions, each representing a day.
	Contributions *CreatedCommitContributionConnection `json:"contributions"`
	// The repository in which the commits were made.
	Repository *Repository `json:"repository"`
	// The HTTP path for the user's commits to the repository in this time range.
	ResourcePath string `json:"resourcePath"`
	// The HTTP URL for the user's commits to the repository in this time range.
	URL string `json:"url"`
}

// An edge in a connection.
type CommitEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *Commit `json:"node"`
}

// The connection type for Commit.
type CommitHistoryConnection struct {
	// A list of edges.
	Edges []*CommitEdge `json:"edges"`
	// A list of nodes.
	Nodes []*Commit `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// Represents a 'connected' event on a given issue or pull request.
type ConnectedEvent struct {
	// Identifies the actor who performed the event.
	Actor Actor `json:"actor"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// Reference originated in a different repository.
	IsCrossRepository bool `json:"isCrossRepository"`
	// Issue or pull request that made the reference.
	Source ReferencedSubject `json:"source"`
	// Issue or pull request which was connected.
	Subject ReferencedSubject `json:"subject"`
}

func (ConnectedEvent) IsPullRequestTimelineItems() {}
func (ConnectedEvent) IsIssueTimelineItems()       {}
func (ConnectedEvent) IsNode()                     {}

// A calendar of contributions made on GitHub by a user.
type ContributionCalendar struct {
	// A list of hex color codes used in this calendar. The darker the color, the more contributions it represents.
	Colors []string `json:"colors"`
	// Determine if the color set was chosen because it's currently Halloween.
	IsHalloween bool `json:"isHalloween"`
	// A list of the months of contributions in this calendar.
	Months []*ContributionCalendarMonth `json:"months"`
	// The count of total contributions in the calendar.
	TotalContributions int `json:"totalContributions"`
	// A list of the weeks of contributions in this calendar.
	Weeks []*ContributionCalendarWeek `json:"weeks"`
}

// Represents a single day of contributions on GitHub by a user.
type ContributionCalendarDay struct {
	// The hex color code that represents how many contributions were made on this day compared to others in the calendar.
	Color string `json:"color"`
	// How many contributions were made by the user on this day.
	ContributionCount int `json:"contributionCount"`
	// The day this square represents.
	Date time.Time `json:"date"`
	// A number representing which day of the week this square represents, e.g., 1 is Monday.
	Weekday int `json:"weekday"`
}

// A month of contributions in a user's contribution graph.
type ContributionCalendarMonth struct {
	// The date of the first day of this month.
	FirstDay time.Time `json:"firstDay"`
	// The name of the month.
	Name string `json:"name"`
	// How many weeks started in this month.
	TotalWeeks int `json:"totalWeeks"`
	// The year the month occurred in.
	Year int `json:"year"`
}

// A week of contributions in a user's contribution graph.
type ContributionCalendarWeek struct {
	// The days of contributions in this week.
	ContributionDays []*ContributionCalendarDay `json:"contributionDays"`
	// The date of the earliest square in this week.
	FirstDay time.Time `json:"firstDay"`
}

// Ordering options for contribution connections.
type ContributionOrder struct {
	// The ordering direction.
	Direction OrderDirection `json:"direction"`
}

// A contributions collection aggregates contributions such as opened issues and commits created by a user.
type ContributionsCollection struct {
	// Commit contributions made by the user, grouped by repository.
	CommitContributionsByRepository []*CommitContributionsByRepository `json:"commitContributionsByRepository"`
	// A calendar of this user's contributions on GitHub.
	ContributionCalendar *ContributionCalendar `json:"contributionCalendar"`
	// The years the user has been making contributions with the most recent year first.
	ContributionYears []int `json:"contributionYears"`
	// Determine if this collection's time span ends in the current month.
	//
	DoesEndInCurrentMonth bool `json:"doesEndInCurrentMonth"`
	// The date of the first restricted contribution the user made in this time period. Can only be non-null when the user has enabled private contribution counts.
	EarliestRestrictedContributionDate *time.Time `json:"earliestRestrictedContributionDate"`
	// The ending date and time of this collection.
	EndedAt string `json:"endedAt"`
	// The first issue the user opened on GitHub. This will be null if that issue was opened outside the collection's time range and ignoreTimeRange is false. If the issue is not visible but the user has opted to show private contributions, a RestrictedContribution will be returned.
	FirstIssueContribution CreatedIssueOrRestrictedContribution `json:"firstIssueContribution"`
	// The first pull request the user opened on GitHub. This will be null if that pull request was opened outside the collection's time range and ignoreTimeRange is not true. If the pull request is not visible but the user has opted to show private contributions, a RestrictedContribution will be returned.
	FirstPullRequestContribution CreatedPullRequestOrRestrictedContribution `json:"firstPullRequestContribution"`
	// The first repository the user created on GitHub. This will be null if that first repository was created outside the collection's time range and ignoreTimeRange is false. If the repository is not visible, then a RestrictedContribution is returned.
	FirstRepositoryContribution CreatedRepositoryOrRestrictedContribution `json:"firstRepositoryContribution"`
	// Does the user have any more activity in the timeline that occurred prior to the collection's time range?
	HasActivityInThePast bool `json:"hasActivityInThePast"`
	// Determine if there are any contributions in this collection.
	HasAnyContributions bool `json:"hasAnyContributions"`
	// Determine if the user made any contributions in this time frame whose details are not visible because they were made in a private repository. Can only be true if the user enabled private contribution counts.
	HasAnyRestrictedContributions bool `json:"hasAnyRestrictedContributions"`
	// Whether or not the collector's time span is all within the same day.
	IsSingleDay bool `json:"isSingleDay"`
	// A list of issues the user opened.
	IssueContributions *CreatedIssueContributionConnection `json:"issueContributions"`
	// Issue contributions made by the user, grouped by repository.
	IssueContributionsByRepository []*IssueContributionsByRepository `json:"issueContributionsByRepository"`
	// When the user signed up for GitHub. This will be null if that sign up date falls outside the collection's time range and ignoreTimeRange is false.
	JoinedGitHubContribution *JoinedGitHubContribution `json:"joinedGitHubContribution"`
	// The date of the most recent restricted contribution the user made in this time period. Can only be non-null when the user has enabled private contribution counts.
	LatestRestrictedContributionDate *time.Time `json:"latestRestrictedContributionDate"`
	// When this collection's time range does not include any activity from the user, use this
	// to get a different collection from an earlier time range that does have activity.
	//
	MostRecentCollectionWithActivity *ContributionsCollection `json:"mostRecentCollectionWithActivity"`
	// Returns a different contributions collection from an earlier time range than this one
	// that does not have any contributions.
	//
	MostRecentCollectionWithoutActivity *ContributionsCollection `json:"mostRecentCollectionWithoutActivity"`
	// The issue the user opened on GitHub that received the most comments in the specified
	// time frame.
	//
	PopularIssueContribution *CreatedIssueContribution `json:"popularIssueContribution"`
	// The pull request the user opened on GitHub that received the most comments in the
	// specified time frame.
	//
	PopularPullRequestContribution *CreatedPullRequestContribution `json:"popularPullRequestContribution"`
	// Pull request contributions made by the user.
	PullRequestContributions *CreatedPullRequestContributionConnection `json:"pullRequestContributions"`
	// Pull request contributions made by the user, grouped by repository.
	PullRequestContributionsByRepository []*PullRequestContributionsByRepository `json:"pullRequestContributionsByRepository"`
	// Pull request review contributions made by the user.
	PullRequestReviewContributions *CreatedPullRequestReviewContributionConnection `json:"pullRequestReviewContributions"`
	// Pull request review contributions made by the user, grouped by repository.
	PullRequestReviewContributionsByRepository []*PullRequestReviewContributionsByRepository `json:"pullRequestReviewContributionsByRepository"`
	// A list of repositories owned by the user that the user created in this time range.
	RepositoryContributions *CreatedRepositoryContributionConnection `json:"repositoryContributions"`
	// A count of contributions made by the user that the viewer cannot access. Only non-zero when the user has chosen to share their private contribution counts.
	RestrictedContributionsCount int `json:"restrictedContributionsCount"`
	// The beginning date and time of this collection.
	StartedAt string `json:"startedAt"`
	// How many commits were made by the user in this time span.
	TotalCommitContributions int `json:"totalCommitContributions"`
	// How many issues the user opened.
	TotalIssueContributions int `json:"totalIssueContributions"`
	// How many pull requests the user opened.
	TotalPullRequestContributions int `json:"totalPullRequestContributions"`
	// How many pull request reviews the user left.
	TotalPullRequestReviewContributions int `json:"totalPullRequestReviewContributions"`
	// How many different repositories the user committed to.
	TotalRepositoriesWithContributedCommits int `json:"totalRepositoriesWithContributedCommits"`
	// How many different repositories the user opened issues in.
	TotalRepositoriesWithContributedIssues int `json:"totalRepositoriesWithContributedIssues"`
	// How many different repositories the user left pull request reviews in.
	TotalRepositoriesWithContributedPullRequestReviews int `json:"totalRepositoriesWithContributedPullRequestReviews"`
	// How many different repositories the user opened pull requests in.
	TotalRepositoriesWithContributedPullRequests int `json:"totalRepositoriesWithContributedPullRequests"`
	// How many repositories the user created.
	TotalRepositoryContributions int `json:"totalRepositoryContributions"`
	// The user who made the contributions in this collection.
	User *User `json:"user"`
}

// Autogenerated input type of ConvertProjectCardNoteToIssue
type ConvertProjectCardNoteToIssueInput struct {
	// The ProjectCard ID to convert.
	ProjectCardID string `json:"projectCardId"`
	// The ID of the repository to create the issue in.
	RepositoryID string `json:"repositoryId"`
	// The title of the newly created issue. Defaults to the card's note text.
	Title *string `json:"title"`
	// The body of the newly created issue.
	Body *string `json:"body"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of ConvertProjectCardNoteToIssue
type ConvertProjectCardNoteToIssuePayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The updated ProjectCard.
	ProjectCard *ProjectCard `json:"projectCard"`
}

// Represents a 'convert_to_draft' event on a given pull request.
type ConvertToDraftEvent struct {
	// Identifies the actor who performed the event.
	Actor Actor `json:"actor"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// PullRequest referenced by event.
	PullRequest *PullRequest `json:"pullRequest"`
	// The HTTP path for this convert to draft event.
	ResourcePath string `json:"resourcePath"`
	// The HTTP URL for this convert to draft event.
	URL string `json:"url"`
}

func (ConvertToDraftEvent) IsPullRequestTimelineItems() {}
func (ConvertToDraftEvent) IsNode()                     {}
func (ConvertToDraftEvent) IsUniformResourceLocatable() {}

// Represents a 'converted_note_to_issue' event on a given issue or pull request.
type ConvertedNoteToIssueEvent struct {
	// Identifies the actor who performed the event.
	Actor Actor `json:"actor"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	// Identifies the primary key from the database.
	DatabaseID *int   `json:"databaseId"`
	ID         string `json:"id"`
}

func (ConvertedNoteToIssueEvent) IsPullRequestTimelineItems() {}
func (ConvertedNoteToIssueEvent) IsIssueTimelineItems()       {}
func (ConvertedNoteToIssueEvent) IsNode()                     {}

// Autogenerated input type of CreateBranchProtectionRule
type CreateBranchProtectionRuleInput struct {
	// The global relay id of the repository in which a new branch protection rule should be created in.
	RepositoryID string `json:"repositoryId"`
	// The glob-like pattern used to determine matching branches.
	Pattern string `json:"pattern"`
	// Are approving reviews required to update matching branches.
	RequiresApprovingReviews *bool `json:"requiresApprovingReviews"`
	// Number of approving reviews required to update matching branches.
	RequiredApprovingReviewCount *int `json:"requiredApprovingReviewCount"`
	// Are commits required to be signed.
	RequiresCommitSignatures *bool `json:"requiresCommitSignatures"`
	// Can admins overwrite branch protection.
	IsAdminEnforced *bool `json:"isAdminEnforced"`
	// Are status checks required to update matching branches.
	RequiresStatusChecks *bool `json:"requiresStatusChecks"`
	// Are branches required to be up to date before merging.
	RequiresStrictStatusChecks *bool `json:"requiresStrictStatusChecks"`
	// Are reviews from code owners required to update matching branches.
	RequiresCodeOwnerReviews *bool `json:"requiresCodeOwnerReviews"`
	// Will new commits pushed to matching branches dismiss pull request review approvals.
	DismissesStaleReviews *bool `json:"dismissesStaleReviews"`
	// Is dismissal of pull request reviews restricted.
	RestrictsReviewDismissals *bool `json:"restrictsReviewDismissals"`
	// A list of User or Team IDs allowed to dismiss reviews on pull requests targeting matching branches.
	ReviewDismissalActorIds []string `json:"reviewDismissalActorIds"`
	// Is pushing to matching branches restricted.
	RestrictsPushes *bool `json:"restrictsPushes"`
	// A list of User, Team or App IDs allowed to push to matching branches.
	PushActorIds []string `json:"pushActorIds"`
	// List of required status check contexts that must pass for commits to be accepted to matching branches.
	RequiredStatusCheckContexts []string `json:"requiredStatusCheckContexts"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of CreateBranchProtectionRule
type CreateBranchProtectionRulePayload struct {
	// The newly created BranchProtectionRule.
	BranchProtectionRule *BranchProtectionRule `json:"branchProtectionRule"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated input type of CreateCheckRun
type CreateCheckRunInput struct {
	// The node ID of the repository.
	RepositoryID string `json:"repositoryId"`
	// The name of the check.
	Name string `json:"name"`
	// The SHA of the head commit.
	HeadSha string `json:"headSha"`
	// The URL of the integrator's site that has the full details of the check.
	DetailsURL *string `json:"detailsUrl"`
	// A reference for the run on the integrator's system.
	ExternalID *string `json:"externalId"`
	// The current status.
	Status *RequestableCheckStatusState `json:"status"`
	// The time that the check run began.
	StartedAt *string `json:"startedAt"`
	// The final conclusion of the check.
	Conclusion *CheckConclusionState `json:"conclusion"`
	// The time that the check run finished.
	CompletedAt *string `json:"completedAt"`
	// Descriptive details about the run.
	Output *CheckRunOutput `json:"output"`
	// Possible further actions the integrator can perform, which a user may trigger.
	Actions []*CheckRunAction `json:"actions"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of CreateCheckRun
type CreateCheckRunPayload struct {
	// The newly created check run.
	CheckRun *CheckRun `json:"checkRun"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated input type of CreateCheckSuite
type CreateCheckSuiteInput struct {
	// The Node ID of the repository.
	RepositoryID string `json:"repositoryId"`
	// The SHA of the head commit.
	HeadSha string `json:"headSha"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of CreateCheckSuite
type CreateCheckSuitePayload struct {
	// The newly created check suite.
	CheckSuite *CheckSuite `json:"checkSuite"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated input type of CreateEnterpriseOrganization
type CreateEnterpriseOrganizationInput struct {
	// The ID of the enterprise owning the new organization.
	EnterpriseID string `json:"enterpriseId"`
	// The login of the new organization.
	Login string `json:"login"`
	// The profile name of the new organization.
	ProfileName string `json:"profileName"`
	// The email used for sending billing receipts.
	BillingEmail string `json:"billingEmail"`
	// The logins for the administrators of the new organization.
	AdminLogins []string `json:"adminLogins"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of CreateEnterpriseOrganization
type CreateEnterpriseOrganizationPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The enterprise that owns the created organization.
	Enterprise *Enterprise `json:"enterprise"`
	// The organization that was created.
	Organization *Organization `json:"organization"`
}

// Autogenerated input type of CreateIpAllowListEntry
type CreateIPAllowListEntryInput struct {
	// The ID of the owner for which to create the new IP allow list entry.
	OwnerID string `json:"ownerId"`
	// An IP address or range of addresses in CIDR notation.
	AllowListValue string `json:"allowListValue"`
	// An optional name for the IP allow list entry.
	Name *string `json:"name"`
	// Whether the IP allow list entry is active when an IP allow list is enabled.
	IsActive bool `json:"isActive"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of CreateIpAllowListEntry
type CreateIPAllowListEntryPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The IP allow list entry that was created.
	IPAllowListEntry *IPAllowListEntry `json:"ipAllowListEntry"`
}

// Autogenerated input type of CreateIssue
type CreateIssueInput struct {
	// The Node ID of the repository.
	RepositoryID string `json:"repositoryId"`
	// The title for the issue.
	Title string `json:"title"`
	// The body for the issue description.
	Body *string `json:"body"`
	// The Node ID for the user assignee for this issue.
	AssigneeIds []string `json:"assigneeIds"`
	// The Node ID of the milestone for this issue.
	MilestoneID *string `json:"milestoneId"`
	// An array of Node IDs of labels for this issue.
	LabelIds []string `json:"labelIds"`
	// An array of Node IDs for projects associated with this issue.
	ProjectIds []string `json:"projectIds"`
	// The name of an issue template in the repository, assigns labels and assignees from the template to the issue
	IssueTemplate *string `json:"issueTemplate"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of CreateIssue
type CreateIssuePayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The new issue.
	Issue *Issue `json:"issue"`
}

// Autogenerated input type of CreateProject
type CreateProjectInput struct {
	// The owner ID to create the project under.
	OwnerID string `json:"ownerId"`
	// The name of project.
	Name string `json:"name"`
	// The description of project.
	Body *string `json:"body"`
	// The name of the GitHub-provided template.
	Template *ProjectTemplate `json:"template"`
	// A list of repository IDs to create as linked repositories for the project
	RepositoryIds []string `json:"repositoryIds"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of CreateProject
type CreateProjectPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The new project.
	Project *Project `json:"project"`
}

// Autogenerated input type of CreatePullRequest
type CreatePullRequestInput struct {
	// The Node ID of the repository.
	RepositoryID string `json:"repositoryId"`
	// The name of the branch you want your changes pulled into. This should be an existing branch
	// on the current repository. You cannot update the base branch on a pull request to point
	// to another repository.
	//
	BaseRefName string `json:"baseRefName"`
	// The name of the branch where your changes are implemented. For cross-repository pull requests
	// in the same network, namespace `head_ref_name` with a user like this: `username:branch`.
	//
	HeadRefName string `json:"headRefName"`
	// The title of the pull request.
	Title string `json:"title"`
	// The contents of the pull request.
	Body *string `json:"body"`
	// Indicates whether maintainers can modify the pull request.
	MaintainerCanModify *bool `json:"maintainerCanModify"`
	// Indicates whether this pull request should be a draft.
	Draft *bool `json:"draft"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of CreatePullRequest
type CreatePullRequestPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The new pull request.
	PullRequest *PullRequest `json:"pullRequest"`
}

// Autogenerated input type of CreateRef
type CreateRefInput struct {
	// The Node ID of the Repository to create the Ref in.
	RepositoryID string `json:"repositoryId"`
	// The fully qualified name of the new Ref (ie: `refs/heads/my_new_branch`).
	Name string `json:"name"`
	// The GitObjectID that the new Ref shall target. Must point to a commit.
	Oid string `json:"oid"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of CreateRef
type CreateRefPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The newly created ref.
	Ref *Ref `json:"ref"`
}

// Autogenerated input type of CreateRepository
type CreateRepositoryInput struct {
	// The name of the new repository.
	Name string `json:"name"`
	// The ID of the owner for the new repository.
	OwnerID *string `json:"ownerId"`
	// A short description of the new repository.
	Description *string `json:"description"`
	// Indicates the repository's visibility level.
	Visibility RepositoryVisibility `json:"visibility"`
	// Whether this repository should be marked as a template such that anyone who can access it can create new repositories with the same files and directory structure.
	Template *bool `json:"template"`
	// The URL for a web page about this repository.
	HomepageURL *string `json:"homepageUrl"`
	// Indicates if the repository should have the wiki feature enabled.
	HasWikiEnabled *bool `json:"hasWikiEnabled"`
	// Indicates if the repository should have the issues feature enabled.
	HasIssuesEnabled *bool `json:"hasIssuesEnabled"`
	// When an organization is specified as the owner, this ID identifies the team that should be granted access to the new repository.
	TeamID *string `json:"teamId"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of CreateRepository
type CreateRepositoryPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The new repository.
	Repository *Repository `json:"repository"`
}

// Autogenerated input type of CreateTeamDiscussionComment
type CreateTeamDiscussionCommentInput struct {
	// The ID of the discussion to which the comment belongs.
	DiscussionID string `json:"discussionId"`
	// The content of the comment.
	Body string `json:"body"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of CreateTeamDiscussionComment
type CreateTeamDiscussionCommentPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The new comment.
	TeamDiscussionComment *TeamDiscussionComment `json:"teamDiscussionComment"`
}

// Autogenerated input type of CreateTeamDiscussion
type CreateTeamDiscussionInput struct {
	// The ID of the team to which the discussion belongs.
	TeamID string `json:"teamId"`
	// The title of the discussion.
	Title string `json:"title"`
	// The content of the discussion.
	Body string `json:"body"`
	// If true, restricts the visiblity of this discussion to team members and organization admins. If false or not specified, allows any organization member to view this discussion.
	Private *bool `json:"private"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of CreateTeamDiscussion
type CreateTeamDiscussionPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The new discussion.
	TeamDiscussion *TeamDiscussion `json:"teamDiscussion"`
}

// Represents the contribution a user made by committing to a repository.
type CreatedCommitContribution struct {
	// How many commits were made on this day to this repository by the user.
	CommitCount int `json:"commitCount"`
	// Whether this contribution is associated with a record you do not have access to. For
	// example, your own 'first issue' contribution may have been made on a repository you can no
	// longer access.
	//
	IsRestricted bool `json:"isRestricted"`
	// When this contribution was made.
	OccurredAt string `json:"occurredAt"`
	// The repository the user made a commit in.
	Repository *Repository `json:"repository"`
	// The HTTP path for this contribution.
	ResourcePath string `json:"resourcePath"`
	// The HTTP URL for this contribution.
	URL string `json:"url"`
	// The user who made this contribution.
	//
	User *User `json:"user"`
}

func (CreatedCommitContribution) IsContribution() {}

// The connection type for CreatedCommitContribution.
type CreatedCommitContributionConnection struct {
	// A list of edges.
	Edges []*CreatedCommitContributionEdge `json:"edges"`
	// A list of nodes.
	Nodes []*CreatedCommitContribution `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of commits across days and repositories in the connection.
	//
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type CreatedCommitContributionEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *CreatedCommitContribution `json:"node"`
}

// Represents the contribution a user made on GitHub by opening an issue.
type CreatedIssueContribution struct {
	// Whether this contribution is associated with a record you do not have access to. For
	// example, your own 'first issue' contribution may have been made on a repository you can no
	// longer access.
	//
	IsRestricted bool `json:"isRestricted"`
	// The issue that was opened.
	Issue *Issue `json:"issue"`
	// When this contribution was made.
	OccurredAt string `json:"occurredAt"`
	// The HTTP path for this contribution.
	ResourcePath string `json:"resourcePath"`
	// The HTTP URL for this contribution.
	URL string `json:"url"`
	// The user who made this contribution.
	//
	User *User `json:"user"`
}

func (CreatedIssueContribution) IsContribution()                         {}
func (CreatedIssueContribution) IsCreatedIssueOrRestrictedContribution() {}

// The connection type for CreatedIssueContribution.
type CreatedIssueContributionConnection struct {
	// A list of edges.
	Edges []*CreatedIssueContributionEdge `json:"edges"`
	// A list of nodes.
	Nodes []*CreatedIssueContribution `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type CreatedIssueContributionEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *CreatedIssueContribution `json:"node"`
}

// Represents the contribution a user made on GitHub by opening a pull request.
type CreatedPullRequestContribution struct {
	// Whether this contribution is associated with a record you do not have access to. For
	// example, your own 'first issue' contribution may have been made on a repository you can no
	// longer access.
	//
	IsRestricted bool `json:"isRestricted"`
	// When this contribution was made.
	OccurredAt string `json:"occurredAt"`
	// The pull request that was opened.
	PullRequest *PullRequest `json:"pullRequest"`
	// The HTTP path for this contribution.
	ResourcePath string `json:"resourcePath"`
	// The HTTP URL for this contribution.
	URL string `json:"url"`
	// The user who made this contribution.
	//
	User *User `json:"user"`
}

func (CreatedPullRequestContribution) IsCreatedPullRequestOrRestrictedContribution() {}
func (CreatedPullRequestContribution) IsContribution()                               {}

// The connection type for CreatedPullRequestContribution.
type CreatedPullRequestContributionConnection struct {
	// A list of edges.
	Edges []*CreatedPullRequestContributionEdge `json:"edges"`
	// A list of nodes.
	Nodes []*CreatedPullRequestContribution `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type CreatedPullRequestContributionEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *CreatedPullRequestContribution `json:"node"`
}

// Represents the contribution a user made by leaving a review on a pull request.
type CreatedPullRequestReviewContribution struct {
	// Whether this contribution is associated with a record you do not have access to. For
	// example, your own 'first issue' contribution may have been made on a repository you can no
	// longer access.
	//
	IsRestricted bool `json:"isRestricted"`
	// When this contribution was made.
	OccurredAt string `json:"occurredAt"`
	// The pull request the user reviewed.
	PullRequest *PullRequest `json:"pullRequest"`
	// The review the user left on the pull request.
	PullRequestReview *PullRequestReview `json:"pullRequestReview"`
	// The repository containing the pull request that the user reviewed.
	Repository *Repository `json:"repository"`
	// The HTTP path for this contribution.
	ResourcePath string `json:"resourcePath"`
	// The HTTP URL for this contribution.
	URL string `json:"url"`
	// The user who made this contribution.
	//
	User *User `json:"user"`
}

func (CreatedPullRequestReviewContribution) IsContribution() {}

// The connection type for CreatedPullRequestReviewContribution.
type CreatedPullRequestReviewContributionConnection struct {
	// A list of edges.
	Edges []*CreatedPullRequestReviewContributionEdge `json:"edges"`
	// A list of nodes.
	Nodes []*CreatedPullRequestReviewContribution `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type CreatedPullRequestReviewContributionEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *CreatedPullRequestReviewContribution `json:"node"`
}

// Represents the contribution a user made on GitHub by creating a repository.
type CreatedRepositoryContribution struct {
	// Whether this contribution is associated with a record you do not have access to. For
	// example, your own 'first issue' contribution may have been made on a repository you can no
	// longer access.
	//
	IsRestricted bool `json:"isRestricted"`
	// When this contribution was made.
	OccurredAt string `json:"occurredAt"`
	// The repository that was created.
	Repository *Repository `json:"repository"`
	// The HTTP path for this contribution.
	ResourcePath string `json:"resourcePath"`
	// The HTTP URL for this contribution.
	URL string `json:"url"`
	// The user who made this contribution.
	//
	User *User `json:"user"`
}

func (CreatedRepositoryContribution) IsCreatedRepositoryOrRestrictedContribution() {}
func (CreatedRepositoryContribution) IsContribution()                              {}

// The connection type for CreatedRepositoryContribution.
type CreatedRepositoryContributionConnection struct {
	// A list of edges.
	Edges []*CreatedRepositoryContributionEdge `json:"edges"`
	// A list of nodes.
	Nodes []*CreatedRepositoryContribution `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type CreatedRepositoryContributionEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *CreatedRepositoryContribution `json:"node"`
}

// Represents a mention made by one issue or pull request to another.
type CrossReferencedEvent struct {
	// Identifies the actor who performed the event.
	Actor Actor `json:"actor"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// Reference originated in a different repository.
	IsCrossRepository bool `json:"isCrossRepository"`
	// Identifies when the reference was made.
	ReferencedAt string `json:"referencedAt"`
	// The HTTP path for this pull request.
	ResourcePath string `json:"resourcePath"`
	// Issue or pull request that made the reference.
	Source ReferencedSubject `json:"source"`
	// Issue or pull request to which the reference was made.
	Target ReferencedSubject `json:"target"`
	// The HTTP URL for this pull request.
	URL string `json:"url"`
	// Checks if the target will be closed when the source is merged.
	WillCloseTarget bool `json:"willCloseTarget"`
}

func (CrossReferencedEvent) IsPullRequestTimelineItems() {}
func (CrossReferencedEvent) IsIssueTimelineItems()       {}
func (CrossReferencedEvent) IsNode()                     {}
func (CrossReferencedEvent) IsUniformResourceLocatable() {}
func (CrossReferencedEvent) IsPullRequestTimelineItem()  {}
func (CrossReferencedEvent) IsIssueTimelineItem()        {}

// Autogenerated input type of DeclineTopicSuggestion
type DeclineTopicSuggestionInput struct {
	// The Node ID of the repository.
	RepositoryID string `json:"repositoryId"`
	// The name of the suggested topic.
	Name string `json:"name"`
	// The reason why the suggested topic is declined.
	Reason TopicSuggestionDeclineReason `json:"reason"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of DeclineTopicSuggestion
type DeclineTopicSuggestionPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The declined topic.
	Topic *Topic `json:"topic"`
}

// Autogenerated input type of DeleteBranchProtectionRule
type DeleteBranchProtectionRuleInput struct {
	// The global relay id of the branch protection rule to be deleted.
	BranchProtectionRuleID string `json:"branchProtectionRuleId"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of DeleteBranchProtectionRule
type DeleteBranchProtectionRulePayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated input type of DeleteDeployment
type DeleteDeploymentInput struct {
	// The Node ID of the deployment to be deleted.
	ID string `json:"id"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of DeleteDeployment
type DeleteDeploymentPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated input type of DeleteIpAllowListEntry
type DeleteIPAllowListEntryInput struct {
	// The ID of the IP allow list entry to delete.
	IPAllowListEntryID string `json:"ipAllowListEntryId"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of DeleteIpAllowListEntry
type DeleteIPAllowListEntryPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The IP allow list entry that was deleted.
	IPAllowListEntry *IPAllowListEntry `json:"ipAllowListEntry"`
}

// Autogenerated input type of DeleteIssueComment
type DeleteIssueCommentInput struct {
	// The ID of the comment to delete.
	ID string `json:"id"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of DeleteIssueComment
type DeleteIssueCommentPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated input type of DeleteIssue
type DeleteIssueInput struct {
	// The ID of the issue to delete.
	IssueID string `json:"issueId"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of DeleteIssue
type DeleteIssuePayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The repository the issue belonged to
	Repository *Repository `json:"repository"`
}

// Autogenerated input type of DeleteProjectCard
type DeleteProjectCardInput struct {
	// The id of the card to delete.
	CardID string `json:"cardId"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of DeleteProjectCard
type DeleteProjectCardPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The column the deleted card was in.
	Column *ProjectColumn `json:"column"`
	// The deleted card ID.
	DeletedCardID *string `json:"deletedCardId"`
}

// Autogenerated input type of DeleteProjectColumn
type DeleteProjectColumnInput struct {
	// The id of the column to delete.
	ColumnID string `json:"columnId"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of DeleteProjectColumn
type DeleteProjectColumnPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The deleted column ID.
	DeletedColumnID *string `json:"deletedColumnId"`
	// The project the deleted column was in.
	Project *Project `json:"project"`
}

// Autogenerated input type of DeleteProject
type DeleteProjectInput struct {
	// The Project ID to update.
	ProjectID string `json:"projectId"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of DeleteProject
type DeleteProjectPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The repository or organization the project was removed from.
	Owner ProjectOwner `json:"owner"`
}

// Autogenerated input type of DeletePullRequestReviewComment
type DeletePullRequestReviewCommentInput struct {
	// The ID of the comment to delete.
	ID string `json:"id"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of DeletePullRequestReviewComment
type DeletePullRequestReviewCommentPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The pull request review the deleted comment belonged to.
	PullRequestReview *PullRequestReview `json:"pullRequestReview"`
}

// Autogenerated input type of DeletePullRequestReview
type DeletePullRequestReviewInput struct {
	// The Node ID of the pull request review to delete.
	PullRequestReviewID string `json:"pullRequestReviewId"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of DeletePullRequestReview
type DeletePullRequestReviewPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The deleted pull request review.
	PullRequestReview *PullRequestReview `json:"pullRequestReview"`
}

// Autogenerated input type of DeleteRef
type DeleteRefInput struct {
	// The Node ID of the Ref to be deleted.
	RefID string `json:"refId"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of DeleteRef
type DeleteRefPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated input type of DeleteTeamDiscussionComment
type DeleteTeamDiscussionCommentInput struct {
	// The ID of the comment to delete.
	ID string `json:"id"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of DeleteTeamDiscussionComment
type DeleteTeamDiscussionCommentPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated input type of DeleteTeamDiscussion
type DeleteTeamDiscussionInput struct {
	// The discussion ID to delete.
	ID string `json:"id"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of DeleteTeamDiscussion
type DeleteTeamDiscussionPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Represents a 'demilestoned' event on a given issue or pull request.
type DemilestonedEvent struct {
	// Identifies the actor who performed the event.
	Actor Actor `json:"actor"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// Identifies the milestone title associated with the 'demilestoned' event.
	MilestoneTitle string `json:"milestoneTitle"`
	// Object referenced by event.
	Subject MilestoneItem `json:"subject"`
}

func (DemilestonedEvent) IsPullRequestTimelineItems() {}
func (DemilestonedEvent) IsIssueTimelineItems()       {}
func (DemilestonedEvent) IsPullRequestTimelineItem()  {}
func (DemilestonedEvent) IsIssueTimelineItem()        {}
func (DemilestonedEvent) IsNode()                     {}

// A repository deploy key.
type DeployKey struct {
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// The deploy key.
	Key string `json:"key"`
	// Whether or not the deploy key is read only.
	ReadOnly bool `json:"readOnly"`
	// The deploy key title.
	Title string `json:"title"`
	// Whether or not the deploy key has been verified.
	Verified bool `json:"verified"`
}

func (DeployKey) IsNode() {}

// The connection type for DeployKey.
type DeployKeyConnection struct {
	// A list of edges.
	Edges []*DeployKeyEdge `json:"edges"`
	// A list of nodes.
	Nodes []*DeployKey `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type DeployKeyEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *DeployKey `json:"node"`
}

// Represents a 'deployed' event on a given pull request.
type DeployedEvent struct {
	// Identifies the actor who performed the event.
	Actor Actor `json:"actor"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	// Identifies the primary key from the database.
	DatabaseID *int `json:"databaseId"`
	// The deployment associated with the 'deployed' event.
	Deployment *Deployment `json:"deployment"`
	ID         string      `json:"id"`
	// PullRequest referenced by event.
	PullRequest *PullRequest `json:"pullRequest"`
	// The ref associated with the 'deployed' event.
	Ref *Ref `json:"ref"`
}

func (DeployedEvent) IsPullRequestTimelineItems() {}
func (DeployedEvent) IsPullRequestTimelineItem()  {}
func (DeployedEvent) IsNode()                     {}

// Represents triggered deployment instance.
type Deployment struct {
	// Identifies the commit sha of the deployment.
	Commit *Commit `json:"commit"`
	// Identifies the oid of the deployment commit, even if the commit has been deleted.
	CommitOid string `json:"commitOid"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	// Identifies the actor who triggered the deployment.
	Creator Actor `json:"creator"`
	// Identifies the primary key from the database.
	DatabaseID *int `json:"databaseId"`
	// The deployment description.
	Description *string `json:"description"`
	// The latest environment to which this deployment was made.
	Environment *string `json:"environment"`
	ID          string  `json:"id"`
	// The latest environment to which this deployment was made.
	LatestEnvironment *string `json:"latestEnvironment"`
	// The latest status of this deployment.
	LatestStatus *DeploymentStatus `json:"latestStatus"`
	// The original environment to which this deployment was made.
	OriginalEnvironment *string `json:"originalEnvironment"`
	// Extra information that a deployment system might need.
	Payload *string `json:"payload"`
	// Identifies the Ref of the deployment, if the deployment was created by ref.
	Ref *Ref `json:"ref"`
	// Identifies the repository associated with the deployment.
	Repository *Repository `json:"repository"`
	// The current state of the deployment.
	State *DeploymentState `json:"state"`
	// A list of statuses associated with the deployment.
	Statuses *DeploymentStatusConnection `json:"statuses"`
	// The deployment task.
	Task *string `json:"task"`
	// Identifies the date and time when the object was last updated.
	UpdatedAt string `json:"updatedAt"`
}

func (Deployment) IsNode() {}

// The connection type for Deployment.
type DeploymentConnection struct {
	// A list of edges.
	Edges []*DeploymentEdge `json:"edges"`
	// A list of nodes.
	Nodes []*Deployment `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type DeploymentEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *Deployment `json:"node"`
}

// Represents a 'deployment_environment_changed' event on a given pull request.
type DeploymentEnvironmentChangedEvent struct {
	// Identifies the actor who performed the event.
	Actor Actor `json:"actor"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	// The deployment status that updated the deployment environment.
	DeploymentStatus *DeploymentStatus `json:"deploymentStatus"`
	ID               string            `json:"id"`
	// PullRequest referenced by event.
	PullRequest *PullRequest `json:"pullRequest"`
}

func (DeploymentEnvironmentChangedEvent) IsPullRequestTimelineItems() {}
func (DeploymentEnvironmentChangedEvent) IsPullRequestTimelineItem()  {}
func (DeploymentEnvironmentChangedEvent) IsNode()                     {}

// Ordering options for deployment connections
type DeploymentOrder struct {
	// The field to order deployments by.
	Field DeploymentOrderField `json:"field"`
	// The ordering direction.
	Direction OrderDirection `json:"direction"`
}

// Describes the status of a given deployment attempt.
type DeploymentStatus struct {
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	// Identifies the actor who triggered the deployment.
	Creator Actor `json:"creator"`
	// Identifies the deployment associated with status.
	Deployment *Deployment `json:"deployment"`
	// Identifies the description of the deployment.
	Description *string `json:"description"`
	// Identifies the environment URL of the deployment.
	EnvironmentURL *string `json:"environmentUrl"`
	ID             string  `json:"id"`
	// Identifies the log URL of the deployment.
	LogURL *string `json:"logUrl"`
	// Identifies the current state of the deployment.
	State DeploymentStatusState `json:"state"`
	// Identifies the date and time when the object was last updated.
	UpdatedAt string `json:"updatedAt"`
}

func (DeploymentStatus) IsNode() {}

// The connection type for DeploymentStatus.
type DeploymentStatusConnection struct {
	// A list of edges.
	Edges []*DeploymentStatusEdge `json:"edges"`
	// A list of nodes.
	Nodes []*DeploymentStatus `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type DeploymentStatusEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *DeploymentStatus `json:"node"`
}

// Represents a 'disconnected' event on a given issue or pull request.
type DisconnectedEvent struct {
	// Identifies the actor who performed the event.
	Actor Actor `json:"actor"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// Reference originated in a different repository.
	IsCrossRepository bool `json:"isCrossRepository"`
	// Issue or pull request from which the issue was disconnected.
	Source ReferencedSubject `json:"source"`
	// Issue or pull request which was disconnected.
	Subject ReferencedSubject `json:"subject"`
}

func (DisconnectedEvent) IsPullRequestTimelineItems() {}
func (DisconnectedEvent) IsIssueTimelineItems()       {}
func (DisconnectedEvent) IsNode()                     {}

// Autogenerated input type of DismissPullRequestReview
type DismissPullRequestReviewInput struct {
	// The Node ID of the pull request review to modify.
	PullRequestReviewID string `json:"pullRequestReviewId"`
	// The contents of the pull request review dismissal message.
	Message string `json:"message"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of DismissPullRequestReview
type DismissPullRequestReviewPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The dismissed pull request review.
	PullRequestReview *PullRequestReview `json:"pullRequestReview"`
}

// Specifies a review comment to be left with a Pull Request Review.
type DraftPullRequestReviewComment struct {
	// Path to the file being commented on.
	Path string `json:"path"`
	// Position in the file to leave a comment on.
	Position int `json:"position"`
	// Body of the comment to leave.
	Body string `json:"body"`
}

// Specifies a review comment thread to be left with a Pull Request Review.
type DraftPullRequestReviewThread struct {
	// Path to the file being commented on.
	Path string `json:"path"`
	// The line of the blob to which the thread refers. The end of the line range for multi-line comments.
	Line int `json:"line"`
	// The side of the diff on which the line resides. For multi-line comments, this is the side for the end of the line range.
	Side *DiffSide `json:"side"`
	// The first line of the range to which the comment refers.
	StartLine *int `json:"startLine"`
	// The side of the diff on which the start line resides.
	StartSide *DiffSide `json:"startSide"`
	// Body of the comment to leave.
	Body string `json:"body"`
}

// An account to manage multiple organizations with consolidated policy and billing.
type Enterprise struct {
	// A URL pointing to the enterprise's public avatar.
	AvatarURL string `json:"avatarUrl"`
	// Enterprise billing information visible to enterprise billing managers.
	BillingInfo *EnterpriseBillingInfo `json:"billingInfo"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	// Identifies the primary key from the database.
	DatabaseID *int `json:"databaseId"`
	// The description of the enterprise.
	Description *string `json:"description"`
	// The description of the enterprise as HTML.
	DescriptionHTML string `json:"descriptionHTML"`
	ID              string `json:"id"`
	// The location of the enterprise.
	Location *string `json:"location"`
	// A list of users who are members of this enterprise.
	Members *EnterpriseMemberConnection `json:"members"`
	// The name of the enterprise.
	Name string `json:"name"`
	// A list of organizations that belong to this enterprise.
	Organizations *OrganizationConnection `json:"organizations"`
	// Enterprise information only visible to enterprise owners.
	OwnerInfo *EnterpriseOwnerInfo `json:"ownerInfo"`
	// The HTTP path for this enterprise.
	ResourcePath string `json:"resourcePath"`
	// The URL-friendly identifier for the enterprise.
	Slug string `json:"slug"`
	// The HTTP URL for this enterprise.
	URL string `json:"url"`
	// A list of user accounts on this enterprise.
	UserAccounts *EnterpriseUserAccountConnection `json:"userAccounts"`
	// Is the current viewer an admin of this enterprise?
	ViewerIsAdmin bool `json:"viewerIsAdmin"`
	// The URL of the enterprise website.
	WebsiteURL *string `json:"websiteUrl"`
}

func (Enterprise) IsNode()             {}
func (Enterprise) IsIPAllowListOwner() {}

// The connection type for User.
type EnterpriseAdministratorConnection struct {
	// A list of edges.
	Edges []*EnterpriseAdministratorEdge `json:"edges"`
	// A list of nodes.
	Nodes []*User `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// A User who is an administrator of an enterprise.
type EnterpriseAdministratorEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *User `json:"node"`
	// The role of the administrator.
	Role EnterpriseAdministratorRole `json:"role"`
}

// An invitation for a user to become an owner or billing manager of an enterprise.
type EnterpriseAdministratorInvitation struct {
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	// The email of the person who was invited to the enterprise.
	Email *string `json:"email"`
	// The enterprise the invitation is for.
	Enterprise *Enterprise `json:"enterprise"`
	ID         string      `json:"id"`
	// The user who was invited to the enterprise.
	Invitee *User `json:"invitee"`
	// The user who created the invitation.
	Inviter *User `json:"inviter"`
	// The invitee's pending role in the enterprise (owner or billing_manager).
	Role EnterpriseAdministratorRole `json:"role"`
}

func (EnterpriseAdministratorInvitation) IsNode() {}

// The connection type for EnterpriseAdministratorInvitation.
type EnterpriseAdministratorInvitationConnection struct {
	// A list of edges.
	Edges []*EnterpriseAdministratorInvitationEdge `json:"edges"`
	// A list of nodes.
	Nodes []*EnterpriseAdministratorInvitation `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type EnterpriseAdministratorInvitationEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *EnterpriseAdministratorInvitation `json:"node"`
}

// Ordering options for enterprise administrator invitation connections
type EnterpriseAdministratorInvitationOrder struct {
	// The field to order enterprise administrator invitations by.
	Field EnterpriseAdministratorInvitationOrderField `json:"field"`
	// The ordering direction.
	Direction OrderDirection `json:"direction"`
}

// Enterprise billing information visible to enterprise billing managers and owners.
type EnterpriseBillingInfo struct {
	// The number of licenseable users/emails across the enterprise.
	AllLicensableUsersCount int `json:"allLicensableUsersCount"`
	// The number of data packs used by all organizations owned by the enterprise.
	AssetPacks int `json:"assetPacks"`
	// The number of available seats across all owned organizations based on the unique number of billable users.
	AvailableSeats int `json:"availableSeats"`
	// The bandwidth quota in GB for all organizations owned by the enterprise.
	BandwidthQuota float64 `json:"bandwidthQuota"`
	// The bandwidth usage in GB for all organizations owned by the enterprise.
	BandwidthUsage float64 `json:"bandwidthUsage"`
	// The bandwidth usage as a percentage of the bandwidth quota.
	BandwidthUsagePercentage int `json:"bandwidthUsagePercentage"`
	// The total seats across all organizations owned by the enterprise.
	Seats int `json:"seats"`
	// The storage quota in GB for all organizations owned by the enterprise.
	StorageQuota float64 `json:"storageQuota"`
	// The storage usage in GB for all organizations owned by the enterprise.
	StorageUsage float64 `json:"storageUsage"`
	// The storage usage as a percentage of the storage quota.
	StorageUsagePercentage int `json:"storageUsagePercentage"`
	// The number of available licenses across all owned organizations based on the unique number of billable users.
	TotalAvailableLicenses int `json:"totalAvailableLicenses"`
	// The total number of licenses allocated.
	TotalLicenses int `json:"totalLicenses"`
}

// An identity provider configured to provision identities for an enterprise.
type EnterpriseIdentityProvider struct {
	// The digest algorithm used to sign SAML requests for the identity provider.
	DigestMethod *SamlDigestAlgorithm `json:"digestMethod"`
	// The enterprise this identity provider belongs to.
	Enterprise *Enterprise `json:"enterprise"`
	// ExternalIdentities provisioned by this identity provider.
	ExternalIdentities *ExternalIdentityConnection `json:"externalIdentities"`
	ID                 string                      `json:"id"`
	// The x509 certificate used by the identity provider to sign assertions and responses.
	IdpCertificate *string `json:"idpCertificate"`
	// The Issuer Entity ID for the SAML identity provider.
	Issuer *string `json:"issuer"`
	// Recovery codes that can be used by admins to access the enterprise if the identity provider is unavailable.
	RecoveryCodes []string `json:"recoveryCodes"`
	// The signature algorithm used to sign SAML requests for the identity provider.
	SignatureMethod *SamlSignatureAlgorithm `json:"signatureMethod"`
	// The URL endpoint for the identity provider's SAML SSO.
	SsoURL *string `json:"ssoUrl"`
}

func (EnterpriseIdentityProvider) IsNode() {}

// The connection type for EnterpriseMember.
type EnterpriseMemberConnection struct {
	// A list of edges.
	Edges []*EnterpriseMemberEdge `json:"edges"`
	// A list of nodes.
	Nodes []EnterpriseMember `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// A User who is a member of an enterprise through one or more organizations.
type EnterpriseMemberEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// Whether the user does not have a license for the enterprise.
	IsUnlicensed bool `json:"isUnlicensed"`
	// The item at the end of the edge.
	Node EnterpriseMember `json:"node"`
}

// Ordering options for enterprise member connections.
type EnterpriseMemberOrder struct {
	// The field to order enterprise members by.
	Field EnterpriseMemberOrderField `json:"field"`
	// The ordering direction.
	Direction OrderDirection `json:"direction"`
}

// The connection type for Organization.
type EnterpriseOrganizationMembershipConnection struct {
	// A list of edges.
	Edges []*EnterpriseOrganizationMembershipEdge `json:"edges"`
	// A list of nodes.
	Nodes []*Organization `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An enterprise organization that a user is a member of.
type EnterpriseOrganizationMembershipEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *Organization `json:"node"`
	// The role of the user in the enterprise membership.
	Role EnterpriseUserAccountMembershipRole `json:"role"`
}

// The connection type for User.
type EnterpriseOutsideCollaboratorConnection struct {
	// A list of edges.
	Edges []*EnterpriseOutsideCollaboratorEdge `json:"edges"`
	// A list of nodes.
	Nodes []*User `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// A User who is an outside collaborator of an enterprise through one or more organizations.
type EnterpriseOutsideCollaboratorEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// Whether the outside collaborator does not have a license for the enterprise.
	IsUnlicensed bool `json:"isUnlicensed"`
	// The item at the end of the edge.
	Node *User `json:"node"`
	// The enterprise organization repositories this user is a member of.
	Repositories *EnterpriseRepositoryInfoConnection `json:"repositories"`
}

// Enterprise information only visible to enterprise owners.
type EnterpriseOwnerInfo struct {
	// A list of enterprise organizations configured with the provided action execution capabilities setting value.
	ActionExecutionCapabilitySettingOrganizations *OrganizationConnection `json:"actionExecutionCapabilitySettingOrganizations"`
	// A list of all of the administrators for this enterprise.
	Admins *EnterpriseAdministratorConnection `json:"admins"`
	// A list of users in the enterprise who currently have two-factor authentication disabled.
	AffiliatedUsersWithTwoFactorDisabled *UserConnection `json:"affiliatedUsersWithTwoFactorDisabled"`
	// Whether or not affiliated users with two-factor authentication disabled exist in the enterprise.
	AffiliatedUsersWithTwoFactorDisabledExist bool `json:"affiliatedUsersWithTwoFactorDisabledExist"`
	// The setting value for whether private repository forking is enabled for repositories in organizations in this enterprise.
	AllowPrivateRepositoryForkingSetting EnterpriseEnabledDisabledSettingValue `json:"allowPrivateRepositoryForkingSetting"`
	// A list of enterprise organizations configured with the provided private repository forking setting value.
	AllowPrivateRepositoryForkingSettingOrganizations *OrganizationConnection `json:"allowPrivateRepositoryForkingSettingOrganizations"`
	// The setting value for base repository permissions for organizations in this enterprise.
	DefaultRepositoryPermissionSetting EnterpriseDefaultRepositoryPermissionSettingValue `json:"defaultRepositoryPermissionSetting"`
	// A list of enterprise organizations configured with the provided default repository permission.
	DefaultRepositoryPermissionSettingOrganizations *OrganizationConnection `json:"defaultRepositoryPermissionSettingOrganizations"`
	// Enterprise Server installations owned by the enterprise.
	EnterpriseServerInstallations *EnterpriseServerInstallationConnection `json:"enterpriseServerInstallations"`
	// The setting value for whether the enterprise has an IP allow list enabled.
	IPAllowListEnabledSetting IPAllowListEnabledSettingValue `json:"ipAllowListEnabledSetting"`
	// The IP addresses that are allowed to access resources owned by the enterprise.
	IPAllowListEntries *IPAllowListEntryConnection `json:"ipAllowListEntries"`
	// Whether or not the default repository permission is currently being updated.
	IsUpdatingDefaultRepositoryPermission bool `json:"isUpdatingDefaultRepositoryPermission"`
	// Whether the two-factor authentication requirement is currently being enforced.
	IsUpdatingTwoFactorRequirement bool `json:"isUpdatingTwoFactorRequirement"`
	// The setting value for whether organization members with admin permissions on a repository can change repository visibility.
	MembersCanChangeRepositoryVisibilitySetting EnterpriseEnabledDisabledSettingValue `json:"membersCanChangeRepositoryVisibilitySetting"`
	// A list of enterprise organizations configured with the provided can change repository visibility setting value.
	MembersCanChangeRepositoryVisibilitySettingOrganizations *OrganizationConnection `json:"membersCanChangeRepositoryVisibilitySettingOrganizations"`
	// The setting value for whether members of organizations in the enterprise can create internal repositories.
	MembersCanCreateInternalRepositoriesSetting *bool `json:"membersCanCreateInternalRepositoriesSetting"`
	// The setting value for whether members of organizations in the enterprise can create private repositories.
	MembersCanCreatePrivateRepositoriesSetting *bool `json:"membersCanCreatePrivateRepositoriesSetting"`
	// The setting value for whether members of organizations in the enterprise can create public repositories.
	MembersCanCreatePublicRepositoriesSetting *bool `json:"membersCanCreatePublicRepositoriesSetting"`
	// The setting value for whether members of organizations in the enterprise can create repositories.
	MembersCanCreateRepositoriesSetting *EnterpriseMembersCanCreateRepositoriesSettingValue `json:"membersCanCreateRepositoriesSetting"`
	// A list of enterprise organizations configured with the provided repository creation setting value.
	MembersCanCreateRepositoriesSettingOrganizations *OrganizationConnection `json:"membersCanCreateRepositoriesSettingOrganizations"`
	// The setting value for whether members with admin permissions for repositories can delete issues.
	MembersCanDeleteIssuesSetting EnterpriseEnabledDisabledSettingValue `json:"membersCanDeleteIssuesSetting"`
	// A list of enterprise organizations configured with the provided members can delete issues setting value.
	MembersCanDeleteIssuesSettingOrganizations *OrganizationConnection `json:"membersCanDeleteIssuesSettingOrganizations"`
	// The setting value for whether members with admin permissions for repositories can delete or transfer repositories.
	MembersCanDeleteRepositoriesSetting EnterpriseEnabledDisabledSettingValue `json:"membersCanDeleteRepositoriesSetting"`
	// A list of enterprise organizations configured with the provided members can delete repositories setting value.
	MembersCanDeleteRepositoriesSettingOrganizations *OrganizationConnection `json:"membersCanDeleteRepositoriesSettingOrganizations"`
	// The setting value for whether members of organizations in the enterprise can invite outside collaborators.
	MembersCanInviteCollaboratorsSetting EnterpriseEnabledDisabledSettingValue `json:"membersCanInviteCollaboratorsSetting"`
	// A list of enterprise organizations configured with the provided members can invite collaborators setting value.
	MembersCanInviteCollaboratorsSettingOrganizations *OrganizationConnection `json:"membersCanInviteCollaboratorsSettingOrganizations"`
	// Indicates whether members of this enterprise's organizations can purchase additional services for those organizations.
	MembersCanMakePurchasesSetting EnterpriseMembersCanMakePurchasesSettingValue `json:"membersCanMakePurchasesSetting"`
	// The setting value for whether members with admin permissions for repositories can update protected branches.
	MembersCanUpdateProtectedBranchesSetting EnterpriseEnabledDisabledSettingValue `json:"membersCanUpdateProtectedBranchesSetting"`
	// A list of enterprise organizations configured with the provided members can update protected branches setting value.
	MembersCanUpdateProtectedBranchesSettingOrganizations *OrganizationConnection `json:"membersCanUpdateProtectedBranchesSettingOrganizations"`
	// The setting value for whether members can view dependency insights.
	MembersCanViewDependencyInsightsSetting EnterpriseEnabledDisabledSettingValue `json:"membersCanViewDependencyInsightsSetting"`
	// A list of enterprise organizations configured with the provided members can view dependency insights setting value.
	MembersCanViewDependencyInsightsSettingOrganizations *OrganizationConnection `json:"membersCanViewDependencyInsightsSettingOrganizations"`
	// The setting value for whether organization projects are enabled for organizations in this enterprise.
	OrganizationProjectsSetting EnterpriseEnabledDisabledSettingValue `json:"organizationProjectsSetting"`
	// A list of enterprise organizations configured with the provided organization projects setting value.
	OrganizationProjectsSettingOrganizations *OrganizationConnection `json:"organizationProjectsSettingOrganizations"`
	// A list of outside collaborators across the repositories in the enterprise.
	OutsideCollaborators *EnterpriseOutsideCollaboratorConnection `json:"outsideCollaborators"`
	// A list of pending administrator invitations for the enterprise.
	PendingAdminInvitations *EnterpriseAdministratorInvitationConnection `json:"pendingAdminInvitations"`
	// A list of pending collaborator invitations across the repositories in the enterprise.
	PendingCollaboratorInvitations *RepositoryInvitationConnection `json:"pendingCollaboratorInvitations"`
	// A list of pending collaborators across the repositories in the enterprise.
	PendingCollaborators *EnterprisePendingCollaboratorConnection `json:"pendingCollaborators"`
	// A list of pending member invitations for organizations in the enterprise.
	PendingMemberInvitations *EnterprisePendingMemberInvitationConnection `json:"pendingMemberInvitations"`
	// The setting value for whether repository projects are enabled in this enterprise.
	RepositoryProjectsSetting EnterpriseEnabledDisabledSettingValue `json:"repositoryProjectsSetting"`
	// A list of enterprise organizations configured with the provided repository projects setting value.
	RepositoryProjectsSettingOrganizations *OrganizationConnection `json:"repositoryProjectsSettingOrganizations"`
	// The SAML Identity Provider for the enterprise.
	SamlIdentityProvider *EnterpriseIdentityProvider `json:"samlIdentityProvider"`
	// A list of enterprise organizations configured with the SAML single sign-on setting value.
	SamlIdentityProviderSettingOrganizations *OrganizationConnection `json:"samlIdentityProviderSettingOrganizations"`
	// The setting value for whether team discussions are enabled for organizations in this enterprise.
	TeamDiscussionsSetting EnterpriseEnabledDisabledSettingValue `json:"teamDiscussionsSetting"`
	// A list of enterprise organizations configured with the provided team discussions setting value.
	TeamDiscussionsSettingOrganizations *OrganizationConnection `json:"teamDiscussionsSettingOrganizations"`
	// The setting value for whether the enterprise requires two-factor authentication for its organizations and users.
	TwoFactorRequiredSetting EnterpriseEnabledSettingValue `json:"twoFactorRequiredSetting"`
	// A list of enterprise organizations configured with the two-factor authentication setting value.
	TwoFactorRequiredSettingOrganizations *OrganizationConnection `json:"twoFactorRequiredSettingOrganizations"`
}

// The connection type for User.
type EnterprisePendingCollaboratorConnection struct {
	// A list of edges.
	Edges []*EnterprisePendingCollaboratorEdge `json:"edges"`
	// A list of nodes.
	Nodes []*User `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// A user with an invitation to be a collaborator on a repository owned by an organization in an enterprise.
type EnterprisePendingCollaboratorEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// Whether the invited collaborator does not have a license for the enterprise.
	IsUnlicensed bool `json:"isUnlicensed"`
	// The item at the end of the edge.
	Node *User `json:"node"`
	// The enterprise organization repositories this user is a member of.
	Repositories *EnterpriseRepositoryInfoConnection `json:"repositories"`
}

// The connection type for OrganizationInvitation.
type EnterprisePendingMemberInvitationConnection struct {
	// A list of edges.
	Edges []*EnterprisePendingMemberInvitationEdge `json:"edges"`
	// A list of nodes.
	Nodes []*OrganizationInvitation `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
	// Identifies the total count of unique users in the connection.
	TotalUniqueUserCount int `json:"totalUniqueUserCount"`
}

// An invitation to be a member in an enterprise organization.
type EnterprisePendingMemberInvitationEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// Whether the invitation has a license for the enterprise.
	IsUnlicensed bool `json:"isUnlicensed"`
	// The item at the end of the edge.
	Node *OrganizationInvitation `json:"node"`
}

// A subset of repository information queryable from an enterprise.
type EnterpriseRepositoryInfo struct {
	ID string `json:"id"`
	// Identifies if the repository is private.
	IsPrivate bool `json:"isPrivate"`
	// The repository's name.
	Name string `json:"name"`
	// The repository's name with owner.
	NameWithOwner string `json:"nameWithOwner"`
}

func (EnterpriseRepositoryInfo) IsNode() {}

// The connection type for EnterpriseRepositoryInfo.
type EnterpriseRepositoryInfoConnection struct {
	// A list of edges.
	Edges []*EnterpriseRepositoryInfoEdge `json:"edges"`
	// A list of nodes.
	Nodes []*EnterpriseRepositoryInfo `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type EnterpriseRepositoryInfoEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *EnterpriseRepositoryInfo `json:"node"`
}

// An Enterprise Server installation.
type EnterpriseServerInstallation struct {
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	// The customer name to which the Enterprise Server installation belongs.
	CustomerName string `json:"customerName"`
	// The host name of the Enterprise Server installation.
	HostName string `json:"hostName"`
	ID       string `json:"id"`
	// Whether or not the installation is connected to an Enterprise Server installation via GitHub Connect.
	IsConnected bool `json:"isConnected"`
	// Identifies the date and time when the object was last updated.
	UpdatedAt string `json:"updatedAt"`
	// User accounts on this Enterprise Server installation.
	UserAccounts *EnterpriseServerUserAccountConnection `json:"userAccounts"`
	// User accounts uploads for the Enterprise Server installation.
	UserAccountsUploads *EnterpriseServerUserAccountsUploadConnection `json:"userAccountsUploads"`
}

func (EnterpriseServerInstallation) IsNode() {}

// The connection type for EnterpriseServerInstallation.
type EnterpriseServerInstallationConnection struct {
	// A list of edges.
	Edges []*EnterpriseServerInstallationEdge `json:"edges"`
	// A list of nodes.
	Nodes []*EnterpriseServerInstallation `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type EnterpriseServerInstallationEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *EnterpriseServerInstallation `json:"node"`
}

// Ordering options for Enterprise Server installation connections.
type EnterpriseServerInstallationOrder struct {
	// The field to order Enterprise Server installations by.
	Field EnterpriseServerInstallationOrderField `json:"field"`
	// The ordering direction.
	Direction OrderDirection `json:"direction"`
}

// A user account on an Enterprise Server installation.
type EnterpriseServerUserAccount struct {
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	// User emails belonging to this user account.
	Emails *EnterpriseServerUserAccountEmailConnection `json:"emails"`
	// The Enterprise Server installation on which this user account exists.
	EnterpriseServerInstallation *EnterpriseServerInstallation `json:"enterpriseServerInstallation"`
	ID                           string                        `json:"id"`
	// Whether the user account is a site administrator on the Enterprise Server installation.
	IsSiteAdmin bool `json:"isSiteAdmin"`
	// The login of the user account on the Enterprise Server installation.
	Login string `json:"login"`
	// The profile name of the user account on the Enterprise Server installation.
	ProfileName *string `json:"profileName"`
	// The date and time when the user account was created on the Enterprise Server installation.
	RemoteCreatedAt string `json:"remoteCreatedAt"`
	// The ID of the user account on the Enterprise Server installation.
	RemoteUserID int `json:"remoteUserId"`
	// Identifies the date and time when the object was last updated.
	UpdatedAt string `json:"updatedAt"`
}

func (EnterpriseServerUserAccount) IsNode() {}

// The connection type for EnterpriseServerUserAccount.
type EnterpriseServerUserAccountConnection struct {
	// A list of edges.
	Edges []*EnterpriseServerUserAccountEdge `json:"edges"`
	// A list of nodes.
	Nodes []*EnterpriseServerUserAccount `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type EnterpriseServerUserAccountEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *EnterpriseServerUserAccount `json:"node"`
}

// An email belonging to a user account on an Enterprise Server installation.
type EnterpriseServerUserAccountEmail struct {
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	// The email address.
	Email string `json:"email"`
	ID    string `json:"id"`
	// Indicates whether this is the primary email of the associated user account.
	IsPrimary bool `json:"isPrimary"`
	// Identifies the date and time when the object was last updated.
	UpdatedAt string `json:"updatedAt"`
	// The user account to which the email belongs.
	UserAccount *EnterpriseServerUserAccount `json:"userAccount"`
}

func (EnterpriseServerUserAccountEmail) IsNode() {}

// The connection type for EnterpriseServerUserAccountEmail.
type EnterpriseServerUserAccountEmailConnection struct {
	// A list of edges.
	Edges []*EnterpriseServerUserAccountEmailEdge `json:"edges"`
	// A list of nodes.
	Nodes []*EnterpriseServerUserAccountEmail `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type EnterpriseServerUserAccountEmailEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *EnterpriseServerUserAccountEmail `json:"node"`
}

// Ordering options for Enterprise Server user account email connections.
type EnterpriseServerUserAccountEmailOrder struct {
	// The field to order emails by.
	Field EnterpriseServerUserAccountEmailOrderField `json:"field"`
	// The ordering direction.
	Direction OrderDirection `json:"direction"`
}

// Ordering options for Enterprise Server user account connections.
type EnterpriseServerUserAccountOrder struct {
	// The field to order user accounts by.
	Field EnterpriseServerUserAccountOrderField `json:"field"`
	// The ordering direction.
	Direction OrderDirection `json:"direction"`
}

// A user accounts upload from an Enterprise Server installation.
type EnterpriseServerUserAccountsUpload struct {
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	// The enterprise to which this upload belongs.
	Enterprise *Enterprise `json:"enterprise"`
	// The Enterprise Server installation for which this upload was generated.
	EnterpriseServerInstallation *EnterpriseServerInstallation `json:"enterpriseServerInstallation"`
	ID                           string                        `json:"id"`
	// The name of the file uploaded.
	Name string `json:"name"`
	// The synchronization state of the upload
	SyncState EnterpriseServerUserAccountsUploadSyncState `json:"syncState"`
	// Identifies the date and time when the object was last updated.
	UpdatedAt string `json:"updatedAt"`
}

func (EnterpriseServerUserAccountsUpload) IsNode() {}

// The connection type for EnterpriseServerUserAccountsUpload.
type EnterpriseServerUserAccountsUploadConnection struct {
	// A list of edges.
	Edges []*EnterpriseServerUserAccountsUploadEdge `json:"edges"`
	// A list of nodes.
	Nodes []*EnterpriseServerUserAccountsUpload `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type EnterpriseServerUserAccountsUploadEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *EnterpriseServerUserAccountsUpload `json:"node"`
}

// Ordering options for Enterprise Server user accounts upload connections.
type EnterpriseServerUserAccountsUploadOrder struct {
	// The field to order user accounts uploads by.
	Field EnterpriseServerUserAccountsUploadOrderField `json:"field"`
	// The ordering direction.
	Direction OrderDirection `json:"direction"`
}

// An account for a user who is an admin of an enterprise or a member of an enterprise through one or more organizations.
type EnterpriseUserAccount struct {
	// A URL pointing to the enterprise user account's public avatar.
	AvatarURL string `json:"avatarUrl"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	// The enterprise in which this user account exists.
	Enterprise *Enterprise `json:"enterprise"`
	ID         string      `json:"id"`
	// An identifier for the enterprise user account, a login or email address
	Login string `json:"login"`
	// The name of the enterprise user account
	Name *string `json:"name"`
	// A list of enterprise organizations this user is a member of.
	Organizations *EnterpriseOrganizationMembershipConnection `json:"organizations"`
	// The HTTP path for this user.
	ResourcePath string `json:"resourcePath"`
	// Identifies the date and time when the object was last updated.
	UpdatedAt string `json:"updatedAt"`
	// The HTTP URL for this user.
	URL string `json:"url"`
	// The user within the enterprise.
	User *User `json:"user"`
}

func (EnterpriseUserAccount) IsEnterpriseMember() {}
func (EnterpriseUserAccount) IsNode()             {}
func (EnterpriseUserAccount) IsActor()            {}

// The connection type for EnterpriseUserAccount.
type EnterpriseUserAccountConnection struct {
	// A list of edges.
	Edges []*EnterpriseUserAccountEdge `json:"edges"`
	// A list of nodes.
	Nodes []*EnterpriseUserAccount `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type EnterpriseUserAccountEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *EnterpriseUserAccount `json:"node"`
}

// An external identity provisioned by SAML SSO or SCIM.
type ExternalIdentity struct {
	// The GUID for this identity
	GUID string `json:"guid"`
	ID   string `json:"id"`
	// Organization invitation for this SCIM-provisioned external identity
	OrganizationInvitation *OrganizationInvitation `json:"organizationInvitation"`
	// SAML Identity attributes
	SamlIdentity *ExternalIdentitySamlAttributes `json:"samlIdentity"`
	// SCIM Identity attributes
	ScimIdentity *ExternalIdentityScimAttributes `json:"scimIdentity"`
	// User linked to this external identity. Will be NULL if this identity has not been claimed by an organization member.
	User *User `json:"user"`
}

func (ExternalIdentity) IsNode() {}

// The connection type for ExternalIdentity.
type ExternalIdentityConnection struct {
	// A list of edges.
	Edges []*ExternalIdentityEdge `json:"edges"`
	// A list of nodes.
	Nodes []*ExternalIdentity `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type ExternalIdentityEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *ExternalIdentity `json:"node"`
}

// SAML attributes for the External Identity
type ExternalIdentitySamlAttributes struct {
	// The emails associated with the SAML identity
	Emails []*UserEmailMetadata `json:"emails"`
	// Family name of the SAML identity
	FamilyName *string `json:"familyName"`
	// Given name of the SAML identity
	GivenName *string `json:"givenName"`
	// The groups linked to this identity in IDP
	Groups []string `json:"groups"`
	// The NameID of the SAML identity
	NameID *string `json:"nameId"`
	// The userName of the SAML identity
	Username *string `json:"username"`
}

// SCIM attributes for the External Identity
type ExternalIdentityScimAttributes struct {
	// The emails associated with the SCIM identity
	Emails []*UserEmailMetadata `json:"emails"`
	// Family name of the SCIM identity
	FamilyName *string `json:"familyName"`
	// Given name of the SCIM identity
	GivenName *string `json:"givenName"`
	// The groups linked to this identity in IDP
	Groups []string `json:"groups"`
	// The userName of the SCIM identity
	Username *string `json:"username"`
}

// Autogenerated input type of FollowUser
type FollowUserInput struct {
	// ID of the user to follow.
	UserID string `json:"userId"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of FollowUser
type FollowUserPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The user that was followed.
	User *User `json:"user"`
}

// The connection type for User.
type FollowerConnection struct {
	// A list of edges.
	Edges []*UserEdge `json:"edges"`
	// A list of nodes.
	Nodes []*User `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// The connection type for User.
type FollowingConnection struct {
	// A list of edges.
	Edges []*UserEdge `json:"edges"`
	// A list of nodes.
	Nodes []*User `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// A funding platform link for a repository.
type FundingLink struct {
	// The funding platform this link is for.
	Platform FundingPlatform `json:"platform"`
	// The configured URL for this funding link.
	URL string `json:"url"`
}

// A generic hovercard context with a message and icon
type GenericHovercardContext struct {
	// A string describing this context
	Message string `json:"message"`
	// An octicon to accompany this context
	Octicon string `json:"octicon"`
}

func (GenericHovercardContext) IsHovercardContext() {}

// A Gist.
type Gist struct {
	// A list of comments associated with the gist
	Comments *GistCommentConnection `json:"comments"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	// The gist description.
	Description *string `json:"description"`
	// The files in this gist.
	Files []*GistFile `json:"files"`
	// A list of forks associated with the gist
	Forks *GistConnection `json:"forks"`
	ID    string          `json:"id"`
	// Identifies if the gist is a fork.
	IsFork bool `json:"isFork"`
	// Whether the gist is public or not.
	IsPublic bool `json:"isPublic"`
	// The gist name.
	Name string `json:"name"`
	// The gist owner.
	Owner RepositoryOwner `json:"owner"`
	// Identifies when the gist was last pushed to.
	PushedAt *string `json:"pushedAt"`
	// The HTML path to this resource.
	ResourcePath string `json:"resourcePath"`
	// Returns a count of how many stargazers there are on this object
	//
	StargazerCount int `json:"stargazerCount"`
	// A list of users who have starred this starrable.
	Stargazers *StargazerConnection `json:"stargazers"`
	// Identifies the date and time when the object was last updated.
	UpdatedAt string `json:"updatedAt"`
	// The HTTP URL for this Gist.
	URL string `json:"url"`
	// Returns a boolean indicating whether the viewing user has starred this starrable.
	ViewerHasStarred bool `json:"viewerHasStarred"`
}

func (Gist) IsNode()                     {}
func (Gist) IsStarrable()                {}
func (Gist) IsUniformResourceLocatable() {}
func (Gist) IsPinnableItem()             {}

// Represents a comment on an Gist.
type GistComment struct {
	// The actor who authored the comment.
	Author Actor `json:"author"`
	// Author's association with the gist.
	AuthorAssociation CommentAuthorAssociation `json:"authorAssociation"`
	// Identifies the comment body.
	Body string `json:"body"`
	// The body rendered to HTML.
	BodyHTML string `json:"bodyHTML"`
	// The body rendered to text.
	BodyText string `json:"bodyText"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	// Check if this comment was created via an email reply.
	CreatedViaEmail bool `json:"createdViaEmail"`
	// Identifies the primary key from the database.
	DatabaseID *int `json:"databaseId"`
	// The actor who edited the comment.
	Editor Actor `json:"editor"`
	// The associated gist.
	Gist *Gist  `json:"gist"`
	ID   string `json:"id"`
	// Check if this comment was edited and includes an edit with the creation data
	IncludesCreatedEdit bool `json:"includesCreatedEdit"`
	// Returns whether or not a comment has been minimized.
	IsMinimized bool `json:"isMinimized"`
	// The moment the editor made the last edit
	LastEditedAt *string `json:"lastEditedAt"`
	// Returns why the comment was minimized.
	MinimizedReason *string `json:"minimizedReason"`
	// Identifies when the comment was published at.
	PublishedAt *string `json:"publishedAt"`
	// Identifies the date and time when the object was last updated.
	UpdatedAt string `json:"updatedAt"`
	// A list of edits to this content.
	UserContentEdits *UserContentEditConnection `json:"userContentEdits"`
	// Check if the current viewer can delete this object.
	ViewerCanDelete bool `json:"viewerCanDelete"`
	// Check if the current viewer can minimize this object.
	ViewerCanMinimize bool `json:"viewerCanMinimize"`
	// Check if the current viewer can update this object.
	ViewerCanUpdate bool `json:"viewerCanUpdate"`
	// Reasons why the current viewer can not update this comment.
	ViewerCannotUpdateReasons []CommentCannotUpdateReason `json:"viewerCannotUpdateReasons"`
	// Did the viewer author this comment.
	ViewerDidAuthor bool `json:"viewerDidAuthor"`
}

func (GistComment) IsNode()             {}
func (GistComment) IsComment()          {}
func (GistComment) IsDeletable()        {}
func (GistComment) IsMinimizable()      {}
func (GistComment) IsUpdatable()        {}
func (GistComment) IsUpdatableComment() {}

// The connection type for GistComment.
type GistCommentConnection struct {
	// A list of edges.
	Edges []*GistCommentEdge `json:"edges"`
	// A list of nodes.
	Nodes []*GistComment `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type GistCommentEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *GistComment `json:"node"`
}

// The connection type for Gist.
type GistConnection struct {
	// A list of edges.
	Edges []*GistEdge `json:"edges"`
	// A list of nodes.
	Nodes []*Gist `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type GistEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *Gist `json:"node"`
}

// A file in a gist.
type GistFile struct {
	// The file name encoded to remove characters that are invalid in URL paths.
	EncodedName *string `json:"encodedName"`
	// The gist file encoding.
	Encoding *string `json:"encoding"`
	// The file extension from the file name.
	Extension *string `json:"extension"`
	// Indicates if this file is an image.
	IsImage bool `json:"isImage"`
	// Whether the file's contents were truncated.
	IsTruncated bool `json:"isTruncated"`
	// The programming language this file is written in.
	Language *Language `json:"language"`
	// The gist file name.
	Name *string `json:"name"`
	// The gist file size in bytes.
	Size *int `json:"size"`
	// UTF8 text data or null if the file is binary
	Text *string `json:"text"`
}

// Ordering options for gist connections
type GistOrder struct {
	// The field to order repositories by.
	Field GistOrderField `json:"field"`
	// The ordering direction.
	Direction OrderDirection `json:"direction"`
}

// Represents an actor in a Git commit (ie. an author or committer).
type GitActor struct {
	// A URL pointing to the author's public avatar.
	AvatarURL string `json:"avatarUrl"`
	// The timestamp of the Git action (authoring or committing).
	Date *string `json:"date"`
	// The email in the Git commit.
	Email *string `json:"email"`
	// The name in the Git commit.
	Name *string `json:"name"`
	// The GitHub user corresponding to the email field. Null if no such user exists.
	User *User `json:"user"`
}

// The connection type for GitActor.
type GitActorConnection struct {
	// A list of edges.
	Edges []*GitActorEdge `json:"edges"`
	// A list of nodes.
	Nodes []*GitActor `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type GitActorEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *GitActor `json:"node"`
}

// Represents information about the GitHub instance.
type GitHubMetadata struct {
	// Returns a String that's a SHA of `github-services`
	GitHubServicesSha string `json:"gitHubServicesSha"`
	// IP addresses that users connect to for git operations
	GitIPAddresses []string `json:"gitIpAddresses"`
	// IP addresses that service hooks are sent from
	HookIPAddresses []string `json:"hookIpAddresses"`
	// IP addresses that the importer connects from
	ImporterIPAddresses []string `json:"importerIpAddresses"`
	// Whether or not users are verified
	IsPasswordAuthenticationVerifiable bool `json:"isPasswordAuthenticationVerifiable"`
	// IP addresses for GitHub Pages' A records
	PagesIPAddresses []string `json:"pagesIpAddresses"`
}

// Represents a GPG signature on a Commit or Tag.
type GpgSignature struct {
	// Email used to sign this object.
	Email string `json:"email"`
	// True if the signature is valid and verified by GitHub.
	IsValid bool `json:"isValid"`
	// Hex-encoded ID of the key that signed this object.
	KeyID *string `json:"keyId"`
	// Payload for GPG signing object. Raw ODB object without the signature header.
	Payload string `json:"payload"`
	// ASCII-armored signature header from object.
	Signature string `json:"signature"`
	// GitHub user corresponding to the email signing this commit.
	Signer *User `json:"signer"`
	// The state of this signature. `VALID` if signature is valid and verified by GitHub, otherwise represents reason why signature is considered invalid.
	State GitSignatureState `json:"state"`
	// True if the signature was made with GitHub's signing key.
	WasSignedByGitHub bool `json:"wasSignedByGitHub"`
}

func (GpgSignature) IsGitSignature() {}

// Represents a 'head_ref_deleted' event on a given pull request.
type HeadRefDeletedEvent struct {
	// Identifies the actor who performed the event.
	Actor Actor `json:"actor"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	// Identifies the Ref associated with the `head_ref_deleted` event.
	HeadRef *Ref `json:"headRef"`
	// Identifies the name of the Ref associated with the `head_ref_deleted` event.
	HeadRefName string `json:"headRefName"`
	ID          string `json:"id"`
	// PullRequest referenced by event.
	PullRequest *PullRequest `json:"pullRequest"`
}

func (HeadRefDeletedEvent) IsPullRequestTimelineItems() {}
func (HeadRefDeletedEvent) IsPullRequestTimelineItem()  {}
func (HeadRefDeletedEvent) IsNode()                     {}

// Represents a 'head_ref_force_pushed' event on a given pull request.
type HeadRefForcePushedEvent struct {
	// Identifies the actor who performed the event.
	Actor Actor `json:"actor"`
	// Identifies the after commit SHA for the 'head_ref_force_pushed' event.
	AfterCommit *Commit `json:"afterCommit"`
	// Identifies the before commit SHA for the 'head_ref_force_pushed' event.
	BeforeCommit *Commit `json:"beforeCommit"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// PullRequest referenced by event.
	PullRequest *PullRequest `json:"pullRequest"`
	// Identifies the fully qualified ref name for the 'head_ref_force_pushed' event.
	Ref *Ref `json:"ref"`
}

func (HeadRefForcePushedEvent) IsPullRequestTimelineItems() {}
func (HeadRefForcePushedEvent) IsPullRequestTimelineItem()  {}
func (HeadRefForcePushedEvent) IsNode()                     {}

// Represents a 'head_ref_restored' event on a given pull request.
type HeadRefRestoredEvent struct {
	// Identifies the actor who performed the event.
	Actor Actor `json:"actor"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// PullRequest referenced by event.
	PullRequest *PullRequest `json:"pullRequest"`
}

func (HeadRefRestoredEvent) IsPullRequestTimelineItems() {}
func (HeadRefRestoredEvent) IsNode()                     {}
func (HeadRefRestoredEvent) IsPullRequestTimelineItem()  {}

// Detail needed to display a hovercard for a user
type Hovercard struct {
	// Each of the contexts for this hovercard
	Contexts []HovercardContext `json:"contexts"`
}

// Autogenerated input type of InviteEnterpriseAdmin
type InviteEnterpriseAdminInput struct {
	// The ID of the enterprise to which you want to invite an administrator.
	EnterpriseID string `json:"enterpriseId"`
	// The login of a user to invite as an administrator.
	Invitee *string `json:"invitee"`
	// The email of the person to invite as an administrator.
	Email *string `json:"email"`
	// The role of the administrator.
	Role *EnterpriseAdministratorRole `json:"role"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of InviteEnterpriseAdmin
type InviteEnterpriseAdminPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The created enterprise administrator invitation.
	Invitation *EnterpriseAdministratorInvitation `json:"invitation"`
}

// An IP address or range of addresses that is allowed to access an owner's resources.
type IPAllowListEntry struct {
	// A single IP address or range of IP addresses in CIDR notation.
	AllowListValue string `json:"allowListValue"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// Whether the entry is currently active.
	IsActive bool `json:"isActive"`
	// The name of the IP allow list entry.
	Name *string `json:"name"`
	// The owner of the IP allow list entry.
	Owner IPAllowListOwner `json:"owner"`
	// Identifies the date and time when the object was last updated.
	UpdatedAt string `json:"updatedAt"`
}

func (IPAllowListEntry) IsNode() {}

// The connection type for IpAllowListEntry.
type IPAllowListEntryConnection struct {
	// A list of edges.
	Edges []*IPAllowListEntryEdge `json:"edges"`
	// A list of nodes.
	Nodes []*IPAllowListEntry `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type IPAllowListEntryEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *IPAllowListEntry `json:"node"`
}

// Ordering options for IP allow list entry connections.
type IPAllowListEntryOrder struct {
	// The field to order IP allow list entries by.
	Field IPAllowListEntryOrderField `json:"field"`
	// The ordering direction.
	Direction OrderDirection `json:"direction"`
}

// An Issue is a place to discuss ideas, enhancements, tasks, and bugs for a project.
type Issue struct {
	// Reason that the conversation was locked.
	ActiveLockReason *LockReason `json:"activeLockReason"`
	// A list of Users assigned to this object.
	Assignees *UserConnection `json:"assignees"`
	// The actor who authored the comment.
	Author Actor `json:"author"`
	// Author's association with the subject of the comment.
	AuthorAssociation CommentAuthorAssociation `json:"authorAssociation"`
	// Identifies the body of the issue.
	Body string `json:"body"`
	// The body rendered to HTML.
	BodyHTML string `json:"bodyHTML"`
	// The http path for this issue body
	BodyResourcePath string `json:"bodyResourcePath"`
	// Identifies the body of the issue rendered to text.
	BodyText string `json:"bodyText"`
	// The http URL for this issue body
	BodyURL string `json:"bodyUrl"`
	// `true` if the object is closed (definition of closed may depend on type)
	Closed bool `json:"closed"`
	// Identifies the date and time when the object was closed.
	ClosedAt *string `json:"closedAt"`
	// A list of comments associated with the Issue.
	Comments *IssueCommentConnection `json:"comments"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	// Check if this comment was created via an email reply.
	CreatedViaEmail bool `json:"createdViaEmail"`
	// Identifies the primary key from the database.
	DatabaseID *int `json:"databaseId"`
	// The actor who edited the comment.
	Editor Actor `json:"editor"`
	// The hovercard information for this issue
	Hovercard *Hovercard `json:"hovercard"`
	ID        string     `json:"id"`
	// Check if this comment was edited and includes an edit with the creation data
	IncludesCreatedEdit bool `json:"includesCreatedEdit"`
	// Is this issue read by the viewer
	IsReadByViewer *bool `json:"isReadByViewer"`
	// A list of labels associated with the object.
	Labels *LabelConnection `json:"labels"`
	// The moment the editor made the last edit
	LastEditedAt *string `json:"lastEditedAt"`
	// `true` if the object is locked
	Locked bool `json:"locked"`
	// Identifies the milestone associated with the issue.
	Milestone *Milestone `json:"milestone"`
	// Identifies the issue number.
	Number int `json:"number"`
	// A list of Users that are participating in the Issue conversation.
	Participants *UserConnection `json:"participants"`
	// List of project cards associated with this issue.
	ProjectCards *ProjectCardConnection `json:"projectCards"`
	// Identifies when the comment was published at.
	PublishedAt *string `json:"publishedAt"`
	// A list of reactions grouped by content left on the subject.
	ReactionGroups []*ReactionGroup `json:"reactionGroups"`
	// A list of Reactions left on the Issue.
	Reactions *ReactionConnection `json:"reactions"`
	// The repository associated with this node.
	Repository *Repository `json:"repository"`
	// The HTTP path for this issue
	ResourcePath string `json:"resourcePath"`
	// Identifies the state of the issue.
	State IssueState `json:"state"`
	// A list of events, comments, commits, etc. associated with the issue.
	Timeline *IssueTimelineConnection `json:"timeline"`
	// A list of events, comments, commits, etc. associated with the issue.
	TimelineItems *IssueTimelineItemsConnection `json:"timelineItems"`
	// Identifies the issue title.
	Title string `json:"title"`
	// Identifies the date and time when the object was last updated.
	UpdatedAt string `json:"updatedAt"`
	// The HTTP URL for this issue
	URL string `json:"url"`
	// A list of edits to this content.
	UserContentEdits *UserContentEditConnection `json:"userContentEdits"`
	// Can user react to this subject
	ViewerCanReact bool `json:"viewerCanReact"`
	// Check if the viewer is able to change their subscription status for the repository.
	ViewerCanSubscribe bool `json:"viewerCanSubscribe"`
	// Check if the current viewer can update this object.
	ViewerCanUpdate bool `json:"viewerCanUpdate"`
	// Reasons why the current viewer can not update this comment.
	ViewerCannotUpdateReasons []CommentCannotUpdateReason `json:"viewerCannotUpdateReasons"`
	// Did the viewer author this comment.
	ViewerDidAuthor bool `json:"viewerDidAuthor"`
	// Identifies if the viewer is watching, not watching, or ignoring the subscribable entity.
	ViewerSubscription *SubscriptionState `json:"viewerSubscription"`
}

func (Issue) IsIssueOrPullRequest()       {}
func (Issue) IsSearchResultItem()         {}
func (Issue) IsReferencedSubject()        {}
func (Issue) IsMilestoneItem()            {}
func (Issue) IsRenamedTitleSubject()      {}
func (Issue) IsProjectCardItem()          {}
func (Issue) IsNode()                     {}
func (Issue) IsAssignable()               {}
func (Issue) IsClosable()                 {}
func (Issue) IsComment()                  {}
func (Issue) IsUpdatable()                {}
func (Issue) IsUpdatableComment()         {}
func (Issue) IsLabelable()                {}
func (Issue) IsLockable()                 {}
func (Issue) IsReactable()                {}
func (Issue) IsRepositoryNode()           {}
func (Issue) IsSubscribable()             {}
func (Issue) IsUniformResourceLocatable() {}

// Represents a comment on an Issue.
type IssueComment struct {
	// The actor who authored the comment.
	Author Actor `json:"author"`
	// Author's association with the subject of the comment.
	AuthorAssociation CommentAuthorAssociation `json:"authorAssociation"`
	// The body as Markdown.
	Body string `json:"body"`
	// The body rendered to HTML.
	BodyHTML string `json:"bodyHTML"`
	// The body rendered to text.
	BodyText string `json:"bodyText"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	// Check if this comment was created via an email reply.
	CreatedViaEmail bool `json:"createdViaEmail"`
	// Identifies the primary key from the database.
	DatabaseID *int `json:"databaseId"`
	// The actor who edited the comment.
	Editor Actor  `json:"editor"`
	ID     string `json:"id"`
	// Check if this comment was edited and includes an edit with the creation data
	IncludesCreatedEdit bool `json:"includesCreatedEdit"`
	// Returns whether or not a comment has been minimized.
	IsMinimized bool `json:"isMinimized"`
	// Identifies the issue associated with the comment.
	Issue *Issue `json:"issue"`
	// The moment the editor made the last edit
	LastEditedAt *string `json:"lastEditedAt"`
	// Returns why the comment was minimized.
	MinimizedReason *string `json:"minimizedReason"`
	// Identifies when the comment was published at.
	PublishedAt *string `json:"publishedAt"`
	// Returns the pull request associated with the comment, if this comment was made on a
	// pull request.
	//
	PullRequest *PullRequest `json:"pullRequest"`
	// A list of reactions grouped by content left on the subject.
	ReactionGroups []*ReactionGroup `json:"reactionGroups"`
	// A list of Reactions left on the Issue.
	Reactions *ReactionConnection `json:"reactions"`
	// The repository associated with this node.
	Repository *Repository `json:"repository"`
	// The HTTP path for this issue comment
	ResourcePath string `json:"resourcePath"`
	// Identifies the date and time when the object was last updated.
	UpdatedAt string `json:"updatedAt"`
	// The HTTP URL for this issue comment
	URL string `json:"url"`
	// A list of edits to this content.
	UserContentEdits *UserContentEditConnection `json:"userContentEdits"`
	// Check if the current viewer can delete this object.
	ViewerCanDelete bool `json:"viewerCanDelete"`
	// Check if the current viewer can minimize this object.
	ViewerCanMinimize bool `json:"viewerCanMinimize"`
	// Can user react to this subject
	ViewerCanReact bool `json:"viewerCanReact"`
	// Check if the current viewer can update this object.
	ViewerCanUpdate bool `json:"viewerCanUpdate"`
	// Reasons why the current viewer can not update this comment.
	ViewerCannotUpdateReasons []CommentCannotUpdateReason `json:"viewerCannotUpdateReasons"`
	// Did the viewer author this comment.
	ViewerDidAuthor bool `json:"viewerDidAuthor"`
}

func (IssueComment) IsPullRequestTimelineItems() {}
func (IssueComment) IsIssueTimelineItems()       {}
func (IssueComment) IsPullRequestTimelineItem()  {}
func (IssueComment) IsIssueTimelineItem()        {}
func (IssueComment) IsNode()                     {}
func (IssueComment) IsComment()                  {}
func (IssueComment) IsDeletable()                {}
func (IssueComment) IsMinimizable()              {}
func (IssueComment) IsUpdatable()                {}
func (IssueComment) IsUpdatableComment()         {}
func (IssueComment) IsReactable()                {}
func (IssueComment) IsRepositoryNode()           {}

// The connection type for IssueComment.
type IssueCommentConnection struct {
	// A list of edges.
	Edges []*IssueCommentEdge `json:"edges"`
	// A list of nodes.
	Nodes []*IssueComment `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type IssueCommentEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *IssueComment `json:"node"`
}

// The connection type for Issue.
type IssueConnection struct {
	// A list of edges.
	Edges []*IssueEdge `json:"edges"`
	// A list of nodes.
	Nodes []*Issue `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// This aggregates issues opened by a user within one repository.
type IssueContributionsByRepository struct {
	// The issue contributions.
	Contributions *CreatedIssueContributionConnection `json:"contributions"`
	// The repository in which the issues were opened.
	Repository *Repository `json:"repository"`
}

// An edge in a connection.
type IssueEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *Issue `json:"node"`
}

// Ways in which to filter lists of issues.
type IssueFilters struct {
	// List issues assigned to given name. Pass in `null` for issues with no assigned user, and `*` for issues assigned to any user.
	Assignee *string `json:"assignee"`
	// List issues created by given name.
	CreatedBy *string `json:"createdBy"`
	// List issues where the list of label names exist on the issue.
	Labels []string `json:"labels"`
	// List issues where the given name is mentioned in the issue.
	Mentioned *string `json:"mentioned"`
	// List issues by given milestone argument. If an string representation of an integer is passed, it should refer to a milestone by its number field. Pass in `null` for issues with no milestone, and `*` for issues that are assigned to any milestone.
	Milestone *string `json:"milestone"`
	// List issues that have been updated at or after the given date.
	Since *string `json:"since"`
	// List issues filtered by the list of states given.
	States []IssueState `json:"states"`
	// List issues subscribed to by viewer.
	ViewerSubscribed *bool `json:"viewerSubscribed"`
}

// Ways in which lists of issues can be ordered upon return.
type IssueOrder struct {
	// The field in which to order issues by.
	Field IssueOrderField `json:"field"`
	// The direction in which to order issues by the specified field.
	Direction OrderDirection `json:"direction"`
}

// A repository issue template.
type IssueTemplate struct {
	// The template purpose.
	About *string `json:"about"`
	// The suggested issue body.
	Body *string `json:"body"`
	// The template name.
	Name string `json:"name"`
	// The suggested issue title.
	Title *string `json:"title"`
}

// The connection type for IssueTimelineItem.
type IssueTimelineConnection struct {
	// A list of edges.
	Edges []*IssueTimelineItemEdge `json:"edges"`
	// A list of nodes.
	Nodes []IssueTimelineItem `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type IssueTimelineItemEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node IssueTimelineItem `json:"node"`
}

// The connection type for IssueTimelineItems.
type IssueTimelineItemsConnection struct {
	// A list of edges.
	Edges []*IssueTimelineItemsEdge `json:"edges"`
	// Identifies the count of items after applying `before` and `after` filters.
	FilteredCount int `json:"filteredCount"`
	// A list of nodes.
	Nodes []IssueTimelineItems `json:"nodes"`
	// Identifies the count of items after applying `before`/`after` filters and `first`/`last`/`skip` slicing.
	PageCount int `json:"pageCount"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
	// Identifies the date and time when the timeline was last updated.
	UpdatedAt string `json:"updatedAt"`
}

// An edge in a connection.
type IssueTimelineItemsEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node IssueTimelineItems `json:"node"`
}

// Represents a user signing up for a GitHub account.
type JoinedGitHubContribution struct {
	// Whether this contribution is associated with a record you do not have access to. For
	// example, your own 'first issue' contribution may have been made on a repository you can no
	// longer access.
	//
	IsRestricted bool `json:"isRestricted"`
	// When this contribution was made.
	OccurredAt string `json:"occurredAt"`
	// The HTTP path for this contribution.
	ResourcePath string `json:"resourcePath"`
	// The HTTP URL for this contribution.
	URL string `json:"url"`
	// The user who made this contribution.
	//
	User *User `json:"user"`
}

func (JoinedGitHubContribution) IsContribution() {}

// A label for categorizing Issues or Milestones with a given Repository.
type Label struct {
	// Identifies the label color.
	Color string `json:"color"`
	// Identifies the date and time when the label was created.
	CreatedAt *string `json:"createdAt"`
	// A brief description of this label.
	Description *string `json:"description"`
	ID          string  `json:"id"`
	// Indicates whether or not this is a default label.
	IsDefault bool `json:"isDefault"`
	// A list of issues associated with this label.
	Issues *IssueConnection `json:"issues"`
	// Identifies the label name.
	Name string `json:"name"`
	// A list of pull requests associated with this label.
	PullRequests *PullRequestConnection `json:"pullRequests"`
	// The repository associated with this label.
	Repository *Repository `json:"repository"`
	// The HTTP path for this label.
	ResourcePath string `json:"resourcePath"`
	// Identifies the date and time when the label was last updated.
	UpdatedAt *string `json:"updatedAt"`
	// The HTTP URL for this label.
	URL string `json:"url"`
}

func (Label) IsNode() {}

// The connection type for Label.
type LabelConnection struct {
	// A list of edges.
	Edges []*LabelEdge `json:"edges"`
	// A list of nodes.
	Nodes []*Label `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type LabelEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *Label `json:"node"`
}

// Ways in which lists of labels can be ordered upon return.
type LabelOrder struct {
	// The field in which to order labels by.
	Field LabelOrderField `json:"field"`
	// The direction in which to order labels by the specified field.
	Direction OrderDirection `json:"direction"`
}

// Represents a 'labeled' event on a given issue or pull request.
type LabeledEvent struct {
	// Identifies the actor who performed the event.
	Actor Actor `json:"actor"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// Identifies the label associated with the 'labeled' event.
	Label *Label `json:"label"`
	// Identifies the `Labelable` associated with the event.
	Labelable Labelable `json:"labelable"`
}

func (LabeledEvent) IsPullRequestTimelineItems() {}
func (LabeledEvent) IsIssueTimelineItems()       {}
func (LabeledEvent) IsPullRequestTimelineItem()  {}
func (LabeledEvent) IsIssueTimelineItem()        {}
func (LabeledEvent) IsNode()                     {}

// Represents a given language found in repositories.
type Language struct {
	// The color defined for the current language.
	Color *string `json:"color"`
	ID    string  `json:"id"`
	// The name of the current language.
	Name string `json:"name"`
}

func (Language) IsNode() {}

// A list of languages associated with the parent.
type LanguageConnection struct {
	// A list of edges.
	Edges []*LanguageEdge `json:"edges"`
	// A list of nodes.
	Nodes []*Language `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
	// The total size in bytes of files written in that language.
	TotalSize int `json:"totalSize"`
}

// Represents the language of a repository.
type LanguageEdge struct {
	Cursor string    `json:"cursor"`
	Node   *Language `json:"node"`
	// The number of bytes of code written in the language.
	Size int `json:"size"`
}

// Ordering options for language connections.
type LanguageOrder struct {
	// The field to order languages by.
	Field LanguageOrderField `json:"field"`
	// The ordering direction.
	Direction OrderDirection `json:"direction"`
}

// A repository's open source license
type License struct {
	// The full text of the license
	Body string `json:"body"`
	// The conditions set by the license
	Conditions []*LicenseRule `json:"conditions"`
	// A human-readable description of the license
	Description *string `json:"description"`
	// Whether the license should be featured
	Featured bool `json:"featured"`
	// Whether the license should be displayed in license pickers
	Hidden bool   `json:"hidden"`
	ID     string `json:"id"`
	// Instructions on how to implement the license
	Implementation *string `json:"implementation"`
	// The lowercased SPDX ID of the license
	Key string `json:"key"`
	// The limitations set by the license
	Limitations []*LicenseRule `json:"limitations"`
	// The license full name specified by <https://spdx.org/licenses>
	Name string `json:"name"`
	// Customary short name if applicable (e.g, GPLv3)
	Nickname *string `json:"nickname"`
	// The permissions set by the license
	Permissions []*LicenseRule `json:"permissions"`
	// Whether the license is a pseudo-license placeholder (e.g., other, no-license)
	PseudoLicense bool `json:"pseudoLicense"`
	// Short identifier specified by <https://spdx.org/licenses>
	SpdxID *string `json:"spdxId"`
	// URL to the license on <https://choosealicense.com>
	URL *string `json:"url"`
}

func (License) IsNode() {}

// Describes a License's conditions, permissions, and limitations
type LicenseRule struct {
	// A description of the rule
	Description string `json:"description"`
	// The machine-readable rule key
	Key string `json:"key"`
	// The human-readable rule label
	Label string `json:"label"`
}

// Autogenerated input type of LinkRepositoryToProject
type LinkRepositoryToProjectInput struct {
	// The ID of the Project to link to a Repository
	ProjectID string `json:"projectId"`
	// The ID of the Repository to link to a Project.
	RepositoryID string `json:"repositoryId"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of LinkRepositoryToProject
type LinkRepositoryToProjectPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The linked Project.
	Project *Project `json:"project"`
	// The linked Repository.
	Repository *Repository `json:"repository"`
}

// Autogenerated input type of LockLockable
type LockLockableInput struct {
	// ID of the issue or pull request to be locked.
	LockableID string `json:"lockableId"`
	// A reason for why the issue or pull request will be locked.
	LockReason *LockReason `json:"lockReason"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of LockLockable
type LockLockablePayload struct {
	// Identifies the actor who performed the event.
	Actor Actor `json:"actor"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The item that was locked.
	LockedRecord Lockable `json:"lockedRecord"`
}

// Represents a 'locked' event on a given issue or pull request.
type LockedEvent struct {
	// Identifies the actor who performed the event.
	Actor Actor `json:"actor"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// Reason that the conversation was locked (optional).
	LockReason *LockReason `json:"lockReason"`
	// Object that was locked.
	Lockable Lockable `json:"lockable"`
}

func (LockedEvent) IsPullRequestTimelineItems() {}
func (LockedEvent) IsIssueTimelineItems()       {}
func (LockedEvent) IsPullRequestTimelineItem()  {}
func (LockedEvent) IsIssueTimelineItem()        {}
func (LockedEvent) IsNode()                     {}

// A placeholder user for attribution of imported data on GitHub.
type Mannequin struct {
	// A URL pointing to the GitHub App's public avatar.
	AvatarURL string `json:"avatarUrl"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	// Identifies the primary key from the database.
	DatabaseID *int `json:"databaseId"`
	// The mannequin's email on the source instance.
	Email *string `json:"email"`
	ID    string  `json:"id"`
	// The username of the actor.
	Login string `json:"login"`
	// The HTML path to this resource.
	ResourcePath string `json:"resourcePath"`
	// Identifies the date and time when the object was last updated.
	UpdatedAt string `json:"updatedAt"`
	// The URL to this resource.
	URL string `json:"url"`
}

func (Mannequin) IsRequestedReviewer()        {}
func (Mannequin) IsNode()                     {}
func (Mannequin) IsActor()                    {}
func (Mannequin) IsUniformResourceLocatable() {}
func (Mannequin) IsAssignee()                 {}

// Autogenerated input type of MarkFileAsViewed
type MarkFileAsViewedInput struct {
	// The Node ID of the pull request.
	PullRequestID string `json:"pullRequestId"`
	// The path of the file to mark as viewed
	Path string `json:"path"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of MarkFileAsViewed
type MarkFileAsViewedPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The updated pull request.
	PullRequest *PullRequest `json:"pullRequest"`
}

// Autogenerated input type of MarkPullRequestReadyForReview
type MarkPullRequestReadyForReviewInput struct {
	// ID of the pull request to be marked as ready for review.
	PullRequestID string `json:"pullRequestId"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of MarkPullRequestReadyForReview
type MarkPullRequestReadyForReviewPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The pull request that is ready for review.
	PullRequest *PullRequest `json:"pullRequest"`
}

// Represents a 'marked_as_duplicate' event on a given issue or pull request.
type MarkedAsDuplicateEvent struct {
	// Identifies the actor who performed the event.
	Actor Actor `json:"actor"`
	// The authoritative issue or pull request which has been duplicated by another.
	Canonical IssueOrPullRequest `json:"canonical"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	// The issue or pull request which has been marked as a duplicate of another.
	Duplicate IssueOrPullRequest `json:"duplicate"`
	ID        string             `json:"id"`
	// Canonical and duplicate belong to different repositories.
	IsCrossRepository bool `json:"isCrossRepository"`
}

func (MarkedAsDuplicateEvent) IsNode()                     {}
func (MarkedAsDuplicateEvent) IsPullRequestTimelineItems() {}
func (MarkedAsDuplicateEvent) IsIssueTimelineItems()       {}

// A public description of a Marketplace category.
type MarketplaceCategory struct {
	// The category's description.
	Description *string `json:"description"`
	// The technical description of how apps listed in this category work with GitHub.
	HowItWorks *string `json:"howItWorks"`
	ID         string  `json:"id"`
	// The category's name.
	Name string `json:"name"`
	// How many Marketplace listings have this as their primary category.
	PrimaryListingCount int `json:"primaryListingCount"`
	// The HTTP path for this Marketplace category.
	ResourcePath string `json:"resourcePath"`
	// How many Marketplace listings have this as their secondary category.
	SecondaryListingCount int `json:"secondaryListingCount"`
	// The short name of the category used in its URL.
	Slug string `json:"slug"`
	// The HTTP URL for this Marketplace category.
	URL string `json:"url"`
}

func (MarketplaceCategory) IsNode() {}

// A listing in the GitHub integration marketplace.
type MarketplaceListing struct {
	// The GitHub App this listing represents.
	App *App `json:"app"`
	// URL to the listing owner's company site.
	CompanyURL *string `json:"companyUrl"`
	// The HTTP path for configuring access to the listing's integration or OAuth app
	ConfigurationResourcePath string `json:"configurationResourcePath"`
	// The HTTP URL for configuring access to the listing's integration or OAuth app
	ConfigurationURL string `json:"configurationUrl"`
	// URL to the listing's documentation.
	DocumentationURL *string `json:"documentationUrl"`
	// The listing's detailed description.
	ExtendedDescription *string `json:"extendedDescription"`
	// The listing's detailed description rendered to HTML.
	ExtendedDescriptionHTML string `json:"extendedDescriptionHTML"`
	// The listing's introductory description.
	FullDescription string `json:"fullDescription"`
	// The listing's introductory description rendered to HTML.
	FullDescriptionHTML string `json:"fullDescriptionHTML"`
	// Does this listing have any plans with a free trial?
	HasPublishedFreeTrialPlans bool `json:"hasPublishedFreeTrialPlans"`
	// Does this listing have a terms of service link?
	HasTermsOfService bool `json:"hasTermsOfService"`
	// A technical description of how this app works with GitHub.
	HowItWorks *string `json:"howItWorks"`
	// The listing's technical description rendered to HTML.
	HowItWorksHTML string `json:"howItWorksHTML"`
	ID             string `json:"id"`
	// URL to install the product to the viewer's account or organization.
	InstallationURL *string `json:"installationUrl"`
	// Whether this listing's app has been installed for the current viewer
	InstalledForViewer bool `json:"installedForViewer"`
	// Whether this listing has been removed from the Marketplace.
	IsArchived bool `json:"isArchived"`
	// Whether this listing is still an editable draft that has not been submitted for review and is not publicly visible in the Marketplace.
	IsDraft bool `json:"isDraft"`
	// Whether the product this listing represents is available as part of a paid plan.
	IsPaid bool `json:"isPaid"`
	// Whether this listing has been approved for display in the Marketplace.
	IsPublic bool `json:"isPublic"`
	// Whether this listing has been rejected by GitHub for display in the Marketplace.
	IsRejected bool `json:"isRejected"`
	// Whether this listing has been approved for unverified display in the Marketplace.
	IsUnverified bool `json:"isUnverified"`
	// Whether this draft listing has been submitted for review for approval to be unverified in the Marketplace.
	IsUnverifiedPending bool `json:"isUnverifiedPending"`
	// Whether this draft listing has been submitted for review from GitHub for approval to be verified in the Marketplace.
	IsVerificationPendingFromDraft bool `json:"isVerificationPendingFromDraft"`
	// Whether this unverified listing has been submitted for review from GitHub for approval to be verified in the Marketplace.
	IsVerificationPendingFromUnverified bool `json:"isVerificationPendingFromUnverified"`
	// Whether this listing has been approved for verified display in the Marketplace.
	IsVerified bool `json:"isVerified"`
	// The hex color code, without the leading '#', for the logo background.
	LogoBackgroundColor string `json:"logoBackgroundColor"`
	// URL for the listing's logo image.
	LogoURL *string `json:"logoUrl"`
	// The listing's full name.
	Name string `json:"name"`
	// The listing's very short description without a trailing period or ampersands.
	NormalizedShortDescription string `json:"normalizedShortDescription"`
	// URL to the listing's detailed pricing.
	PricingURL *string `json:"pricingUrl"`
	// The category that best describes the listing.
	PrimaryCategory *MarketplaceCategory `json:"primaryCategory"`
	// URL to the listing's privacy policy, may return an empty string for listings that do not require a privacy policy URL.
	PrivacyPolicyURL string `json:"privacyPolicyUrl"`
	// The HTTP path for the Marketplace listing.
	ResourcePath string `json:"resourcePath"`
	// The URLs for the listing's screenshots.
	ScreenshotUrls []*string `json:"screenshotUrls"`
	// An alternate category that describes the listing.
	SecondaryCategory *MarketplaceCategory `json:"secondaryCategory"`
	// The listing's very short description.
	ShortDescription string `json:"shortDescription"`
	// The short name of the listing used in its URL.
	Slug string `json:"slug"`
	// URL to the listing's status page.
	StatusURL *string `json:"statusUrl"`
	// An email address for support for this listing's app.
	SupportEmail *string `json:"supportEmail"`
	// Either a URL or an email address for support for this listing's app, may return an empty string for listings that do not require a support URL.
	SupportURL string `json:"supportUrl"`
	// URL to the listing's terms of service.
	TermsOfServiceURL *string `json:"termsOfServiceUrl"`
	// The HTTP URL for the Marketplace listing.
	URL string `json:"url"`
	// Can the current viewer add plans for this Marketplace listing.
	ViewerCanAddPlans bool `json:"viewerCanAddPlans"`
	// Can the current viewer approve this Marketplace listing.
	ViewerCanApprove bool `json:"viewerCanApprove"`
	// Can the current viewer delist this Marketplace listing.
	ViewerCanDelist bool `json:"viewerCanDelist"`
	// Can the current viewer edit this Marketplace listing.
	ViewerCanEdit bool `json:"viewerCanEdit"`
	// Can the current viewer edit the primary and secondary category of this
	// Marketplace listing.
	//
	ViewerCanEditCategories bool `json:"viewerCanEditCategories"`
	// Can the current viewer edit the plans for this Marketplace listing.
	ViewerCanEditPlans bool `json:"viewerCanEditPlans"`
	// Can the current viewer return this Marketplace listing to draft state
	// so it becomes editable again.
	//
	ViewerCanRedraft bool `json:"viewerCanRedraft"`
	// Can the current viewer reject this Marketplace listing by returning it to
	// an editable draft state or rejecting it entirely.
	//
	ViewerCanReject bool `json:"viewerCanReject"`
	// Can the current viewer request this listing be reviewed for display in
	// the Marketplace as verified.
	//
	ViewerCanRequestApproval bool `json:"viewerCanRequestApproval"`
	// Indicates whether the current user has an active subscription to this Marketplace listing.
	//
	ViewerHasPurchased bool `json:"viewerHasPurchased"`
	// Indicates if the current user has purchased a subscription to this Marketplace listing
	// for all of the organizations the user owns.
	//
	ViewerHasPurchasedForAllOrganizations bool `json:"viewerHasPurchasedForAllOrganizations"`
	// Does the current viewer role allow them to administer this Marketplace listing.
	//
	ViewerIsListingAdmin bool `json:"viewerIsListingAdmin"`
}

func (MarketplaceListing) IsSearchResultItem() {}
func (MarketplaceListing) IsNode()             {}

// Look up Marketplace Listings
type MarketplaceListingConnection struct {
	// A list of edges.
	Edges []*MarketplaceListingEdge `json:"edges"`
	// A list of nodes.
	Nodes []*MarketplaceListing `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type MarketplaceListingEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *MarketplaceListing `json:"node"`
}

// Audit log entry for a members_can_delete_repos.clear event.
type MembersCanDeleteReposClearAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	// The HTTP path for this enterprise.
	EnterpriseResourcePath *string `json:"enterpriseResourcePath"`
	// The slug of the enterprise.
	EnterpriseSlug *string `json:"enterpriseSlug"`
	// The HTTP URL for this enterprise.
	EnterpriseURL *string `json:"enterpriseUrl"`
	ID            string  `json:"id"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
}

func (MembersCanDeleteReposClearAuditEntry) IsNode()                       {}
func (MembersCanDeleteReposClearAuditEntry) IsAuditEntry()                 {}
func (MembersCanDeleteReposClearAuditEntry) IsEnterpriseAuditEntryData()   {}
func (MembersCanDeleteReposClearAuditEntry) IsOrganizationAuditEntryData() {}
func (MembersCanDeleteReposClearAuditEntry) IsOrganizationAuditEntry()     {}

// Audit log entry for a members_can_delete_repos.disable event.
type MembersCanDeleteReposDisableAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	// The HTTP path for this enterprise.
	EnterpriseResourcePath *string `json:"enterpriseResourcePath"`
	// The slug of the enterprise.
	EnterpriseSlug *string `json:"enterpriseSlug"`
	// The HTTP URL for this enterprise.
	EnterpriseURL *string `json:"enterpriseUrl"`
	ID            string  `json:"id"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
}

func (MembersCanDeleteReposDisableAuditEntry) IsNode()                       {}
func (MembersCanDeleteReposDisableAuditEntry) IsAuditEntry()                 {}
func (MembersCanDeleteReposDisableAuditEntry) IsEnterpriseAuditEntryData()   {}
func (MembersCanDeleteReposDisableAuditEntry) IsOrganizationAuditEntryData() {}
func (MembersCanDeleteReposDisableAuditEntry) IsOrganizationAuditEntry()     {}

// Audit log entry for a members_can_delete_repos.enable event.
type MembersCanDeleteReposEnableAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	// The HTTP path for this enterprise.
	EnterpriseResourcePath *string `json:"enterpriseResourcePath"`
	// The slug of the enterprise.
	EnterpriseSlug *string `json:"enterpriseSlug"`
	// The HTTP URL for this enterprise.
	EnterpriseURL *string `json:"enterpriseUrl"`
	ID            string  `json:"id"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
}

func (MembersCanDeleteReposEnableAuditEntry) IsOrganizationAuditEntry()     {}
func (MembersCanDeleteReposEnableAuditEntry) IsNode()                       {}
func (MembersCanDeleteReposEnableAuditEntry) IsAuditEntry()                 {}
func (MembersCanDeleteReposEnableAuditEntry) IsEnterpriseAuditEntryData()   {}
func (MembersCanDeleteReposEnableAuditEntry) IsOrganizationAuditEntryData() {}

// Represents a 'mentioned' event on a given issue or pull request.
type MentionedEvent struct {
	// Identifies the actor who performed the event.
	Actor Actor `json:"actor"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	// Identifies the primary key from the database.
	DatabaseID *int   `json:"databaseId"`
	ID         string `json:"id"`
}

func (MentionedEvent) IsPullRequestTimelineItems() {}
func (MentionedEvent) IsIssueTimelineItems()       {}
func (MentionedEvent) IsNode()                     {}

// Autogenerated input type of MergeBranch
type MergeBranchInput struct {
	// The Node ID of the Repository containing the base branch that will be modified.
	RepositoryID string `json:"repositoryId"`
	// The name of the base branch that the provided head will be merged into.
	Base string `json:"base"`
	// The head to merge into the base branch. This can be a branch name or a commit GitObjectID.
	Head string `json:"head"`
	// Message to use for the merge commit. If omitted, a default will be used.
	CommitMessage *string `json:"commitMessage"`
	// The email address to associate with this commit.
	AuthorEmail *string `json:"authorEmail"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of MergeBranch
type MergeBranchPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The resulting merge Commit.
	MergeCommit *Commit `json:"mergeCommit"`
}

// Autogenerated input type of MergePullRequest
type MergePullRequestInput struct {
	// ID of the pull request to be merged.
	PullRequestID string `json:"pullRequestId"`
	// Commit headline to use for the merge commit; if omitted, a default message will be used.
	CommitHeadline *string `json:"commitHeadline"`
	// Commit body to use for the merge commit; if omitted, a default message will be used
	CommitBody *string `json:"commitBody"`
	// OID that the pull request head ref must match to allow merge; if omitted, no check is performed.
	ExpectedHeadOid *string `json:"expectedHeadOid"`
	// The merge method to use. If omitted, defaults to 'MERGE'
	MergeMethod *PullRequestMergeMethod `json:"mergeMethod"`
	// The email address to associate with this merge.
	AuthorEmail *string `json:"authorEmail"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of MergePullRequest
type MergePullRequestPayload struct {
	// Identifies the actor who performed the event.
	Actor Actor `json:"actor"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The pull request that was merged.
	PullRequest *PullRequest `json:"pullRequest"`
}

// Represents a 'merged' event on a given pull request.
type MergedEvent struct {
	// Identifies the actor who performed the event.
	Actor Actor `json:"actor"`
	// Identifies the commit associated with the `merge` event.
	Commit *Commit `json:"commit"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// Identifies the Ref associated with the `merge` event.
	MergeRef *Ref `json:"mergeRef"`
	// Identifies the name of the Ref associated with the `merge` event.
	MergeRefName string `json:"mergeRefName"`
	// PullRequest referenced by event.
	PullRequest *PullRequest `json:"pullRequest"`
	// The HTTP path for this merged event.
	ResourcePath string `json:"resourcePath"`
	// The HTTP URL for this merged event.
	URL string `json:"url"`
}

func (MergedEvent) IsPullRequestTimelineItems() {}
func (MergedEvent) IsPullRequestTimelineItem()  {}
func (MergedEvent) IsNode()                     {}
func (MergedEvent) IsUniformResourceLocatable() {}

// Represents a Milestone object on a given repository.
type Milestone struct {
	// `true` if the object is closed (definition of closed may depend on type)
	Closed bool `json:"closed"`
	// Identifies the date and time when the object was closed.
	ClosedAt *string `json:"closedAt"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	// Identifies the actor who created the milestone.
	Creator Actor `json:"creator"`
	// Identifies the description of the milestone.
	Description *string `json:"description"`
	// Identifies the due date of the milestone.
	DueOn *string `json:"dueOn"`
	ID    string  `json:"id"`
	// Just for debugging on review-lab
	IssuePrioritiesDebug string `json:"issuePrioritiesDebug"`
	// A list of issues associated with the milestone.
	Issues *IssueConnection `json:"issues"`
	// Identifies the number of the milestone.
	Number int `json:"number"`
	// Indentifies the percentage complete for the milestone
	ProgressPercentage float64 `json:"progressPercentage"`
	// A list of pull requests associated with the milestone.
	PullRequests *PullRequestConnection `json:"pullRequests"`
	// The repository associated with this milestone.
	Repository *Repository `json:"repository"`
	// The HTTP path for this milestone
	ResourcePath string `json:"resourcePath"`
	// Identifies the state of the milestone.
	State MilestoneState `json:"state"`
	// Identifies the title of the milestone.
	Title string `json:"title"`
	// Identifies the date and time when the object was last updated.
	UpdatedAt string `json:"updatedAt"`
	// The HTTP URL for this milestone
	URL string `json:"url"`
}

func (Milestone) IsNode()                     {}
func (Milestone) IsClosable()                 {}
func (Milestone) IsUniformResourceLocatable() {}

// The connection type for Milestone.
type MilestoneConnection struct {
	// A list of edges.
	Edges []*MilestoneEdge `json:"edges"`
	// A list of nodes.
	Nodes []*Milestone `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type MilestoneEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *Milestone `json:"node"`
}

// Ordering options for milestone connections.
type MilestoneOrder struct {
	// The field to order milestones by.
	Field MilestoneOrderField `json:"field"`
	// The ordering direction.
	Direction OrderDirection `json:"direction"`
}

// Represents a 'milestoned' event on a given issue or pull request.
type MilestonedEvent struct {
	// Identifies the actor who performed the event.
	Actor Actor `json:"actor"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// Identifies the milestone title associated with the 'milestoned' event.
	MilestoneTitle string `json:"milestoneTitle"`
	// Object referenced by event.
	Subject MilestoneItem `json:"subject"`
}

func (MilestonedEvent) IsPullRequestTimelineItems() {}
func (MilestonedEvent) IsIssueTimelineItems()       {}
func (MilestonedEvent) IsPullRequestTimelineItem()  {}
func (MilestonedEvent) IsIssueTimelineItem()        {}
func (MilestonedEvent) IsNode()                     {}

// Autogenerated input type of MinimizeComment
type MinimizeCommentInput struct {
	// The Node ID of the subject to modify.
	SubjectID string `json:"subjectId"`
	// The classification of comment
	Classifier ReportedContentClassifiers `json:"classifier"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of MinimizeComment
type MinimizeCommentPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The comment that was minimized.
	MinimizedComment Minimizable `json:"minimizedComment"`
}

// Autogenerated input type of MoveProjectCard
type MoveProjectCardInput struct {
	// The id of the card to move.
	CardID string `json:"cardId"`
	// The id of the column to move it into.
	ColumnID string `json:"columnId"`
	// Place the new card after the card with this id. Pass null to place it at the top.
	AfterCardID *string `json:"afterCardId"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of MoveProjectCard
type MoveProjectCardPayload struct {
	// The new edge of the moved card.
	CardEdge *ProjectCardEdge `json:"cardEdge"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated input type of MoveProjectColumn
type MoveProjectColumnInput struct {
	// The id of the column to move.
	ColumnID string `json:"columnId"`
	// Place the new column after the column with this id. Pass null to place it at the front.
	AfterColumnID *string `json:"afterColumnId"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of MoveProjectColumn
type MoveProjectColumnPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The new edge of the moved column.
	ColumnEdge *ProjectColumnEdge `json:"columnEdge"`
}

// Represents a 'moved_columns_in_project' event on a given issue or pull request.
type MovedColumnsInProjectEvent struct {
	// Identifies the actor who performed the event.
	Actor Actor `json:"actor"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	// Identifies the primary key from the database.
	DatabaseID *int   `json:"databaseId"`
	ID         string `json:"id"`
}

func (MovedColumnsInProjectEvent) IsPullRequestTimelineItems() {}
func (MovedColumnsInProjectEvent) IsIssueTimelineItems()       {}
func (MovedColumnsInProjectEvent) IsNode()                     {}

// Audit log entry for a oauth_application.create event.
type OauthApplicationCreateAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The application URL of the OAuth Application.
	ApplicationURL *string `json:"applicationUrl"`
	// The callback URL of the OAuth Application.
	CallbackURL *string `json:"callbackUrl"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// The name of the OAuth Application.
	OauthApplicationName *string `json:"oauthApplicationName"`
	// The HTTP path for the OAuth Application
	OauthApplicationResourcePath *string `json:"oauthApplicationResourcePath"`
	// The HTTP URL for the OAuth Application
	OauthApplicationURL *string `json:"oauthApplicationUrl"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The rate limit of the OAuth Application.
	RateLimit *int `json:"rateLimit"`
	// The state of the OAuth Application.
	State *OauthApplicationCreateAuditEntryState `json:"state"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
}

func (OauthApplicationCreateAuditEntry) IsOrganizationAuditEntry()         {}
func (OauthApplicationCreateAuditEntry) IsNode()                           {}
func (OauthApplicationCreateAuditEntry) IsAuditEntry()                     {}
func (OauthApplicationCreateAuditEntry) IsOauthApplicationAuditEntryData() {}
func (OauthApplicationCreateAuditEntry) IsOrganizationAuditEntryData()     {}

// Audit log entry for a org.add_billing_manager
type OrgAddBillingManagerAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// The email address used to invite a billing manager for the organization.
	InvitationEmail *string `json:"invitationEmail"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
}

func (OrgAddBillingManagerAuditEntry) IsNode()                       {}
func (OrgAddBillingManagerAuditEntry) IsAuditEntry()                 {}
func (OrgAddBillingManagerAuditEntry) IsOrganizationAuditEntryData() {}
func (OrgAddBillingManagerAuditEntry) IsOrganizationAuditEntry()     {}

// Audit log entry for a org.add_member
type OrgAddMemberAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The permission level of the member added to the organization.
	Permission *OrgAddMemberAuditEntryPermission `json:"permission"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
}

func (OrgAddMemberAuditEntry) IsNode()                       {}
func (OrgAddMemberAuditEntry) IsAuditEntry()                 {}
func (OrgAddMemberAuditEntry) IsOrganizationAuditEntryData() {}
func (OrgAddMemberAuditEntry) IsOrganizationAuditEntry()     {}

// Audit log entry for a org.block_user
type OrgBlockUserAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The blocked user.
	BlockedUser *User `json:"blockedUser"`
	// The username of the blocked user.
	BlockedUserName *string `json:"blockedUserName"`
	// The HTTP path for the blocked user.
	BlockedUserResourcePath *string `json:"blockedUserResourcePath"`
	// The HTTP URL for the blocked user.
	BlockedUserURL *string `json:"blockedUserUrl"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
}

func (OrgBlockUserAuditEntry) IsNode()                       {}
func (OrgBlockUserAuditEntry) IsAuditEntry()                 {}
func (OrgBlockUserAuditEntry) IsOrganizationAuditEntryData() {}
func (OrgBlockUserAuditEntry) IsOrganizationAuditEntry()     {}

// Audit log entry for a org.config.disable_collaborators_only event.
type OrgConfigDisableCollaboratorsOnlyAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
}

func (OrgConfigDisableCollaboratorsOnlyAuditEntry) IsNode()                       {}
func (OrgConfigDisableCollaboratorsOnlyAuditEntry) IsAuditEntry()                 {}
func (OrgConfigDisableCollaboratorsOnlyAuditEntry) IsOrganizationAuditEntryData() {}
func (OrgConfigDisableCollaboratorsOnlyAuditEntry) IsOrganizationAuditEntry()     {}

// Audit log entry for a org.config.enable_collaborators_only event.
type OrgConfigEnableCollaboratorsOnlyAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
}

func (OrgConfigEnableCollaboratorsOnlyAuditEntry) IsOrganizationAuditEntry()     {}
func (OrgConfigEnableCollaboratorsOnlyAuditEntry) IsNode()                       {}
func (OrgConfigEnableCollaboratorsOnlyAuditEntry) IsAuditEntry()                 {}
func (OrgConfigEnableCollaboratorsOnlyAuditEntry) IsOrganizationAuditEntryData() {}

// Audit log entry for a org.create event.
type OrgCreateAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The billing plan for the Organization.
	BillingPlan *OrgCreateAuditEntryBillingPlan `json:"billingPlan"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
}

func (OrgCreateAuditEntry) IsNode()                       {}
func (OrgCreateAuditEntry) IsAuditEntry()                 {}
func (OrgCreateAuditEntry) IsOrganizationAuditEntryData() {}
func (OrgCreateAuditEntry) IsOrganizationAuditEntry()     {}

// Audit log entry for a org.disable_oauth_app_restrictions event.
type OrgDisableOauthAppRestrictionsAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
}

func (OrgDisableOauthAppRestrictionsAuditEntry) IsOrganizationAuditEntry()     {}
func (OrgDisableOauthAppRestrictionsAuditEntry) IsNode()                       {}
func (OrgDisableOauthAppRestrictionsAuditEntry) IsAuditEntry()                 {}
func (OrgDisableOauthAppRestrictionsAuditEntry) IsOrganizationAuditEntryData() {}

// Audit log entry for a org.disable_saml event.
type OrgDisableSamlAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	// The SAML provider's digest algorithm URL.
	DigestMethodURL *string `json:"digestMethodUrl"`
	ID              string  `json:"id"`
	// The SAML provider's issuer URL.
	IssuerURL *string `json:"issuerUrl"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The SAML provider's signature algorithm URL.
	SignatureMethodURL *string `json:"signatureMethodUrl"`
	// The SAML provider's single sign-on URL.
	SingleSignOnURL *string `json:"singleSignOnUrl"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
}

func (OrgDisableSamlAuditEntry) IsNode()                       {}
func (OrgDisableSamlAuditEntry) IsAuditEntry()                 {}
func (OrgDisableSamlAuditEntry) IsOrganizationAuditEntryData() {}
func (OrgDisableSamlAuditEntry) IsOrganizationAuditEntry()     {}

// Audit log entry for a org.disable_two_factor_requirement event.
type OrgDisableTwoFactorRequirementAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
}

func (OrgDisableTwoFactorRequirementAuditEntry) IsNode()                       {}
func (OrgDisableTwoFactorRequirementAuditEntry) IsAuditEntry()                 {}
func (OrgDisableTwoFactorRequirementAuditEntry) IsOrganizationAuditEntryData() {}
func (OrgDisableTwoFactorRequirementAuditEntry) IsOrganizationAuditEntry()     {}

// Audit log entry for a org.enable_oauth_app_restrictions event.
type OrgEnableOauthAppRestrictionsAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
}

func (OrgEnableOauthAppRestrictionsAuditEntry) IsOrganizationAuditEntry()     {}
func (OrgEnableOauthAppRestrictionsAuditEntry) IsNode()                       {}
func (OrgEnableOauthAppRestrictionsAuditEntry) IsAuditEntry()                 {}
func (OrgEnableOauthAppRestrictionsAuditEntry) IsOrganizationAuditEntryData() {}

// Audit log entry for a org.enable_saml event.
type OrgEnableSamlAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	// The SAML provider's digest algorithm URL.
	DigestMethodURL *string `json:"digestMethodUrl"`
	ID              string  `json:"id"`
	// The SAML provider's issuer URL.
	IssuerURL *string `json:"issuerUrl"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The SAML provider's signature algorithm URL.
	SignatureMethodURL *string `json:"signatureMethodUrl"`
	// The SAML provider's single sign-on URL.
	SingleSignOnURL *string `json:"singleSignOnUrl"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
}

func (OrgEnableSamlAuditEntry) IsOrganizationAuditEntry()     {}
func (OrgEnableSamlAuditEntry) IsNode()                       {}
func (OrgEnableSamlAuditEntry) IsAuditEntry()                 {}
func (OrgEnableSamlAuditEntry) IsOrganizationAuditEntryData() {}

// Audit log entry for a org.enable_two_factor_requirement event.
type OrgEnableTwoFactorRequirementAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
}

func (OrgEnableTwoFactorRequirementAuditEntry) IsNode()                       {}
func (OrgEnableTwoFactorRequirementAuditEntry) IsAuditEntry()                 {}
func (OrgEnableTwoFactorRequirementAuditEntry) IsOrganizationAuditEntryData() {}
func (OrgEnableTwoFactorRequirementAuditEntry) IsOrganizationAuditEntry()     {}

// Audit log entry for a org.invite_member event.
type OrgInviteMemberAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	// The email address of the organization invitation.
	Email *string `json:"email"`
	ID    string  `json:"id"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The organization invitation.
	OrganizationInvitation *OrganizationInvitation `json:"organizationInvitation"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
}

func (OrgInviteMemberAuditEntry) IsNode()                       {}
func (OrgInviteMemberAuditEntry) IsAuditEntry()                 {}
func (OrgInviteMemberAuditEntry) IsOrganizationAuditEntryData() {}
func (OrgInviteMemberAuditEntry) IsOrganizationAuditEntry()     {}

// Audit log entry for a org.invite_to_business event.
type OrgInviteToBusinessAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	// The HTTP path for this enterprise.
	EnterpriseResourcePath *string `json:"enterpriseResourcePath"`
	// The slug of the enterprise.
	EnterpriseSlug *string `json:"enterpriseSlug"`
	// The HTTP URL for this enterprise.
	EnterpriseURL *string `json:"enterpriseUrl"`
	ID            string  `json:"id"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
}

func (OrgInviteToBusinessAuditEntry) IsNode()                       {}
func (OrgInviteToBusinessAuditEntry) IsAuditEntry()                 {}
func (OrgInviteToBusinessAuditEntry) IsEnterpriseAuditEntryData()   {}
func (OrgInviteToBusinessAuditEntry) IsOrganizationAuditEntryData() {}
func (OrgInviteToBusinessAuditEntry) IsOrganizationAuditEntry()     {}

// Audit log entry for a org.oauth_app_access_approved event.
type OrgOauthAppAccessApprovedAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// The name of the OAuth Application.
	OauthApplicationName *string `json:"oauthApplicationName"`
	// The HTTP path for the OAuth Application
	OauthApplicationResourcePath *string `json:"oauthApplicationResourcePath"`
	// The HTTP URL for the OAuth Application
	OauthApplicationURL *string `json:"oauthApplicationUrl"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
}

func (OrgOauthAppAccessApprovedAuditEntry) IsNode()                           {}
func (OrgOauthAppAccessApprovedAuditEntry) IsAuditEntry()                     {}
func (OrgOauthAppAccessApprovedAuditEntry) IsOauthApplicationAuditEntryData() {}
func (OrgOauthAppAccessApprovedAuditEntry) IsOrganizationAuditEntryData()     {}
func (OrgOauthAppAccessApprovedAuditEntry) IsOrganizationAuditEntry()         {}

// Audit log entry for a org.oauth_app_access_denied event.
type OrgOauthAppAccessDeniedAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// The name of the OAuth Application.
	OauthApplicationName *string `json:"oauthApplicationName"`
	// The HTTP path for the OAuth Application
	OauthApplicationResourcePath *string `json:"oauthApplicationResourcePath"`
	// The HTTP URL for the OAuth Application
	OauthApplicationURL *string `json:"oauthApplicationUrl"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
}

func (OrgOauthAppAccessDeniedAuditEntry) IsNode()                           {}
func (OrgOauthAppAccessDeniedAuditEntry) IsAuditEntry()                     {}
func (OrgOauthAppAccessDeniedAuditEntry) IsOauthApplicationAuditEntryData() {}
func (OrgOauthAppAccessDeniedAuditEntry) IsOrganizationAuditEntryData()     {}
func (OrgOauthAppAccessDeniedAuditEntry) IsOrganizationAuditEntry()         {}

// Audit log entry for a org.oauth_app_access_requested event.
type OrgOauthAppAccessRequestedAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// The name of the OAuth Application.
	OauthApplicationName *string `json:"oauthApplicationName"`
	// The HTTP path for the OAuth Application
	OauthApplicationResourcePath *string `json:"oauthApplicationResourcePath"`
	// The HTTP URL for the OAuth Application
	OauthApplicationURL *string `json:"oauthApplicationUrl"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
}

func (OrgOauthAppAccessRequestedAuditEntry) IsNode()                           {}
func (OrgOauthAppAccessRequestedAuditEntry) IsAuditEntry()                     {}
func (OrgOauthAppAccessRequestedAuditEntry) IsOauthApplicationAuditEntryData() {}
func (OrgOauthAppAccessRequestedAuditEntry) IsOrganizationAuditEntryData()     {}
func (OrgOauthAppAccessRequestedAuditEntry) IsOrganizationAuditEntry()         {}

// Audit log entry for a org.remove_billing_manager event.
type OrgRemoveBillingManagerAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The reason for the billing manager being removed.
	Reason *OrgRemoveBillingManagerAuditEntryReason `json:"reason"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
}

func (OrgRemoveBillingManagerAuditEntry) IsOrganizationAuditEntry()     {}
func (OrgRemoveBillingManagerAuditEntry) IsNode()                       {}
func (OrgRemoveBillingManagerAuditEntry) IsAuditEntry()                 {}
func (OrgRemoveBillingManagerAuditEntry) IsOrganizationAuditEntryData() {}

// Audit log entry for a org.remove_member event.
type OrgRemoveMemberAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// The types of membership the member has with the organization.
	MembershipTypes []OrgRemoveMemberAuditEntryMembershipType `json:"membershipTypes"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The reason for the member being removed.
	Reason *OrgRemoveMemberAuditEntryReason `json:"reason"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
}

func (OrgRemoveMemberAuditEntry) IsOrganizationAuditEntry()     {}
func (OrgRemoveMemberAuditEntry) IsNode()                       {}
func (OrgRemoveMemberAuditEntry) IsAuditEntry()                 {}
func (OrgRemoveMemberAuditEntry) IsOrganizationAuditEntryData() {}

// Audit log entry for a org.remove_outside_collaborator event.
type OrgRemoveOutsideCollaboratorAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// The types of membership the outside collaborator has with the organization.
	MembershipTypes []OrgRemoveOutsideCollaboratorAuditEntryMembershipType `json:"membershipTypes"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The reason for the outside collaborator being removed from the Organization.
	Reason *OrgRemoveOutsideCollaboratorAuditEntryReason `json:"reason"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
}

func (OrgRemoveOutsideCollaboratorAuditEntry) IsNode()                       {}
func (OrgRemoveOutsideCollaboratorAuditEntry) IsAuditEntry()                 {}
func (OrgRemoveOutsideCollaboratorAuditEntry) IsOrganizationAuditEntryData() {}
func (OrgRemoveOutsideCollaboratorAuditEntry) IsOrganizationAuditEntry()     {}

// Audit log entry for a org.restore_member event.
type OrgRestoreMemberAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The number of custom email routings for the restored member.
	RestoredCustomEmailRoutingsCount *int `json:"restoredCustomEmailRoutingsCount"`
	// The number of issue assignemnts for the restored member.
	RestoredIssueAssignmentsCount *int `json:"restoredIssueAssignmentsCount"`
	// Restored organization membership objects.
	RestoredMemberships []OrgRestoreMemberAuditEntryMembership `json:"restoredMemberships"`
	// The number of restored memberships.
	RestoredMembershipsCount *int `json:"restoredMembershipsCount"`
	// The number of repositories of the restored member.
	RestoredRepositoriesCount *int `json:"restoredRepositoriesCount"`
	// The number of starred repositories for the restored member.
	RestoredRepositoryStarsCount *int `json:"restoredRepositoryStarsCount"`
	// The number of watched repositories for the restored member.
	RestoredRepositoryWatchesCount *int `json:"restoredRepositoryWatchesCount"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
}

func (OrgRestoreMemberAuditEntry) IsNode()                       {}
func (OrgRestoreMemberAuditEntry) IsAuditEntry()                 {}
func (OrgRestoreMemberAuditEntry) IsOrganizationAuditEntryData() {}
func (OrgRestoreMemberAuditEntry) IsOrganizationAuditEntry()     {}

// Metadata for an organization membership for org.restore_member actions
type OrgRestoreMemberMembershipOrganizationAuditEntryData struct {
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
}

func (OrgRestoreMemberMembershipOrganizationAuditEntryData) IsOrgRestoreMemberAuditEntryMembership() {
}
func (OrgRestoreMemberMembershipOrganizationAuditEntryData) IsOrganizationAuditEntryData() {}

// Metadata for a repository membership for org.restore_member actions
type OrgRestoreMemberMembershipRepositoryAuditEntryData struct {
	// The repository associated with the action
	Repository *Repository `json:"repository"`
	// The name of the repository
	RepositoryName *string `json:"repositoryName"`
	// The HTTP path for the repository
	RepositoryResourcePath *string `json:"repositoryResourcePath"`
	// The HTTP URL for the repository
	RepositoryURL *string `json:"repositoryUrl"`
}

func (OrgRestoreMemberMembershipRepositoryAuditEntryData) IsOrgRestoreMemberAuditEntryMembership() {}
func (OrgRestoreMemberMembershipRepositoryAuditEntryData) IsRepositoryAuditEntryData()             {}

// Metadata for a team membership for org.restore_member actions
type OrgRestoreMemberMembershipTeamAuditEntryData struct {
	// The team associated with the action
	Team *Team `json:"team"`
	// The name of the team
	TeamName *string `json:"teamName"`
	// The HTTP path for this team
	TeamResourcePath *string `json:"teamResourcePath"`
	// The HTTP URL for this team
	TeamURL *string `json:"teamUrl"`
}

func (OrgRestoreMemberMembershipTeamAuditEntryData) IsOrgRestoreMemberAuditEntryMembership() {}
func (OrgRestoreMemberMembershipTeamAuditEntryData) IsTeamAuditEntryData()                   {}

// Audit log entry for a org.unblock_user
type OrgUnblockUserAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The user being unblocked by the organization.
	BlockedUser *User `json:"blockedUser"`
	// The username of the blocked user.
	BlockedUserName *string `json:"blockedUserName"`
	// The HTTP path for the blocked user.
	BlockedUserResourcePath *string `json:"blockedUserResourcePath"`
	// The HTTP URL for the blocked user.
	BlockedUserURL *string `json:"blockedUserUrl"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
}

func (OrgUnblockUserAuditEntry) IsNode()                       {}
func (OrgUnblockUserAuditEntry) IsAuditEntry()                 {}
func (OrgUnblockUserAuditEntry) IsOrganizationAuditEntryData() {}
func (OrgUnblockUserAuditEntry) IsOrganizationAuditEntry()     {}

// Audit log entry for a org.update_default_repository_permission
type OrgUpdateDefaultRepositoryPermissionAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The new default repository permission level for the organization.
	Permission *OrgUpdateDefaultRepositoryPermissionAuditEntryPermission `json:"permission"`
	// The former default repository permission level for the organization.
	PermissionWas *OrgUpdateDefaultRepositoryPermissionAuditEntryPermission `json:"permissionWas"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
}

func (OrgUpdateDefaultRepositoryPermissionAuditEntry) IsNode()                       {}
func (OrgUpdateDefaultRepositoryPermissionAuditEntry) IsAuditEntry()                 {}
func (OrgUpdateDefaultRepositoryPermissionAuditEntry) IsOrganizationAuditEntryData() {}
func (OrgUpdateDefaultRepositoryPermissionAuditEntry) IsOrganizationAuditEntry()     {}

// Audit log entry for a org.update_member event.
type OrgUpdateMemberAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The new member permission level for the organization.
	Permission *OrgUpdateMemberAuditEntryPermission `json:"permission"`
	// The former member permission level for the organization.
	PermissionWas *OrgUpdateMemberAuditEntryPermission `json:"permissionWas"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
}

func (OrgUpdateMemberAuditEntry) IsNode()                       {}
func (OrgUpdateMemberAuditEntry) IsAuditEntry()                 {}
func (OrgUpdateMemberAuditEntry) IsOrganizationAuditEntryData() {}
func (OrgUpdateMemberAuditEntry) IsOrganizationAuditEntry()     {}

// Audit log entry for a org.update_member_repository_creation_permission event.
type OrgUpdateMemberRepositoryCreationPermissionAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// Can members create repositories in the organization.
	CanCreateRepositories *bool `json:"canCreateRepositories"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
	// The permission for visibility level of repositories for this organization.
	Visibility *OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibility `json:"visibility"`
}

func (OrgUpdateMemberRepositoryCreationPermissionAuditEntry) IsOrganizationAuditEntry()     {}
func (OrgUpdateMemberRepositoryCreationPermissionAuditEntry) IsNode()                       {}
func (OrgUpdateMemberRepositoryCreationPermissionAuditEntry) IsAuditEntry()                 {}
func (OrgUpdateMemberRepositoryCreationPermissionAuditEntry) IsOrganizationAuditEntryData() {}

// Audit log entry for a org.update_member_repository_invitation_permission event.
type OrgUpdateMemberRepositoryInvitationPermissionAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// Can outside collaborators be invited to repositories in the organization.
	CanInviteOutsideCollaboratorsToRepositories *bool `json:"canInviteOutsideCollaboratorsToRepositories"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
}

func (OrgUpdateMemberRepositoryInvitationPermissionAuditEntry) IsNode()                       {}
func (OrgUpdateMemberRepositoryInvitationPermissionAuditEntry) IsAuditEntry()                 {}
func (OrgUpdateMemberRepositoryInvitationPermissionAuditEntry) IsOrganizationAuditEntryData() {}
func (OrgUpdateMemberRepositoryInvitationPermissionAuditEntry) IsOrganizationAuditEntry()     {}

// An account on GitHub, with one or more owners, that has repositories, members and teams.
type Organization struct {
	// Determine if this repository owner has any items that can be pinned to their profile.
	AnyPinnableItems bool `json:"anyPinnableItems"`
	// Audit log entries of the organization
	AuditLog *OrganizationAuditEntryConnection `json:"auditLog"`
	// A URL pointing to the organization's public avatar.
	AvatarURL string `json:"avatarUrl"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	// Identifies the primary key from the database.
	DatabaseID *int `json:"databaseId"`
	// The organization's public profile description.
	Description *string `json:"description"`
	// The organization's public profile description rendered to HTML.
	DescriptionHTML *string `json:"descriptionHTML"`
	// The organization's public email.
	Email *string `json:"email"`
	ID    string  `json:"id"`
	// The setting value for whether the organization has an IP allow list enabled.
	IPAllowListEnabledSetting IPAllowListEnabledSettingValue `json:"ipAllowListEnabledSetting"`
	// The IP addresses that are allowed to access resources owned by the organization.
	IPAllowListEntries *IPAllowListEntryConnection `json:"ipAllowListEntries"`
	// Whether the organization has verified its profile email and website.
	IsVerified bool `json:"isVerified"`
	// Showcases a selection of repositories and gists that the profile owner has either curated or that have been selected automatically based on popularity.
	ItemShowcase *ProfileItemShowcase `json:"itemShowcase"`
	// The organization's public profile location.
	Location *string `json:"location"`
	// The organization's login name.
	Login string `json:"login"`
	// Get the status messages members of this entity have set that are either public or visible only to the organization.
	MemberStatuses *UserStatusConnection `json:"memberStatuses"`
	// A list of users who are members of this organization.
	MembersWithRole *OrganizationMemberConnection `json:"membersWithRole"`
	// The organization's public profile name.
	Name *string `json:"name"`
	// The HTTP path creating a new team
	NewTeamResourcePath string `json:"newTeamResourcePath"`
	// The HTTP URL creating a new team
	NewTeamURL string `json:"newTeamUrl"`
	// The billing email for the organization.
	OrganizationBillingEmail *string `json:"organizationBillingEmail"`
	// A list of packages under the owner.
	Packages *PackageConnection `json:"packages"`
	// A list of users who have been invited to join this organization.
	PendingMembers *UserConnection `json:"pendingMembers"`
	// A list of repositories and gists this profile owner can pin to their profile.
	PinnableItems *PinnableItemConnection `json:"pinnableItems"`
	// A list of repositories and gists this profile owner has pinned to their profile
	PinnedItems *PinnableItemConnection `json:"pinnedItems"`
	// Returns how many more items this profile owner can pin to their profile.
	PinnedItemsRemaining int `json:"pinnedItemsRemaining"`
	// Find project by number.
	Project *Project `json:"project"`
	// A list of projects under the owner.
	Projects *ProjectConnection `json:"projects"`
	// The HTTP path listing organization's projects
	ProjectsResourcePath string `json:"projectsResourcePath"`
	// The HTTP URL listing organization's projects
	ProjectsURL string `json:"projectsUrl"`
	// A list of repositories that the user owns.
	Repositories *RepositoryConnection `json:"repositories"`
	// Find Repository.
	Repository *Repository `json:"repository"`
	// When true the organization requires all members, billing managers, and outside collaborators to enable two-factor authentication.
	RequiresTwoFactorAuthentication *bool `json:"requiresTwoFactorAuthentication"`
	// The HTTP path for this organization.
	ResourcePath string `json:"resourcePath"`
	// The Organization's SAML identity providers
	SamlIdentityProvider *OrganizationIdentityProvider `json:"samlIdentityProvider"`
	// The GitHub Sponsors listing for this user.
	SponsorsListing *SponsorsListing `json:"sponsorsListing"`
	// This object's sponsorships as the maintainer.
	SponsorshipsAsMaintainer *SponsorshipConnection `json:"sponsorshipsAsMaintainer"`
	// This object's sponsorships as the sponsor.
	SponsorshipsAsSponsor *SponsorshipConnection `json:"sponsorshipsAsSponsor"`
	// Find an organization's team by its slug.
	Team *Team `json:"team"`
	// A list of teams in this organization.
	Teams *TeamConnection `json:"teams"`
	// The HTTP path listing organization's teams
	TeamsResourcePath string `json:"teamsResourcePath"`
	// The HTTP URL listing organization's teams
	TeamsURL string `json:"teamsUrl"`
	// The organization's Twitter username.
	TwitterUsername *string `json:"twitterUsername"`
	// Identifies the date and time when the object was last updated.
	UpdatedAt string `json:"updatedAt"`
	// The HTTP URL for this organization.
	URL string `json:"url"`
	// Organization is adminable by the viewer.
	ViewerCanAdminister bool `json:"viewerCanAdminister"`
	// Can the viewer pin repositories and gists to the profile?
	ViewerCanChangePinnedItems bool `json:"viewerCanChangePinnedItems"`
	// Can the current viewer create new projects on this owner.
	ViewerCanCreateProjects bool `json:"viewerCanCreateProjects"`
	// Viewer can create repositories on this organization
	ViewerCanCreateRepositories bool `json:"viewerCanCreateRepositories"`
	// Viewer can create teams on this organization.
	ViewerCanCreateTeams bool `json:"viewerCanCreateTeams"`
	// Viewer is an active member of this organization.
	ViewerIsAMember bool `json:"viewerIsAMember"`
	// The organization's public profile URL.
	WebsiteURL *string `json:"websiteUrl"`
}

func (Organization) IsSearchResultItem()         {}
func (Organization) IsPermissionGranter()        {}
func (Organization) IsSponsor()                  {}
func (Organization) IsAuditEntryActor()          {}
func (Organization) IsIPAllowListOwner()         {}
func (Organization) IsNode()                     {}
func (Organization) IsActor()                    {}
func (Organization) IsPackageOwner()             {}
func (Organization) IsProjectOwner()             {}
func (Organization) IsRepositoryOwner()          {}
func (Organization) IsUniformResourceLocatable() {}
func (Organization) IsMemberStatusable()         {}
func (Organization) IsProfileOwner()             {}
func (Organization) IsSponsorable()              {}
func (Organization) IsAssignee()                 {}

// The connection type for OrganizationAuditEntry.
type OrganizationAuditEntryConnection struct {
	// A list of edges.
	Edges []*OrganizationAuditEntryEdge `json:"edges"`
	// A list of nodes.
	Nodes []OrganizationAuditEntry `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type OrganizationAuditEntryEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node OrganizationAuditEntry `json:"node"`
}

// The connection type for Organization.
type OrganizationConnection struct {
	// A list of edges.
	Edges []*OrganizationEdge `json:"edges"`
	// A list of nodes.
	Nodes []*Organization `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type OrganizationEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *Organization `json:"node"`
}

// An Identity Provider configured to provision SAML and SCIM identities for Organizations
type OrganizationIdentityProvider struct {
	// The digest algorithm used to sign SAML requests for the Identity Provider.
	DigestMethod *string `json:"digestMethod"`
	// External Identities provisioned by this Identity Provider
	ExternalIdentities *ExternalIdentityConnection `json:"externalIdentities"`
	ID                 string                      `json:"id"`
	// The x509 certificate used by the Identity Provder to sign assertions and responses.
	IdpCertificate *string `json:"idpCertificate"`
	// The Issuer Entity ID for the SAML Identity Provider
	Issuer *string `json:"issuer"`
	// Organization this Identity Provider belongs to
	Organization *Organization `json:"organization"`
	// The signature algorithm used to sign SAML requests for the Identity Provider.
	SignatureMethod *string `json:"signatureMethod"`
	// The URL endpoint for the Identity Provider's SAML SSO.
	SsoURL *string `json:"ssoUrl"`
}

func (OrganizationIdentityProvider) IsNode() {}

// An Invitation for a user to an organization.
type OrganizationInvitation struct {
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	// The email address of the user invited to the organization.
	Email *string `json:"email"`
	ID    string  `json:"id"`
	// The type of invitation that was sent (e.g. email, user).
	InvitationType OrganizationInvitationType `json:"invitationType"`
	// The user who was invited to the organization.
	Invitee *User `json:"invitee"`
	// The user who created the invitation.
	Inviter *User `json:"inviter"`
	// The organization the invite is for
	Organization *Organization `json:"organization"`
	// The user's pending role in the organization (e.g. member, owner).
	Role OrganizationInvitationRole `json:"role"`
}

func (OrganizationInvitation) IsNode() {}

// The connection type for OrganizationInvitation.
type OrganizationInvitationConnection struct {
	// A list of edges.
	Edges []*OrganizationInvitationEdge `json:"edges"`
	// A list of nodes.
	Nodes []*OrganizationInvitation `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type OrganizationInvitationEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *OrganizationInvitation `json:"node"`
}

// The connection type for User.
type OrganizationMemberConnection struct {
	// A list of edges.
	Edges []*OrganizationMemberEdge `json:"edges"`
	// A list of nodes.
	Nodes []*User `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// Represents a user within an organization.
type OrganizationMemberEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// Whether the organization member has two factor enabled or not. Returns null if information is not available to viewer.
	HasTwoFactorEnabled *bool `json:"hasTwoFactorEnabled"`
	// The item at the end of the edge.
	Node *User `json:"node"`
	// The role this user has in the organization.
	Role *OrganizationMemberRole `json:"role"`
}

// Ordering options for organization connections.
type OrganizationOrder struct {
	// The field to order organizations by.
	Field OrganizationOrderField `json:"field"`
	// The ordering direction.
	Direction OrderDirection `json:"direction"`
}

// An organization teams hovercard context
type OrganizationTeamsHovercardContext struct {
	// A string describing this context
	Message string `json:"message"`
	// An octicon to accompany this context
	Octicon string `json:"octicon"`
	// Teams in this organization the user is a member of that are relevant
	RelevantTeams *TeamConnection `json:"relevantTeams"`
	// The path for the full team list for this user
	TeamsResourcePath string `json:"teamsResourcePath"`
	// The URL for the full team list for this user
	TeamsURL string `json:"teamsUrl"`
	// The total number of teams the user is on in the organization
	TotalTeamCount int `json:"totalTeamCount"`
}

func (OrganizationTeamsHovercardContext) IsHovercardContext() {}

// An organization list hovercard context
type OrganizationsHovercardContext struct {
	// A string describing this context
	Message string `json:"message"`
	// An octicon to accompany this context
	Octicon string `json:"octicon"`
	// Organizations this user is a member of that are relevant
	RelevantOrganizations *OrganizationConnection `json:"relevantOrganizations"`
	// The total number of organizations this user is in
	TotalOrganizationCount int `json:"totalOrganizationCount"`
}

func (OrganizationsHovercardContext) IsHovercardContext() {}

// Information for an uploaded package.
type Package struct {
	ID string `json:"id"`
	// Find the latest version for the package.
	LatestVersion *PackageVersion `json:"latestVersion"`
	// Identifies the name of the package.
	Name string `json:"name"`
	// Identifies the type of the package.
	PackageType PackageType `json:"packageType"`
	// The repository this package belongs to.
	Repository *Repository `json:"repository"`
	// Statistics about package activity.
	Statistics *PackageStatistics `json:"statistics"`
	// Find package version by version string.
	Version *PackageVersion `json:"version"`
	// list of versions for this package
	Versions *PackageVersionConnection `json:"versions"`
}

func (Package) IsNode() {}

// The connection type for Package.
type PackageConnection struct {
	// A list of edges.
	Edges []*PackageEdge `json:"edges"`
	// A list of nodes.
	Nodes []*Package `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type PackageEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *Package `json:"node"`
}

// A file in a package version.
type PackageFile struct {
	ID string `json:"id"`
	// MD5 hash of the file.
	Md5 *string `json:"md5"`
	// Name of the file.
	Name string `json:"name"`
	// The package version this file belongs to.
	PackageVersion *PackageVersion `json:"packageVersion"`
	// SHA1 hash of the file.
	Sha1 *string `json:"sha1"`
	// SHA256 hash of the file.
	Sha256 *string `json:"sha256"`
	// Size of the file in bytes.
	Size *int `json:"size"`
	// Identifies the date and time when the object was last updated.
	UpdatedAt string `json:"updatedAt"`
	// URL to download the asset.
	URL *string `json:"url"`
}

func (PackageFile) IsNode() {}

// The connection type for PackageFile.
type PackageFileConnection struct {
	// A list of edges.
	Edges []*PackageFileEdge `json:"edges"`
	// A list of nodes.
	Nodes []*PackageFile `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type PackageFileEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *PackageFile `json:"node"`
}

// Ways in which lists of package files can be ordered upon return.
type PackageFileOrder struct {
	// The field in which to order package files by.
	Field *PackageFileOrderField `json:"field"`
	// The direction in which to order package files by the specified field.
	Direction *OrderDirection `json:"direction"`
}

// Ways in which lists of packages can be ordered upon return.
type PackageOrder struct {
	// The field in which to order packages by.
	Field *PackageOrderField `json:"field"`
	// The direction in which to order packages by the specified field.
	Direction *OrderDirection `json:"direction"`
}

// Represents a object that contains package activity statistics such as downloads.
type PackageStatistics struct {
	// Number of times the package was downloaded since it was created.
	DownloadsTotalCount int `json:"downloadsTotalCount"`
}

// A version tag contains the mapping between a tag name and a version.
type PackageTag struct {
	ID string `json:"id"`
	// Identifies the tag name of the version.
	Name string `json:"name"`
	// Version that the tag is associated with.
	Version *PackageVersion `json:"version"`
}

func (PackageTag) IsNode() {}

// Information about a specific package version.
type PackageVersion struct {
	// List of files associated with this package version
	Files *PackageFileConnection `json:"files"`
	ID    string                 `json:"id"`
	// The package associated with this version.
	Package *Package `json:"package"`
	// The platform this version was built for.
	Platform *string `json:"platform"`
	// Whether or not this version is a pre-release.
	PreRelease bool `json:"preRelease"`
	// The README of this package version.
	Readme *string `json:"readme"`
	// The release associated with this package version.
	Release *Release `json:"release"`
	// Statistics about package activity.
	Statistics *PackageVersionStatistics `json:"statistics"`
	// The package version summary.
	Summary *string `json:"summary"`
	// The version string.
	Version string `json:"version"`
}

func (PackageVersion) IsNode() {}

// The connection type for PackageVersion.
type PackageVersionConnection struct {
	// A list of edges.
	Edges []*PackageVersionEdge `json:"edges"`
	// A list of nodes.
	Nodes []*PackageVersion `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type PackageVersionEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *PackageVersion `json:"node"`
}

// Ways in which lists of package versions can be ordered upon return.
type PackageVersionOrder struct {
	// The field in which to order package versions by.
	Field *PackageVersionOrderField `json:"field"`
	// The direction in which to order package versions by the specified field.
	Direction *OrderDirection `json:"direction"`
}

// Represents a object that contains package version activity statistics such as downloads.
type PackageVersionStatistics struct {
	// Number of times the package was downloaded since it was created.
	DownloadsTotalCount int `json:"downloadsTotalCount"`
}

// Information about pagination in a connection.
type PageInfo struct {
	// When paginating forwards, the cursor to continue.
	EndCursor *string `json:"endCursor"`
	// When paginating forwards, are there more items?
	HasNextPage bool `json:"hasNextPage"`
	// When paginating backwards, are there more items?
	HasPreviousPage bool `json:"hasPreviousPage"`
	// When paginating backwards, the cursor to continue.
	StartCursor *string `json:"startCursor"`
}

// A level of permission and source for a user's access to a repository.
type PermissionSource struct {
	// The organization the repository belongs to.
	Organization *Organization `json:"organization"`
	// The level of access this source has granted to the user.
	Permission DefaultRepositoryPermissionField `json:"permission"`
	// The source of this permission.
	Source PermissionGranter `json:"source"`
}

// The connection type for PinnableItem.
type PinnableItemConnection struct {
	// A list of edges.
	Edges []*PinnableItemEdge `json:"edges"`
	// A list of nodes.
	Nodes []PinnableItem `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type PinnableItemEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node PinnableItem `json:"node"`
}

// Represents a 'pinned' event on a given issue or pull request.
type PinnedEvent struct {
	// Identifies the actor who performed the event.
	Actor Actor `json:"actor"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// Identifies the issue associated with the event.
	Issue *Issue `json:"issue"`
}

func (PinnedEvent) IsPullRequestTimelineItems() {}
func (PinnedEvent) IsIssueTimelineItems()       {}
func (PinnedEvent) IsNode()                     {}

// Audit log entry for a private_repository_forking.disable event.
type PrivateRepositoryForkingDisableAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	// The HTTP path for this enterprise.
	EnterpriseResourcePath *string `json:"enterpriseResourcePath"`
	// The slug of the enterprise.
	EnterpriseSlug *string `json:"enterpriseSlug"`
	// The HTTP URL for this enterprise.
	EnterpriseURL *string `json:"enterpriseUrl"`
	ID            string  `json:"id"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The repository associated with the action
	Repository *Repository `json:"repository"`
	// The name of the repository
	RepositoryName *string `json:"repositoryName"`
	// The HTTP path for the repository
	RepositoryResourcePath *string `json:"repositoryResourcePath"`
	// The HTTP URL for the repository
	RepositoryURL *string `json:"repositoryUrl"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
}

func (PrivateRepositoryForkingDisableAuditEntry) IsNode()                       {}
func (PrivateRepositoryForkingDisableAuditEntry) IsAuditEntry()                 {}
func (PrivateRepositoryForkingDisableAuditEntry) IsEnterpriseAuditEntryData()   {}
func (PrivateRepositoryForkingDisableAuditEntry) IsOrganizationAuditEntryData() {}
func (PrivateRepositoryForkingDisableAuditEntry) IsRepositoryAuditEntryData()   {}
func (PrivateRepositoryForkingDisableAuditEntry) IsOrganizationAuditEntry()     {}

// Audit log entry for a private_repository_forking.enable event.
type PrivateRepositoryForkingEnableAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	// The HTTP path for this enterprise.
	EnterpriseResourcePath *string `json:"enterpriseResourcePath"`
	// The slug of the enterprise.
	EnterpriseSlug *string `json:"enterpriseSlug"`
	// The HTTP URL for this enterprise.
	EnterpriseURL *string `json:"enterpriseUrl"`
	ID            string  `json:"id"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The repository associated with the action
	Repository *Repository `json:"repository"`
	// The name of the repository
	RepositoryName *string `json:"repositoryName"`
	// The HTTP path for the repository
	RepositoryResourcePath *string `json:"repositoryResourcePath"`
	// The HTTP URL for the repository
	RepositoryURL *string `json:"repositoryUrl"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
}

func (PrivateRepositoryForkingEnableAuditEntry) IsOrganizationAuditEntry()     {}
func (PrivateRepositoryForkingEnableAuditEntry) IsNode()                       {}
func (PrivateRepositoryForkingEnableAuditEntry) IsAuditEntry()                 {}
func (PrivateRepositoryForkingEnableAuditEntry) IsEnterpriseAuditEntryData()   {}
func (PrivateRepositoryForkingEnableAuditEntry) IsOrganizationAuditEntryData() {}
func (PrivateRepositoryForkingEnableAuditEntry) IsRepositoryAuditEntryData()   {}

// A curatable list of repositories relating to a repository owner, which defaults to showing the most popular repositories they own.
type ProfileItemShowcase struct {
	// Whether or not the owner has pinned any repositories or gists.
	HasPinnedItems bool `json:"hasPinnedItems"`
	// The repositories and gists in the showcase. If the profile owner has any pinned items, those will be returned. Otherwise, the profile owner's popular repositories will be returned.
	Items *PinnableItemConnection `json:"items"`
}

// Projects manage issues, pull requests and notes within a project owner.
type Project struct {
	// The project's description body.
	Body *string `json:"body"`
	// The projects description body rendered to HTML.
	BodyHTML string `json:"bodyHTML"`
	// `true` if the object is closed (definition of closed may depend on type)
	Closed bool `json:"closed"`
	// Identifies the date and time when the object was closed.
	ClosedAt *string `json:"closedAt"`
	// List of columns in the project
	Columns *ProjectColumnConnection `json:"columns"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	// The actor who originally created the project.
	Creator Actor `json:"creator"`
	// Identifies the primary key from the database.
	DatabaseID *int   `json:"databaseId"`
	ID         string `json:"id"`
	// The project's name.
	Name string `json:"name"`
	// The project's number.
	Number int `json:"number"`
	// The project's owner. Currently limited to repositories, organizations, and users.
	Owner ProjectOwner `json:"owner"`
	// List of pending cards in this project
	PendingCards *ProjectCardConnection `json:"pendingCards"`
	// Project progress details.
	Progress *ProjectProgress `json:"progress"`
	// The HTTP path for this project
	ResourcePath string `json:"resourcePath"`
	// Whether the project is open or closed.
	State ProjectState `json:"state"`
	// Identifies the date and time when the object was last updated.
	UpdatedAt string `json:"updatedAt"`
	// The HTTP URL for this project
	URL string `json:"url"`
	// Check if the current viewer can update this object.
	ViewerCanUpdate bool `json:"viewerCanUpdate"`
}

func (Project) IsNode()      {}
func (Project) IsClosable()  {}
func (Project) IsUpdatable() {}

// A card in a project.
type ProjectCard struct {
	// The project column this card is associated under. A card may only belong to one
	// project column at a time. The column field will be null if the card is created
	// in a pending state and has yet to be associated with a column. Once cards are
	// associated with a column, they will not become pending in the future.
	//
	Column *ProjectColumn `json:"column"`
	// The card content item
	Content ProjectCardItem `json:"content"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	// The actor who created this card
	Creator Actor `json:"creator"`
	// Identifies the primary key from the database.
	DatabaseID *int   `json:"databaseId"`
	ID         string `json:"id"`
	// Whether the card is archived
	IsArchived bool `json:"isArchived"`
	// The card note
	Note *string `json:"note"`
	// The project that contains this card.
	Project *Project `json:"project"`
	// The HTTP path for this card
	ResourcePath string `json:"resourcePath"`
	// The state of ProjectCard
	State *ProjectCardState `json:"state"`
	// Identifies the date and time when the object was last updated.
	UpdatedAt string `json:"updatedAt"`
	// The HTTP URL for this card
	URL string `json:"url"`
}

func (ProjectCard) IsNode() {}

// The connection type for ProjectCard.
type ProjectCardConnection struct {
	// A list of edges.
	Edges []*ProjectCardEdge `json:"edges"`
	// A list of nodes.
	Nodes []*ProjectCard `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type ProjectCardEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *ProjectCard `json:"node"`
}

// A column inside a project.
type ProjectColumn struct {
	// List of cards in the column
	Cards *ProjectCardConnection `json:"cards"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	// Identifies the primary key from the database.
	DatabaseID *int   `json:"databaseId"`
	ID         string `json:"id"`
	// The project column's name.
	Name string `json:"name"`
	// The project that contains this column.
	Project *Project `json:"project"`
	// The semantic purpose of the column
	Purpose *ProjectColumnPurpose `json:"purpose"`
	// The HTTP path for this project column
	ResourcePath string `json:"resourcePath"`
	// Identifies the date and time when the object was last updated.
	UpdatedAt string `json:"updatedAt"`
	// The HTTP URL for this project column
	URL string `json:"url"`
}

func (ProjectColumn) IsNode() {}

// The connection type for ProjectColumn.
type ProjectColumnConnection struct {
	// A list of edges.
	Edges []*ProjectColumnEdge `json:"edges"`
	// A list of nodes.
	Nodes []*ProjectColumn `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type ProjectColumnEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *ProjectColumn `json:"node"`
}

// A list of projects associated with the owner.
type ProjectConnection struct {
	// A list of edges.
	Edges []*ProjectEdge `json:"edges"`
	// A list of nodes.
	Nodes []*Project `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type ProjectEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *Project `json:"node"`
}

// Ways in which lists of projects can be ordered upon return.
type ProjectOrder struct {
	// The field in which to order projects by.
	Field ProjectOrderField `json:"field"`
	// The direction in which to order projects by the specified field.
	Direction OrderDirection `json:"direction"`
}

// Project progress stats.
type ProjectProgress struct {
	// The number of done cards.
	DoneCount int `json:"doneCount"`
	// The percentage of done cards.
	DonePercentage float64 `json:"donePercentage"`
	// Whether progress tracking is enabled and cards with purpose exist for this project
	Enabled bool `json:"enabled"`
	// The number of in-progress cards.
	InProgressCount int `json:"inProgressCount"`
	// The percentage of in-progress cards.
	InProgressPercentage float64 `json:"inProgressPercentage"`
	// The number of to do cards.
	TodoCount int `json:"todoCount"`
	// The percentage of to do cards.
	TodoPercentage float64 `json:"todoPercentage"`
}

// A user's public key.
type PublicKey struct {
	// The last time this authorization was used to perform an action. Values will be null for keys not owned by the user.
	AccessedAt *string `json:"accessedAt"`
	// Identifies the date and time when the key was created. Keys created before March 5th, 2014 have inaccurate values. Values will be null for keys not owned by the user.
	CreatedAt *string `json:"createdAt"`
	// The fingerprint for this PublicKey.
	Fingerprint string `json:"fingerprint"`
	ID          string `json:"id"`
	// Whether this PublicKey is read-only or not. Values will be null for keys not owned by the user.
	IsReadOnly *bool `json:"isReadOnly"`
	// The public key string.
	Key string `json:"key"`
	// Identifies the date and time when the key was updated. Keys created before March 5th, 2014 may have inaccurate values. Values will be null for keys not owned by the user.
	UpdatedAt *string `json:"updatedAt"`
}

func (PublicKey) IsNode() {}

// The connection type for PublicKey.
type PublicKeyConnection struct {
	// A list of edges.
	Edges []*PublicKeyEdge `json:"edges"`
	// A list of nodes.
	Nodes []*PublicKey `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type PublicKeyEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *PublicKey `json:"node"`
}

// A repository pull request.
type PullRequest struct {
	// Reason that the conversation was locked.
	ActiveLockReason *LockReason `json:"activeLockReason"`
	// The number of additions in this pull request.
	Additions int `json:"additions"`
	// A list of Users assigned to this object.
	Assignees *UserConnection `json:"assignees"`
	// The actor who authored the comment.
	Author Actor `json:"author"`
	// Author's association with the subject of the comment.
	AuthorAssociation CommentAuthorAssociation `json:"authorAssociation"`
	// Identifies the base Ref associated with the pull request.
	BaseRef *Ref `json:"baseRef"`
	// Identifies the name of the base Ref associated with the pull request, even if the ref has been deleted.
	BaseRefName string `json:"baseRefName"`
	// Identifies the oid of the base ref associated with the pull request, even if the ref has been deleted.
	BaseRefOid string `json:"baseRefOid"`
	// The repository associated with this pull request's base Ref.
	BaseRepository *Repository `json:"baseRepository"`
	// The body as Markdown.
	Body string `json:"body"`
	// The body rendered to HTML.
	BodyHTML string `json:"bodyHTML"`
	// The body rendered to text.
	BodyText string `json:"bodyText"`
	// The number of changed files in this pull request.
	ChangedFiles int `json:"changedFiles"`
	// The HTTP path for the checks of this pull request.
	ChecksResourcePath string `json:"checksResourcePath"`
	// The HTTP URL for the checks of this pull request.
	ChecksURL string `json:"checksUrl"`
	// `true` if the pull request is closed
	Closed bool `json:"closed"`
	// Identifies the date and time when the object was closed.
	ClosedAt *string `json:"closedAt"`
	// A list of comments associated with the pull request.
	Comments *IssueCommentConnection `json:"comments"`
	// A list of commits present in this pull request's head branch not present in the base branch.
	Commits *PullRequestCommitConnection `json:"commits"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	// Check if this comment was created via an email reply.
	CreatedViaEmail bool `json:"createdViaEmail"`
	// Identifies the primary key from the database.
	DatabaseID *int `json:"databaseId"`
	// The number of deletions in this pull request.
	Deletions int `json:"deletions"`
	// The actor who edited this pull request's body.
	Editor Actor `json:"editor"`
	// Lists the files changed within this pull request.
	Files *PullRequestChangedFileConnection `json:"files"`
	// Identifies the head Ref associated with the pull request.
	HeadRef *Ref `json:"headRef"`
	// Identifies the name of the head Ref associated with the pull request, even if the ref has been deleted.
	HeadRefName string `json:"headRefName"`
	// Identifies the oid of the head ref associated with the pull request, even if the ref has been deleted.
	HeadRefOid string `json:"headRefOid"`
	// The repository associated with this pull request's head Ref.
	HeadRepository *Repository `json:"headRepository"`
	// The owner of the repository associated with this pull request's head Ref.
	HeadRepositoryOwner RepositoryOwner `json:"headRepositoryOwner"`
	// The hovercard information for this issue
	Hovercard *Hovercard `json:"hovercard"`
	ID        string     `json:"id"`
	// Check if this comment was edited and includes an edit with the creation data
	IncludesCreatedEdit bool `json:"includesCreatedEdit"`
	// The head and base repositories are different.
	IsCrossRepository bool `json:"isCrossRepository"`
	// Identifies if the pull request is a draft.
	IsDraft bool `json:"isDraft"`
	// Is this pull request read by the viewer
	IsReadByViewer *bool `json:"isReadByViewer"`
	// A list of labels associated with the object.
	Labels *LabelConnection `json:"labels"`
	// The moment the editor made the last edit
	LastEditedAt *string `json:"lastEditedAt"`
	// A list of latest reviews per user associated with the pull request.
	LatestOpinionatedReviews *PullRequestReviewConnection `json:"latestOpinionatedReviews"`
	// A list of latest reviews per user associated with the pull request that are not also pending review.
	LatestReviews *PullRequestReviewConnection `json:"latestReviews"`
	// `true` if the pull request is locked
	Locked bool `json:"locked"`
	// Indicates whether maintainers can modify the pull request.
	MaintainerCanModify bool `json:"maintainerCanModify"`
	// The commit that was created when this pull request was merged.
	MergeCommit *Commit `json:"mergeCommit"`
	// Whether or not the pull request can be merged based on the existence of merge conflicts.
	Mergeable MergeableState `json:"mergeable"`
	// Whether or not the pull request was merged.
	Merged bool `json:"merged"`
	// The date and time that the pull request was merged.
	MergedAt *string `json:"mergedAt"`
	// The actor who merged the pull request.
	MergedBy Actor `json:"mergedBy"`
	// Identifies the milestone associated with the pull request.
	Milestone *Milestone `json:"milestone"`
	// Identifies the pull request number.
	Number int `json:"number"`
	// A list of Users that are participating in the Pull Request conversation.
	Participants *UserConnection `json:"participants"`
	// The permalink to the pull request.
	Permalink string `json:"permalink"`
	// The commit that GitHub automatically generated to test if this pull request could be merged. This field will not return a value if the pull request is merged, or if the test merge commit is still being generated. See the `mergeable` field for more details on the mergeability of the pull request.
	PotentialMergeCommit *Commit `json:"potentialMergeCommit"`
	// List of project cards associated with this pull request.
	ProjectCards *ProjectCardConnection `json:"projectCards"`
	// Identifies when the comment was published at.
	PublishedAt *string `json:"publishedAt"`
	// A list of reactions grouped by content left on the subject.
	ReactionGroups []*ReactionGroup `json:"reactionGroups"`
	// A list of Reactions left on the Issue.
	Reactions *ReactionConnection `json:"reactions"`
	// The repository associated with this node.
	Repository *Repository `json:"repository"`
	// The HTTP path for this pull request.
	ResourcePath string `json:"resourcePath"`
	// The HTTP path for reverting this pull request.
	RevertResourcePath string `json:"revertResourcePath"`
	// The HTTP URL for reverting this pull request.
	RevertURL string `json:"revertUrl"`
	// The current status of this pull request with respect to code review.
	ReviewDecision *PullRequestReviewDecision `json:"reviewDecision"`
	// A list of review requests associated with the pull request.
	ReviewRequests *ReviewRequestConnection `json:"reviewRequests"`
	// The list of all review threads for this pull request.
	ReviewThreads *PullRequestReviewThreadConnection `json:"reviewThreads"`
	// A list of reviews associated with the pull request.
	Reviews *PullRequestReviewConnection `json:"reviews"`
	// Identifies the state of the pull request.
	State PullRequestState `json:"state"`
	// A list of reviewer suggestions based on commit history and past review comments.
	SuggestedReviewers []*SuggestedReviewer `json:"suggestedReviewers"`
	// A list of events, comments, commits, etc. associated with the pull request.
	Timeline *PullRequestTimelineConnection `json:"timeline"`
	// A list of events, comments, commits, etc. associated with the pull request.
	TimelineItems *PullRequestTimelineItemsConnection `json:"timelineItems"`
	// Identifies the pull request title.
	Title string `json:"title"`
	// Identifies the date and time when the object was last updated.
	UpdatedAt string `json:"updatedAt"`
	// The HTTP URL for this pull request.
	URL string `json:"url"`
	// A list of edits to this content.
	UserContentEdits *UserContentEditConnection `json:"userContentEdits"`
	// Whether or not the viewer can apply suggestion.
	ViewerCanApplySuggestion bool `json:"viewerCanApplySuggestion"`
	// Check if the viewer can restore the deleted head ref.
	ViewerCanDeleteHeadRef bool `json:"viewerCanDeleteHeadRef"`
	// Can user react to this subject
	ViewerCanReact bool `json:"viewerCanReact"`
	// Check if the viewer is able to change their subscription status for the repository.
	ViewerCanSubscribe bool `json:"viewerCanSubscribe"`
	// Check if the current viewer can update this object.
	ViewerCanUpdate bool `json:"viewerCanUpdate"`
	// Reasons why the current viewer can not update this comment.
	ViewerCannotUpdateReasons []CommentCannotUpdateReason `json:"viewerCannotUpdateReasons"`
	// Did the viewer author this comment.
	ViewerDidAuthor bool `json:"viewerDidAuthor"`
	// The merge body text for the viewer and method.
	ViewerMergeBodyText string `json:"viewerMergeBodyText"`
	// The merge headline text for the viewer and method.
	ViewerMergeHeadlineText string `json:"viewerMergeHeadlineText"`
	// Identifies if the viewer is watching, not watching, or ignoring the subscribable entity.
	ViewerSubscription *SubscriptionState `json:"viewerSubscription"`
}

func (PullRequest) IsIssueOrPullRequest()       {}
func (PullRequest) IsNode()                     {}
func (PullRequest) IsAssignable()               {}
func (PullRequest) IsClosable()                 {}
func (PullRequest) IsComment()                  {}
func (PullRequest) IsUpdatable()                {}
func (PullRequest) IsUpdatableComment()         {}
func (PullRequest) IsLabelable()                {}
func (PullRequest) IsLockable()                 {}
func (PullRequest) IsReactable()                {}
func (PullRequest) IsRepositoryNode()           {}
func (PullRequest) IsSubscribable()             {}
func (PullRequest) IsUniformResourceLocatable() {}
func (PullRequest) IsSearchResultItem()         {}
func (PullRequest) IsReferencedSubject()        {}
func (PullRequest) IsMilestoneItem()            {}
func (PullRequest) IsRenamedTitleSubject()      {}
func (PullRequest) IsCloser()                   {}
func (PullRequest) IsProjectCardItem()          {}

// A file changed in a pull request.
type PullRequestChangedFile struct {
	// The number of additions to the file.
	Additions int `json:"additions"`
	// The number of deletions to the file.
	Deletions int `json:"deletions"`
	// The path of the file.
	Path string `json:"path"`
	// The state of the file for the viewer.
	ViewerViewedState FileViewedState `json:"viewerViewedState"`
}

// The connection type for PullRequestChangedFile.
type PullRequestChangedFileConnection struct {
	// A list of edges.
	Edges []*PullRequestChangedFileEdge `json:"edges"`
	// A list of nodes.
	Nodes []*PullRequestChangedFile `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type PullRequestChangedFileEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *PullRequestChangedFile `json:"node"`
}

// Represents a Git commit part of a pull request.
type PullRequestCommit struct {
	// The Git commit object
	Commit *Commit `json:"commit"`
	ID     string  `json:"id"`
	// The pull request this commit belongs to
	PullRequest *PullRequest `json:"pullRequest"`
	// The HTTP path for this pull request commit
	ResourcePath string `json:"resourcePath"`
	// The HTTP URL for this pull request commit
	URL string `json:"url"`
}

func (PullRequestCommit) IsPullRequestTimelineItems() {}
func (PullRequestCommit) IsNode()                     {}
func (PullRequestCommit) IsUniformResourceLocatable() {}

// Represents a commit comment thread part of a pull request.
type PullRequestCommitCommentThread struct {
	// The comments that exist in this thread.
	Comments *CommitCommentConnection `json:"comments"`
	// The commit the comments were made on.
	Commit *Commit `json:"commit"`
	ID     string  `json:"id"`
	// The file the comments were made on.
	Path *string `json:"path"`
	// The position in the diff for the commit that the comment was made on.
	Position *int `json:"position"`
	// The pull request this commit comment thread belongs to
	PullRequest *PullRequest `json:"pullRequest"`
	// The repository associated with this node.
	Repository *Repository `json:"repository"`
}

func (PullRequestCommitCommentThread) IsPullRequestTimelineItems() {}
func (PullRequestCommitCommentThread) IsNode()                     {}
func (PullRequestCommitCommentThread) IsRepositoryNode()           {}

// The connection type for PullRequestCommit.
type PullRequestCommitConnection struct {
	// A list of edges.
	Edges []*PullRequestCommitEdge `json:"edges"`
	// A list of nodes.
	Nodes []*PullRequestCommit `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type PullRequestCommitEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *PullRequestCommit `json:"node"`
}

// The connection type for PullRequest.
type PullRequestConnection struct {
	// A list of edges.
	Edges []*PullRequestEdge `json:"edges"`
	// A list of nodes.
	Nodes []*PullRequest `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// This aggregates pull requests opened by a user within one repository.
type PullRequestContributionsByRepository struct {
	// The pull request contributions.
	Contributions *CreatedPullRequestContributionConnection `json:"contributions"`
	// The repository in which the pull requests were opened.
	Repository *Repository `json:"repository"`
}

// An edge in a connection.
type PullRequestEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *PullRequest `json:"node"`
}

// Ways in which lists of issues can be ordered upon return.
type PullRequestOrder struct {
	// The field in which to order pull requests by.
	Field PullRequestOrderField `json:"field"`
	// The direction in which to order pull requests by the specified field.
	Direction OrderDirection `json:"direction"`
}

// A review object for a given pull request.
type PullRequestReview struct {
	// The actor who authored the comment.
	Author Actor `json:"author"`
	// Author's association with the subject of the comment.
	AuthorAssociation CommentAuthorAssociation `json:"authorAssociation"`
	// Indicates whether the author of this review has push access to the repository.
	AuthorCanPushToRepository bool `json:"authorCanPushToRepository"`
	// Identifies the pull request review body.
	Body string `json:"body"`
	// The body rendered to HTML.
	BodyHTML string `json:"bodyHTML"`
	// The body of this review rendered as plain text.
	BodyText string `json:"bodyText"`
	// A list of review comments for the current pull request review.
	Comments *PullRequestReviewCommentConnection `json:"comments"`
	// Identifies the commit associated with this pull request review.
	Commit *Commit `json:"commit"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	// Check if this comment was created via an email reply.
	CreatedViaEmail bool `json:"createdViaEmail"`
	// Identifies the primary key from the database.
	DatabaseID *int `json:"databaseId"`
	// The actor who edited the comment.
	Editor Actor  `json:"editor"`
	ID     string `json:"id"`
	// Check if this comment was edited and includes an edit with the creation data
	IncludesCreatedEdit bool `json:"includesCreatedEdit"`
	// The moment the editor made the last edit
	LastEditedAt *string `json:"lastEditedAt"`
	// A list of teams that this review was made on behalf of.
	OnBehalfOf *TeamConnection `json:"onBehalfOf"`
	// Identifies when the comment was published at.
	PublishedAt *string `json:"publishedAt"`
	// Identifies the pull request associated with this pull request review.
	PullRequest *PullRequest `json:"pullRequest"`
	// A list of reactions grouped by content left on the subject.
	ReactionGroups []*ReactionGroup `json:"reactionGroups"`
	// A list of Reactions left on the Issue.
	Reactions *ReactionConnection `json:"reactions"`
	// The repository associated with this node.
	Repository *Repository `json:"repository"`
	// The HTTP path permalink for this PullRequestReview.
	ResourcePath string `json:"resourcePath"`
	// Identifies the current state of the pull request review.
	State PullRequestReviewState `json:"state"`
	// Identifies when the Pull Request Review was submitted
	SubmittedAt *string `json:"submittedAt"`
	// Identifies the date and time when the object was last updated.
	UpdatedAt string `json:"updatedAt"`
	// The HTTP URL permalink for this PullRequestReview.
	URL string `json:"url"`
	// A list of edits to this content.
	UserContentEdits *UserContentEditConnection `json:"userContentEdits"`
	// Check if the current viewer can delete this object.
	ViewerCanDelete bool `json:"viewerCanDelete"`
	// Can user react to this subject
	ViewerCanReact bool `json:"viewerCanReact"`
	// Check if the current viewer can update this object.
	ViewerCanUpdate bool `json:"viewerCanUpdate"`
	// Reasons why the current viewer can not update this comment.
	ViewerCannotUpdateReasons []CommentCannotUpdateReason `json:"viewerCannotUpdateReasons"`
	// Did the viewer author this comment.
	ViewerDidAuthor bool `json:"viewerDidAuthor"`
}

func (PullRequestReview) IsPullRequestTimelineItems() {}
func (PullRequestReview) IsPullRequestTimelineItem()  {}
func (PullRequestReview) IsNode()                     {}
func (PullRequestReview) IsComment()                  {}
func (PullRequestReview) IsDeletable()                {}
func (PullRequestReview) IsUpdatable()                {}
func (PullRequestReview) IsUpdatableComment()         {}
func (PullRequestReview) IsReactable()                {}
func (PullRequestReview) IsRepositoryNode()           {}

// A review comment associated with a given repository pull request.
type PullRequestReviewComment struct {
	// The actor who authored the comment.
	Author Actor `json:"author"`
	// Author's association with the subject of the comment.
	AuthorAssociation CommentAuthorAssociation `json:"authorAssociation"`
	// The comment body of this review comment.
	Body string `json:"body"`
	// The body rendered to HTML.
	BodyHTML string `json:"bodyHTML"`
	// The comment body of this review comment rendered as plain text.
	BodyText string `json:"bodyText"`
	// Identifies the commit associated with the comment.
	Commit *Commit `json:"commit"`
	// Identifies when the comment was created.
	CreatedAt string `json:"createdAt"`
	// Check if this comment was created via an email reply.
	CreatedViaEmail bool `json:"createdViaEmail"`
	// Identifies the primary key from the database.
	DatabaseID *int `json:"databaseId"`
	// The diff hunk to which the comment applies.
	DiffHunk string `json:"diffHunk"`
	// Identifies when the comment was created in a draft state.
	DraftedAt string `json:"draftedAt"`
	// The actor who edited the comment.
	Editor Actor  `json:"editor"`
	ID     string `json:"id"`
	// Check if this comment was edited and includes an edit with the creation data
	IncludesCreatedEdit bool `json:"includesCreatedEdit"`
	// Returns whether or not a comment has been minimized.
	IsMinimized bool `json:"isMinimized"`
	// The moment the editor made the last edit
	LastEditedAt *string `json:"lastEditedAt"`
	// Returns why the comment was minimized.
	MinimizedReason *string `json:"minimizedReason"`
	// Identifies the original commit associated with the comment.
	OriginalCommit *Commit `json:"originalCommit"`
	// The original line index in the diff to which the comment applies.
	OriginalPosition int `json:"originalPosition"`
	// Identifies when the comment body is outdated
	Outdated bool `json:"outdated"`
	// The path to which the comment applies.
	Path string `json:"path"`
	// The line index in the diff to which the comment applies.
	Position *int `json:"position"`
	// Identifies when the comment was published at.
	PublishedAt *string `json:"publishedAt"`
	// The pull request associated with this review comment.
	PullRequest *PullRequest `json:"pullRequest"`
	// The pull request review associated with this review comment.
	PullRequestReview *PullRequestReview `json:"pullRequestReview"`
	// A list of reactions grouped by content left on the subject.
	ReactionGroups []*ReactionGroup `json:"reactionGroups"`
	// A list of Reactions left on the Issue.
	Reactions *ReactionConnection `json:"reactions"`
	// The comment this is a reply to.
	ReplyTo *PullRequestReviewComment `json:"replyTo"`
	// The repository associated with this node.
	Repository *Repository `json:"repository"`
	// The HTTP path permalink for this review comment.
	ResourcePath string `json:"resourcePath"`
	// Identifies the state of the comment.
	State PullRequestReviewCommentState `json:"state"`
	// Identifies when the comment was last updated.
	UpdatedAt string `json:"updatedAt"`
	// The HTTP URL permalink for this review comment.
	URL string `json:"url"`
	// A list of edits to this content.
	UserContentEdits *UserContentEditConnection `json:"userContentEdits"`
	// Check if the current viewer can delete this object.
	ViewerCanDelete bool `json:"viewerCanDelete"`
	// Check if the current viewer can minimize this object.
	ViewerCanMinimize bool `json:"viewerCanMinimize"`
	// Can user react to this subject
	ViewerCanReact bool `json:"viewerCanReact"`
	// Check if the current viewer can update this object.
	ViewerCanUpdate bool `json:"viewerCanUpdate"`
	// Reasons why the current viewer can not update this comment.
	ViewerCannotUpdateReasons []CommentCannotUpdateReason `json:"viewerCannotUpdateReasons"`
	// Did the viewer author this comment.
	ViewerDidAuthor bool `json:"viewerDidAuthor"`
}

func (PullRequestReviewComment) IsPullRequestTimelineItem() {}
func (PullRequestReviewComment) IsNode()                    {}
func (PullRequestReviewComment) IsComment()                 {}
func (PullRequestReviewComment) IsDeletable()               {}
func (PullRequestReviewComment) IsMinimizable()             {}
func (PullRequestReviewComment) IsUpdatable()               {}
func (PullRequestReviewComment) IsUpdatableComment()        {}
func (PullRequestReviewComment) IsReactable()               {}
func (PullRequestReviewComment) IsRepositoryNode()          {}

// The connection type for PullRequestReviewComment.
type PullRequestReviewCommentConnection struct {
	// A list of edges.
	Edges []*PullRequestReviewCommentEdge `json:"edges"`
	// A list of nodes.
	Nodes []*PullRequestReviewComment `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type PullRequestReviewCommentEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *PullRequestReviewComment `json:"node"`
}

// The connection type for PullRequestReview.
type PullRequestReviewConnection struct {
	// A list of edges.
	Edges []*PullRequestReviewEdge `json:"edges"`
	// A list of nodes.
	Nodes []*PullRequestReview `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// This aggregates pull request reviews made by a user within one repository.
type PullRequestReviewContributionsByRepository struct {
	// The pull request review contributions.
	Contributions *CreatedPullRequestReviewContributionConnection `json:"contributions"`
	// The repository in which the pull request reviews were made.
	Repository *Repository `json:"repository"`
}

// An edge in a connection.
type PullRequestReviewEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *PullRequestReview `json:"node"`
}

// A threaded list of comments for a given pull request.
type PullRequestReviewThread struct {
	// A list of pull request comments associated with the thread.
	Comments *PullRequestReviewCommentConnection `json:"comments"`
	// The side of the diff on which this thread was placed.
	DiffSide DiffSide `json:"diffSide"`
	ID       string   `json:"id"`
	// Whether or not the thread has been collapsed (outdated or resolved)
	IsCollapsed bool `json:"isCollapsed"`
	// Indicates whether this thread was outdated by newer changes.
	IsOutdated bool `json:"isOutdated"`
	// Whether this thread has been resolved
	IsResolved bool `json:"isResolved"`
	// The line in the file to which this thread refers
	Line *int `json:"line"`
	// The original line in the file to which this thread refers.
	OriginalLine *int `json:"originalLine"`
	// The original start line in the file to which this thread refers (multi-line only).
	OriginalStartLine *int `json:"originalStartLine"`
	// Identifies the file path of this thread.
	Path string `json:"path"`
	// Identifies the pull request associated with this thread.
	PullRequest *PullRequest `json:"pullRequest"`
	// Identifies the repository associated with this thread.
	Repository *Repository `json:"repository"`
	// The user who resolved this thread
	ResolvedBy *User `json:"resolvedBy"`
	// The side of the diff that the first line of the thread starts on (multi-line only)
	StartDiffSide *DiffSide `json:"startDiffSide"`
	// The start line in the file to which this thread refers (multi-line only)
	StartLine *int `json:"startLine"`
	// Indicates whether the current viewer can reply to this thread.
	ViewerCanReply bool `json:"viewerCanReply"`
	// Whether or not the viewer can resolve this thread
	ViewerCanResolve bool `json:"viewerCanResolve"`
	// Whether or not the viewer can unresolve this thread
	ViewerCanUnresolve bool `json:"viewerCanUnresolve"`
}

func (PullRequestReviewThread) IsPullRequestTimelineItems() {}
func (PullRequestReviewThread) IsPullRequestTimelineItem()  {}
func (PullRequestReviewThread) IsNode()                     {}

// Review comment threads for a pull request review.
type PullRequestReviewThreadConnection struct {
	// A list of edges.
	Edges []*PullRequestReviewThreadEdge `json:"edges"`
	// A list of nodes.
	Nodes []*PullRequestReviewThread `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type PullRequestReviewThreadEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *PullRequestReviewThread `json:"node"`
}

// Represents the latest point in the pull request timeline for which the viewer has seen the pull request's commits.
type PullRequestRevisionMarker struct {
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	// The last commit the viewer has seen.
	LastSeenCommit *Commit `json:"lastSeenCommit"`
	// The pull request to which the marker belongs.
	PullRequest *PullRequest `json:"pullRequest"`
}

func (PullRequestRevisionMarker) IsPullRequestTimelineItems() {}

// The connection type for PullRequestTimelineItem.
type PullRequestTimelineConnection struct {
	// A list of edges.
	Edges []*PullRequestTimelineItemEdge `json:"edges"`
	// A list of nodes.
	Nodes []PullRequestTimelineItem `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type PullRequestTimelineItemEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node PullRequestTimelineItem `json:"node"`
}

// The connection type for PullRequestTimelineItems.
type PullRequestTimelineItemsConnection struct {
	// A list of edges.
	Edges []*PullRequestTimelineItemsEdge `json:"edges"`
	// Identifies the count of items after applying `before` and `after` filters.
	FilteredCount int `json:"filteredCount"`
	// A list of nodes.
	Nodes []PullRequestTimelineItems `json:"nodes"`
	// Identifies the count of items after applying `before`/`after` filters and `first`/`last`/`skip` slicing.
	PageCount int `json:"pageCount"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
	// Identifies the date and time when the timeline was last updated.
	UpdatedAt string `json:"updatedAt"`
}

// An edge in a connection.
type PullRequestTimelineItemsEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node PullRequestTimelineItems `json:"node"`
}

// A Git push.
type Push struct {
	ID string `json:"id"`
	// The SHA after the push
	NextSha *string `json:"nextSha"`
	// The permalink for this push.
	Permalink string `json:"permalink"`
	// The SHA before the push
	PreviousSha *string `json:"previousSha"`
	// The user who pushed
	Pusher *User `json:"pusher"`
	// The repository that was pushed to
	Repository *Repository `json:"repository"`
}

func (Push) IsNode() {}

// A team, user or app who has the ability to push to a protected branch.
type PushAllowance struct {
	// The actor that can push.
	Actor PushAllowanceActor `json:"actor"`
	// Identifies the branch protection rule associated with the allowed user or team.
	BranchProtectionRule *BranchProtectionRule `json:"branchProtectionRule"`
	ID                   string                `json:"id"`
}

func (PushAllowance) IsNode() {}

// The connection type for PushAllowance.
type PushAllowanceConnection struct {
	// A list of edges.
	Edges []*PushAllowanceEdge `json:"edges"`
	// A list of nodes.
	Nodes []*PushAllowance `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type PushAllowanceEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *PushAllowance `json:"node"`
}

// Represents the client's rate limit.
type RateLimit struct {
	// The point cost for the current query counting against the rate limit.
	Cost int `json:"cost"`
	// The maximum number of points the client is permitted to consume in a 60 minute window.
	Limit int `json:"limit"`
	// The maximum number of nodes this query may return
	NodeCount int `json:"nodeCount"`
	// The number of points remaining in the current rate limit window.
	Remaining int `json:"remaining"`
	// The time at which the current rate limit window resets in UTC epoch seconds.
	ResetAt string `json:"resetAt"`
	// The number of points used in the current rate limit window.
	Used int `json:"used"`
}

// The connection type for User.
type ReactingUserConnection struct {
	// A list of edges.
	Edges []*ReactingUserEdge `json:"edges"`
	// A list of nodes.
	Nodes []*User `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// Represents a user that's made a reaction.
type ReactingUserEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	Node   *User  `json:"node"`
	// The moment when the user made the reaction.
	ReactedAt string `json:"reactedAt"`
}

// An emoji reaction to a particular piece of content.
type Reaction struct {
	// Identifies the emoji reaction.
	Content ReactionContent `json:"content"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	// Identifies the primary key from the database.
	DatabaseID *int   `json:"databaseId"`
	ID         string `json:"id"`
	// The reactable piece of content
	Reactable Reactable `json:"reactable"`
	// Identifies the user who created this reaction.
	User *User `json:"user"`
}

func (Reaction) IsNode() {}

// A list of reactions that have been left on the subject.
type ReactionConnection struct {
	// A list of edges.
	Edges []*ReactionEdge `json:"edges"`
	// A list of nodes.
	Nodes []*Reaction `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
	// Whether or not the authenticated user has left a reaction on the subject.
	ViewerHasReacted bool `json:"viewerHasReacted"`
}

// An edge in a connection.
type ReactionEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *Reaction `json:"node"`
}

// A group of emoji reactions to a particular piece of content.
type ReactionGroup struct {
	// Identifies the emoji reaction.
	Content ReactionContent `json:"content"`
	// Identifies when the reaction was created.
	CreatedAt *string `json:"createdAt"`
	// The subject that was reacted to.
	Subject Reactable `json:"subject"`
	// Users who have reacted to the reaction subject with the emotion represented by this reaction group
	Users *ReactingUserConnection `json:"users"`
	// Whether or not the authenticated user has left a reaction on the subject.
	ViewerHasReacted bool `json:"viewerHasReacted"`
}

// Ways in which lists of reactions can be ordered upon return.
type ReactionOrder struct {
	// The field in which to order reactions by.
	Field ReactionOrderField `json:"field"`
	// The direction in which to order reactions by the specified field.
	Direction OrderDirection `json:"direction"`
}

// Represents a 'ready_for_review' event on a given pull request.
type ReadyForReviewEvent struct {
	// Identifies the actor who performed the event.
	Actor Actor `json:"actor"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// PullRequest referenced by event.
	PullRequest *PullRequest `json:"pullRequest"`
	// The HTTP path for this ready for review event.
	ResourcePath string `json:"resourcePath"`
	// The HTTP URL for this ready for review event.
	URL string `json:"url"`
}

func (ReadyForReviewEvent) IsPullRequestTimelineItems() {}
func (ReadyForReviewEvent) IsNode()                     {}
func (ReadyForReviewEvent) IsUniformResourceLocatable() {}

// Represents a Git reference.
type Ref struct {
	// A list of pull requests with this ref as the head ref.
	AssociatedPullRequests *PullRequestConnection `json:"associatedPullRequests"`
	// Branch protection rules for this ref
	BranchProtectionRule *BranchProtectionRule `json:"branchProtectionRule"`
	ID                   string                `json:"id"`
	// The ref name.
	Name string `json:"name"`
	// The ref's prefix, such as `refs/heads/` or `refs/tags/`.
	Prefix string `json:"prefix"`
	// Branch protection rules that are viewable by non-admins
	RefUpdateRule *RefUpdateRule `json:"refUpdateRule"`
	// The repository the ref belongs to.
	Repository *Repository `json:"repository"`
	// The object the ref points to. Returns null when object does not exist.
	Target GitObject `json:"target"`
}

func (Ref) IsNode() {}

// The connection type for Ref.
type RefConnection struct {
	// A list of edges.
	Edges []*RefEdge `json:"edges"`
	// A list of nodes.
	Nodes []*Ref `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type RefEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *Ref `json:"node"`
}

// Ways in which lists of git refs can be ordered upon return.
type RefOrder struct {
	// The field in which to order refs by.
	Field RefOrderField `json:"field"`
	// The direction in which to order refs by the specified field.
	Direction OrderDirection `json:"direction"`
}

// A ref update rules for a viewer.
type RefUpdateRule struct {
	// Can this branch be deleted.
	AllowsDeletions bool `json:"allowsDeletions"`
	// Are force pushes allowed on this branch.
	AllowsForcePushes bool `json:"allowsForcePushes"`
	// Identifies the protection rule pattern.
	Pattern string `json:"pattern"`
	// Number of approving reviews required to update matching branches.
	RequiredApprovingReviewCount *int `json:"requiredApprovingReviewCount"`
	// List of required status check contexts that must pass for commits to be accepted to matching branches.
	RequiredStatusCheckContexts []*string `json:"requiredStatusCheckContexts"`
	// Are merge commits prohibited from being pushed to this branch.
	RequiresLinearHistory bool `json:"requiresLinearHistory"`
	// Are commits required to be signed.
	RequiresSignatures bool `json:"requiresSignatures"`
	// Can the viewer push to the branch
	ViewerCanPush bool `json:"viewerCanPush"`
}

// Represents a 'referenced' event on a given `ReferencedSubject`.
type ReferencedEvent struct {
	// Identifies the actor who performed the event.
	Actor Actor `json:"actor"`
	// Identifies the commit associated with the 'referenced' event.
	Commit *Commit `json:"commit"`
	// Identifies the repository associated with the 'referenced' event.
	CommitRepository *Repository `json:"commitRepository"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// Reference originated in a different repository.
	IsCrossRepository bool `json:"isCrossRepository"`
	// Checks if the commit message itself references the subject. Can be false in the case of a commit comment reference.
	IsDirectReference bool `json:"isDirectReference"`
	// Object referenced by event.
	Subject ReferencedSubject `json:"subject"`
}

func (ReferencedEvent) IsPullRequestTimelineItems() {}
func (ReferencedEvent) IsIssueTimelineItems()       {}
func (ReferencedEvent) IsPullRequestTimelineItem()  {}
func (ReferencedEvent) IsIssueTimelineItem()        {}
func (ReferencedEvent) IsNode()                     {}

// Autogenerated input type of RegenerateEnterpriseIdentityProviderRecoveryCodes
type RegenerateEnterpriseIdentityProviderRecoveryCodesInput struct {
	// The ID of the enterprise on which to set an identity provider.
	EnterpriseID string `json:"enterpriseId"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of RegenerateEnterpriseIdentityProviderRecoveryCodes
type RegenerateEnterpriseIdentityProviderRecoveryCodesPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The identity provider for the enterprise.
	IdentityProvider *EnterpriseIdentityProvider `json:"identityProvider"`
}

// A release contains the content for a release.
type Release struct {
	// The author of the release
	Author *User `json:"author"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	// The description of the release.
	Description *string `json:"description"`
	// The description of this release rendered to HTML.
	DescriptionHTML *string `json:"descriptionHTML"`
	ID              string  `json:"id"`
	// Whether or not the release is a draft
	IsDraft bool `json:"isDraft"`
	// Whether or not the release is a prerelease
	IsPrerelease bool `json:"isPrerelease"`
	// The title of the release.
	Name *string `json:"name"`
	// Identifies the date and time when the release was created.
	PublishedAt *string `json:"publishedAt"`
	// List of releases assets which are dependent on this release.
	ReleaseAssets *ReleaseAssetConnection `json:"releaseAssets"`
	// The HTTP path for this issue
	ResourcePath string `json:"resourcePath"`
	// A description of the release, rendered to HTML without any links in it.
	ShortDescriptionHTML *string `json:"shortDescriptionHTML"`
	// The Git tag the release points to
	Tag *Ref `json:"tag"`
	// The name of the release's Git tag
	TagName string `json:"tagName"`
	// Identifies the date and time when the object was last updated.
	UpdatedAt string `json:"updatedAt"`
	// The HTTP URL for this issue
	URL string `json:"url"`
}

func (Release) IsNode()                     {}
func (Release) IsUniformResourceLocatable() {}

// A release asset contains the content for a release asset.
type ReleaseAsset struct {
	// The asset's content-type
	ContentType string `json:"contentType"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	// The number of times this asset was downloaded
	DownloadCount int `json:"downloadCount"`
	// Identifies the URL where you can download the release asset via the browser.
	DownloadURL string `json:"downloadUrl"`
	ID          string `json:"id"`
	// Identifies the title of the release asset.
	Name string `json:"name"`
	// Release that the asset is associated with
	Release *Release `json:"release"`
	// The size (in bytes) of the asset
	Size int `json:"size"`
	// Identifies the date and time when the object was last updated.
	UpdatedAt string `json:"updatedAt"`
	// The user that performed the upload
	UploadedBy *User `json:"uploadedBy"`
	// Identifies the URL of the release asset.
	URL string `json:"url"`
}

func (ReleaseAsset) IsNode() {}

// The connection type for ReleaseAsset.
type ReleaseAssetConnection struct {
	// A list of edges.
	Edges []*ReleaseAssetEdge `json:"edges"`
	// A list of nodes.
	Nodes []*ReleaseAsset `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type ReleaseAssetEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *ReleaseAsset `json:"node"`
}

// The connection type for Release.
type ReleaseConnection struct {
	// A list of edges.
	Edges []*ReleaseEdge `json:"edges"`
	// A list of nodes.
	Nodes []*Release `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type ReleaseEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *Release `json:"node"`
}

// Ways in which lists of releases can be ordered upon return.
type ReleaseOrder struct {
	// The field in which to order releases by.
	Field ReleaseOrderField `json:"field"`
	// The direction in which to order releases by the specified field.
	Direction OrderDirection `json:"direction"`
}

// Autogenerated input type of RemoveAssigneesFromAssignable
type RemoveAssigneesFromAssignableInput struct {
	// The id of the assignable object to remove assignees from.
	AssignableID string `json:"assignableId"`
	// The id of users to remove as assignees.
	AssigneeIds []string `json:"assigneeIds"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of RemoveAssigneesFromAssignable
type RemoveAssigneesFromAssignablePayload struct {
	// The item that was unassigned.
	Assignable Assignable `json:"assignable"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated input type of RemoveEnterpriseAdmin
type RemoveEnterpriseAdminInput struct {
	// The Enterprise ID from which to remove the administrator.
	EnterpriseID string `json:"enterpriseId"`
	// The login of the user to remove as an administrator.
	Login string `json:"login"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of RemoveEnterpriseAdmin
type RemoveEnterpriseAdminPayload struct {
	// The user who was removed as an administrator.
	Admin *User `json:"admin"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The updated enterprise.
	Enterprise *Enterprise `json:"enterprise"`
	// A message confirming the result of removing an administrator.
	Message *string `json:"message"`
	// The viewer performing the mutation.
	Viewer *User `json:"viewer"`
}

// Autogenerated input type of RemoveEnterpriseIdentityProvider
type RemoveEnterpriseIdentityProviderInput struct {
	// The ID of the enterprise from which to remove the identity provider.
	EnterpriseID string `json:"enterpriseId"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of RemoveEnterpriseIdentityProvider
type RemoveEnterpriseIdentityProviderPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The identity provider that was removed from the enterprise.
	IdentityProvider *EnterpriseIdentityProvider `json:"identityProvider"`
}

// Autogenerated input type of RemoveEnterpriseOrganization
type RemoveEnterpriseOrganizationInput struct {
	// The ID of the enterprise from which the organization should be removed.
	EnterpriseID string `json:"enterpriseId"`
	// The ID of the organization to remove from the enterprise.
	OrganizationID string `json:"organizationId"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of RemoveEnterpriseOrganization
type RemoveEnterpriseOrganizationPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The updated enterprise.
	Enterprise *Enterprise `json:"enterprise"`
	// The organization that was removed from the enterprise.
	Organization *Organization `json:"organization"`
	// The viewer performing the mutation.
	Viewer *User `json:"viewer"`
}

// Autogenerated input type of RemoveLabelsFromLabelable
type RemoveLabelsFromLabelableInput struct {
	// The id of the Labelable to remove labels from.
	LabelableID string `json:"labelableId"`
	// The ids of labels to remove.
	LabelIds []string `json:"labelIds"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of RemoveLabelsFromLabelable
type RemoveLabelsFromLabelablePayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The Labelable the labels were removed from.
	Labelable Labelable `json:"labelable"`
}

// Autogenerated input type of RemoveOutsideCollaborator
type RemoveOutsideCollaboratorInput struct {
	// The ID of the outside collaborator to remove.
	UserID string `json:"userId"`
	// The ID of the organization to remove the outside collaborator from.
	OrganizationID string `json:"organizationId"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of RemoveOutsideCollaborator
type RemoveOutsideCollaboratorPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The user that was removed as an outside collaborator.
	RemovedUser *User `json:"removedUser"`
}

// Autogenerated input type of RemoveReaction
type RemoveReactionInput struct {
	// The Node ID of the subject to modify.
	SubjectID string `json:"subjectId"`
	// The name of the emoji reaction to remove.
	Content ReactionContent `json:"content"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of RemoveReaction
type RemoveReactionPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The reaction object.
	Reaction *Reaction `json:"reaction"`
	// The reactable subject.
	Subject Reactable `json:"subject"`
}

// Autogenerated input type of RemoveStar
type RemoveStarInput struct {
	// The Starrable ID to unstar.
	StarrableID string `json:"starrableId"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of RemoveStar
type RemoveStarPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The starrable.
	Starrable Starrable `json:"starrable"`
}

// Represents a 'removed_from_project' event on a given issue or pull request.
type RemovedFromProjectEvent struct {
	// Identifies the actor who performed the event.
	Actor Actor `json:"actor"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	// Identifies the primary key from the database.
	DatabaseID *int   `json:"databaseId"`
	ID         string `json:"id"`
}

func (RemovedFromProjectEvent) IsNode()                     {}
func (RemovedFromProjectEvent) IsPullRequestTimelineItems() {}
func (RemovedFromProjectEvent) IsIssueTimelineItems()       {}

// Represents a 'renamed' event on a given issue or pull request
type RenamedTitleEvent struct {
	// Identifies the actor who performed the event.
	Actor Actor `json:"actor"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	// Identifies the current title of the issue or pull request.
	CurrentTitle string `json:"currentTitle"`
	ID           string `json:"id"`
	// Identifies the previous title of the issue or pull request.
	PreviousTitle string `json:"previousTitle"`
	// Subject that was renamed.
	Subject RenamedTitleSubject `json:"subject"`
}

func (RenamedTitleEvent) IsPullRequestTimelineItems() {}
func (RenamedTitleEvent) IsIssueTimelineItems()       {}
func (RenamedTitleEvent) IsPullRequestTimelineItem()  {}
func (RenamedTitleEvent) IsIssueTimelineItem()        {}
func (RenamedTitleEvent) IsNode()                     {}

// Autogenerated input type of ReopenIssue
type ReopenIssueInput struct {
	// ID of the issue to be opened.
	IssueID string `json:"issueId"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of ReopenIssue
type ReopenIssuePayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The issue that was opened.
	Issue *Issue `json:"issue"`
}

// Autogenerated input type of ReopenPullRequest
type ReopenPullRequestInput struct {
	// ID of the pull request to be reopened.
	PullRequestID string `json:"pullRequestId"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of ReopenPullRequest
type ReopenPullRequestPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The pull request that was reopened.
	PullRequest *PullRequest `json:"pullRequest"`
}

// Represents a 'reopened' event on any `Closable`.
type ReopenedEvent struct {
	// Identifies the actor who performed the event.
	Actor Actor `json:"actor"`
	// Object that was reopened.
	Closable Closable `json:"closable"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
}

func (ReopenedEvent) IsPullRequestTimelineItems() {}
func (ReopenedEvent) IsIssueTimelineItems()       {}
func (ReopenedEvent) IsNode()                     {}
func (ReopenedEvent) IsPullRequestTimelineItem()  {}
func (ReopenedEvent) IsIssueTimelineItem()        {}

// Audit log entry for a repo.access event.
type RepoAccessAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The repository associated with the action
	Repository *Repository `json:"repository"`
	// The name of the repository
	RepositoryName *string `json:"repositoryName"`
	// The HTTP path for the repository
	RepositoryResourcePath *string `json:"repositoryResourcePath"`
	// The HTTP URL for the repository
	RepositoryURL *string `json:"repositoryUrl"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
	// The visibility of the repository
	Visibility *RepoAccessAuditEntryVisibility `json:"visibility"`
}

func (RepoAccessAuditEntry) IsNode()                       {}
func (RepoAccessAuditEntry) IsAuditEntry()                 {}
func (RepoAccessAuditEntry) IsOrganizationAuditEntryData() {}
func (RepoAccessAuditEntry) IsRepositoryAuditEntryData()   {}
func (RepoAccessAuditEntry) IsOrganizationAuditEntry()     {}

// Audit log entry for a repo.add_member event.
type RepoAddMemberAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The repository associated with the action
	Repository *Repository `json:"repository"`
	// The name of the repository
	RepositoryName *string `json:"repositoryName"`
	// The HTTP path for the repository
	RepositoryResourcePath *string `json:"repositoryResourcePath"`
	// The HTTP URL for the repository
	RepositoryURL *string `json:"repositoryUrl"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
	// The visibility of the repository
	Visibility *RepoAddMemberAuditEntryVisibility `json:"visibility"`
}

func (RepoAddMemberAuditEntry) IsOrganizationAuditEntry()     {}
func (RepoAddMemberAuditEntry) IsNode()                       {}
func (RepoAddMemberAuditEntry) IsAuditEntry()                 {}
func (RepoAddMemberAuditEntry) IsOrganizationAuditEntryData() {}
func (RepoAddMemberAuditEntry) IsRepositoryAuditEntryData()   {}

// Audit log entry for a repo.add_topic event.
type RepoAddTopicAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The repository associated with the action
	Repository *Repository `json:"repository"`
	// The name of the repository
	RepositoryName *string `json:"repositoryName"`
	// The HTTP path for the repository
	RepositoryResourcePath *string `json:"repositoryResourcePath"`
	// The HTTP URL for the repository
	RepositoryURL *string `json:"repositoryUrl"`
	// The name of the topic added to the repository
	Topic *Topic `json:"topic"`
	// The name of the topic added to the repository
	TopicName *string `json:"topicName"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
}

func (RepoAddTopicAuditEntry) IsNode()                       {}
func (RepoAddTopicAuditEntry) IsAuditEntry()                 {}
func (RepoAddTopicAuditEntry) IsRepositoryAuditEntryData()   {}
func (RepoAddTopicAuditEntry) IsOrganizationAuditEntryData() {}
func (RepoAddTopicAuditEntry) IsTopicAuditEntryData()        {}
func (RepoAddTopicAuditEntry) IsOrganizationAuditEntry()     {}

// Audit log entry for a repo.archived event.
type RepoArchivedAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The repository associated with the action
	Repository *Repository `json:"repository"`
	// The name of the repository
	RepositoryName *string `json:"repositoryName"`
	// The HTTP path for the repository
	RepositoryResourcePath *string `json:"repositoryResourcePath"`
	// The HTTP URL for the repository
	RepositoryURL *string `json:"repositoryUrl"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
	// The visibility of the repository
	Visibility *RepoArchivedAuditEntryVisibility `json:"visibility"`
}

func (RepoArchivedAuditEntry) IsOrganizationAuditEntry()     {}
func (RepoArchivedAuditEntry) IsNode()                       {}
func (RepoArchivedAuditEntry) IsAuditEntry()                 {}
func (RepoArchivedAuditEntry) IsRepositoryAuditEntryData()   {}
func (RepoArchivedAuditEntry) IsOrganizationAuditEntryData() {}

// Audit log entry for a repo.change_merge_setting event.
type RepoChangeMergeSettingAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// Whether the change was to enable (true) or disable (false) the merge type
	IsEnabled *bool `json:"isEnabled"`
	// The merge method affected by the change
	MergeType *RepoChangeMergeSettingAuditEntryMergeType `json:"mergeType"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The repository associated with the action
	Repository *Repository `json:"repository"`
	// The name of the repository
	RepositoryName *string `json:"repositoryName"`
	// The HTTP path for the repository
	RepositoryResourcePath *string `json:"repositoryResourcePath"`
	// The HTTP URL for the repository
	RepositoryURL *string `json:"repositoryUrl"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
}

func (RepoChangeMergeSettingAuditEntry) IsNode()                       {}
func (RepoChangeMergeSettingAuditEntry) IsAuditEntry()                 {}
func (RepoChangeMergeSettingAuditEntry) IsRepositoryAuditEntryData()   {}
func (RepoChangeMergeSettingAuditEntry) IsOrganizationAuditEntryData() {}
func (RepoChangeMergeSettingAuditEntry) IsOrganizationAuditEntry()     {}

// Audit log entry for a repo.config.disable_anonymous_git_access event.
type RepoConfigDisableAnonymousGitAccessAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The repository associated with the action
	Repository *Repository `json:"repository"`
	// The name of the repository
	RepositoryName *string `json:"repositoryName"`
	// The HTTP path for the repository
	RepositoryResourcePath *string `json:"repositoryResourcePath"`
	// The HTTP URL for the repository
	RepositoryURL *string `json:"repositoryUrl"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
}

func (RepoConfigDisableAnonymousGitAccessAuditEntry) IsNode()                       {}
func (RepoConfigDisableAnonymousGitAccessAuditEntry) IsAuditEntry()                 {}
func (RepoConfigDisableAnonymousGitAccessAuditEntry) IsOrganizationAuditEntryData() {}
func (RepoConfigDisableAnonymousGitAccessAuditEntry) IsRepositoryAuditEntryData()   {}
func (RepoConfigDisableAnonymousGitAccessAuditEntry) IsOrganizationAuditEntry()     {}

// Audit log entry for a repo.config.disable_collaborators_only event.
type RepoConfigDisableCollaboratorsOnlyAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The repository associated with the action
	Repository *Repository `json:"repository"`
	// The name of the repository
	RepositoryName *string `json:"repositoryName"`
	// The HTTP path for the repository
	RepositoryResourcePath *string `json:"repositoryResourcePath"`
	// The HTTP URL for the repository
	RepositoryURL *string `json:"repositoryUrl"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
}

func (RepoConfigDisableCollaboratorsOnlyAuditEntry) IsOrganizationAuditEntry()     {}
func (RepoConfigDisableCollaboratorsOnlyAuditEntry) IsNode()                       {}
func (RepoConfigDisableCollaboratorsOnlyAuditEntry) IsAuditEntry()                 {}
func (RepoConfigDisableCollaboratorsOnlyAuditEntry) IsOrganizationAuditEntryData() {}
func (RepoConfigDisableCollaboratorsOnlyAuditEntry) IsRepositoryAuditEntryData()   {}

// Audit log entry for a repo.config.disable_contributors_only event.
type RepoConfigDisableContributorsOnlyAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The repository associated with the action
	Repository *Repository `json:"repository"`
	// The name of the repository
	RepositoryName *string `json:"repositoryName"`
	// The HTTP path for the repository
	RepositoryResourcePath *string `json:"repositoryResourcePath"`
	// The HTTP URL for the repository
	RepositoryURL *string `json:"repositoryUrl"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
}

func (RepoConfigDisableContributorsOnlyAuditEntry) IsNode()                       {}
func (RepoConfigDisableContributorsOnlyAuditEntry) IsAuditEntry()                 {}
func (RepoConfigDisableContributorsOnlyAuditEntry) IsOrganizationAuditEntryData() {}
func (RepoConfigDisableContributorsOnlyAuditEntry) IsRepositoryAuditEntryData()   {}
func (RepoConfigDisableContributorsOnlyAuditEntry) IsOrganizationAuditEntry()     {}

// Audit log entry for a repo.config.disable_sockpuppet_disallowed event.
type RepoConfigDisableSockpuppetDisallowedAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The repository associated with the action
	Repository *Repository `json:"repository"`
	// The name of the repository
	RepositoryName *string `json:"repositoryName"`
	// The HTTP path for the repository
	RepositoryResourcePath *string `json:"repositoryResourcePath"`
	// The HTTP URL for the repository
	RepositoryURL *string `json:"repositoryUrl"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
}

func (RepoConfigDisableSockpuppetDisallowedAuditEntry) IsOrganizationAuditEntry()     {}
func (RepoConfigDisableSockpuppetDisallowedAuditEntry) IsNode()                       {}
func (RepoConfigDisableSockpuppetDisallowedAuditEntry) IsAuditEntry()                 {}
func (RepoConfigDisableSockpuppetDisallowedAuditEntry) IsOrganizationAuditEntryData() {}
func (RepoConfigDisableSockpuppetDisallowedAuditEntry) IsRepositoryAuditEntryData()   {}

// Audit log entry for a repo.config.enable_anonymous_git_access event.
type RepoConfigEnableAnonymousGitAccessAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The repository associated with the action
	Repository *Repository `json:"repository"`
	// The name of the repository
	RepositoryName *string `json:"repositoryName"`
	// The HTTP path for the repository
	RepositoryResourcePath *string `json:"repositoryResourcePath"`
	// The HTTP URL for the repository
	RepositoryURL *string `json:"repositoryUrl"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
}

func (RepoConfigEnableAnonymousGitAccessAuditEntry) IsOrganizationAuditEntry()     {}
func (RepoConfigEnableAnonymousGitAccessAuditEntry) IsNode()                       {}
func (RepoConfigEnableAnonymousGitAccessAuditEntry) IsAuditEntry()                 {}
func (RepoConfigEnableAnonymousGitAccessAuditEntry) IsOrganizationAuditEntryData() {}
func (RepoConfigEnableAnonymousGitAccessAuditEntry) IsRepositoryAuditEntryData()   {}

// Audit log entry for a repo.config.enable_collaborators_only event.
type RepoConfigEnableCollaboratorsOnlyAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The repository associated with the action
	Repository *Repository `json:"repository"`
	// The name of the repository
	RepositoryName *string `json:"repositoryName"`
	// The HTTP path for the repository
	RepositoryResourcePath *string `json:"repositoryResourcePath"`
	// The HTTP URL for the repository
	RepositoryURL *string `json:"repositoryUrl"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
}

func (RepoConfigEnableCollaboratorsOnlyAuditEntry) IsNode()                       {}
func (RepoConfigEnableCollaboratorsOnlyAuditEntry) IsAuditEntry()                 {}
func (RepoConfigEnableCollaboratorsOnlyAuditEntry) IsOrganizationAuditEntryData() {}
func (RepoConfigEnableCollaboratorsOnlyAuditEntry) IsRepositoryAuditEntryData()   {}
func (RepoConfigEnableCollaboratorsOnlyAuditEntry) IsOrganizationAuditEntry()     {}

// Audit log entry for a repo.config.enable_contributors_only event.
type RepoConfigEnableContributorsOnlyAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The repository associated with the action
	Repository *Repository `json:"repository"`
	// The name of the repository
	RepositoryName *string `json:"repositoryName"`
	// The HTTP path for the repository
	RepositoryResourcePath *string `json:"repositoryResourcePath"`
	// The HTTP URL for the repository
	RepositoryURL *string `json:"repositoryUrl"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
}

func (RepoConfigEnableContributorsOnlyAuditEntry) IsNode()                       {}
func (RepoConfigEnableContributorsOnlyAuditEntry) IsAuditEntry()                 {}
func (RepoConfigEnableContributorsOnlyAuditEntry) IsOrganizationAuditEntryData() {}
func (RepoConfigEnableContributorsOnlyAuditEntry) IsRepositoryAuditEntryData()   {}
func (RepoConfigEnableContributorsOnlyAuditEntry) IsOrganizationAuditEntry()     {}

// Audit log entry for a repo.config.enable_sockpuppet_disallowed event.
type RepoConfigEnableSockpuppetDisallowedAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The repository associated with the action
	Repository *Repository `json:"repository"`
	// The name of the repository
	RepositoryName *string `json:"repositoryName"`
	// The HTTP path for the repository
	RepositoryResourcePath *string `json:"repositoryResourcePath"`
	// The HTTP URL for the repository
	RepositoryURL *string `json:"repositoryUrl"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
}

func (RepoConfigEnableSockpuppetDisallowedAuditEntry) IsOrganizationAuditEntry()     {}
func (RepoConfigEnableSockpuppetDisallowedAuditEntry) IsNode()                       {}
func (RepoConfigEnableSockpuppetDisallowedAuditEntry) IsAuditEntry()                 {}
func (RepoConfigEnableSockpuppetDisallowedAuditEntry) IsOrganizationAuditEntryData() {}
func (RepoConfigEnableSockpuppetDisallowedAuditEntry) IsRepositoryAuditEntryData()   {}

// Audit log entry for a repo.config.lock_anonymous_git_access event.
type RepoConfigLockAnonymousGitAccessAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The repository associated with the action
	Repository *Repository `json:"repository"`
	// The name of the repository
	RepositoryName *string `json:"repositoryName"`
	// The HTTP path for the repository
	RepositoryResourcePath *string `json:"repositoryResourcePath"`
	// The HTTP URL for the repository
	RepositoryURL *string `json:"repositoryUrl"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
}

func (RepoConfigLockAnonymousGitAccessAuditEntry) IsNode()                       {}
func (RepoConfigLockAnonymousGitAccessAuditEntry) IsAuditEntry()                 {}
func (RepoConfigLockAnonymousGitAccessAuditEntry) IsOrganizationAuditEntryData() {}
func (RepoConfigLockAnonymousGitAccessAuditEntry) IsRepositoryAuditEntryData()   {}
func (RepoConfigLockAnonymousGitAccessAuditEntry) IsOrganizationAuditEntry()     {}

// Audit log entry for a repo.config.unlock_anonymous_git_access event.
type RepoConfigUnlockAnonymousGitAccessAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The repository associated with the action
	Repository *Repository `json:"repository"`
	// The name of the repository
	RepositoryName *string `json:"repositoryName"`
	// The HTTP path for the repository
	RepositoryResourcePath *string `json:"repositoryResourcePath"`
	// The HTTP URL for the repository
	RepositoryURL *string `json:"repositoryUrl"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
}

func (RepoConfigUnlockAnonymousGitAccessAuditEntry) IsOrganizationAuditEntry()     {}
func (RepoConfigUnlockAnonymousGitAccessAuditEntry) IsNode()                       {}
func (RepoConfigUnlockAnonymousGitAccessAuditEntry) IsAuditEntry()                 {}
func (RepoConfigUnlockAnonymousGitAccessAuditEntry) IsOrganizationAuditEntryData() {}
func (RepoConfigUnlockAnonymousGitAccessAuditEntry) IsRepositoryAuditEntryData()   {}

// Audit log entry for a repo.create event.
type RepoCreateAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	// The name of the parent repository for this forked repository.
	ForkParentName *string `json:"forkParentName"`
	// The name of the root repository for this netork.
	ForkSourceName *string `json:"forkSourceName"`
	ID             string  `json:"id"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The repository associated with the action
	Repository *Repository `json:"repository"`
	// The name of the repository
	RepositoryName *string `json:"repositoryName"`
	// The HTTP path for the repository
	RepositoryResourcePath *string `json:"repositoryResourcePath"`
	// The HTTP URL for the repository
	RepositoryURL *string `json:"repositoryUrl"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
	// The visibility of the repository
	Visibility *RepoCreateAuditEntryVisibility `json:"visibility"`
}

func (RepoCreateAuditEntry) IsNode()                       {}
func (RepoCreateAuditEntry) IsAuditEntry()                 {}
func (RepoCreateAuditEntry) IsRepositoryAuditEntryData()   {}
func (RepoCreateAuditEntry) IsOrganizationAuditEntryData() {}
func (RepoCreateAuditEntry) IsOrganizationAuditEntry()     {}

// Audit log entry for a repo.destroy event.
type RepoDestroyAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The repository associated with the action
	Repository *Repository `json:"repository"`
	// The name of the repository
	RepositoryName *string `json:"repositoryName"`
	// The HTTP path for the repository
	RepositoryResourcePath *string `json:"repositoryResourcePath"`
	// The HTTP URL for the repository
	RepositoryURL *string `json:"repositoryUrl"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
	// The visibility of the repository
	Visibility *RepoDestroyAuditEntryVisibility `json:"visibility"`
}

func (RepoDestroyAuditEntry) IsNode()                       {}
func (RepoDestroyAuditEntry) IsAuditEntry()                 {}
func (RepoDestroyAuditEntry) IsRepositoryAuditEntryData()   {}
func (RepoDestroyAuditEntry) IsOrganizationAuditEntryData() {}
func (RepoDestroyAuditEntry) IsOrganizationAuditEntry()     {}

// Audit log entry for a repo.remove_member event.
type RepoRemoveMemberAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The repository associated with the action
	Repository *Repository `json:"repository"`
	// The name of the repository
	RepositoryName *string `json:"repositoryName"`
	// The HTTP path for the repository
	RepositoryResourcePath *string `json:"repositoryResourcePath"`
	// The HTTP URL for the repository
	RepositoryURL *string `json:"repositoryUrl"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
	// The visibility of the repository
	Visibility *RepoRemoveMemberAuditEntryVisibility `json:"visibility"`
}

func (RepoRemoveMemberAuditEntry) IsNode()                       {}
func (RepoRemoveMemberAuditEntry) IsAuditEntry()                 {}
func (RepoRemoveMemberAuditEntry) IsOrganizationAuditEntryData() {}
func (RepoRemoveMemberAuditEntry) IsRepositoryAuditEntryData()   {}
func (RepoRemoveMemberAuditEntry) IsOrganizationAuditEntry()     {}

// Audit log entry for a repo.remove_topic event.
type RepoRemoveTopicAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The repository associated with the action
	Repository *Repository `json:"repository"`
	// The name of the repository
	RepositoryName *string `json:"repositoryName"`
	// The HTTP path for the repository
	RepositoryResourcePath *string `json:"repositoryResourcePath"`
	// The HTTP URL for the repository
	RepositoryURL *string `json:"repositoryUrl"`
	// The name of the topic added to the repository
	Topic *Topic `json:"topic"`
	// The name of the topic added to the repository
	TopicName *string `json:"topicName"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
}

func (RepoRemoveTopicAuditEntry) IsNode()                       {}
func (RepoRemoveTopicAuditEntry) IsAuditEntry()                 {}
func (RepoRemoveTopicAuditEntry) IsRepositoryAuditEntryData()   {}
func (RepoRemoveTopicAuditEntry) IsOrganizationAuditEntryData() {}
func (RepoRemoveTopicAuditEntry) IsTopicAuditEntryData()        {}
func (RepoRemoveTopicAuditEntry) IsOrganizationAuditEntry()     {}

// A repository contains the content for a project.
type Repository struct {
	// A list of users that can be assigned to issues in this repository.
	AssignableUsers *UserConnection `json:"assignableUsers"`
	// A list of branch protection rules for this repository.
	BranchProtectionRules *BranchProtectionRuleConnection `json:"branchProtectionRules"`
	// Returns the code of conduct for this repository
	CodeOfConduct *CodeOfConduct `json:"codeOfConduct"`
	// A list of collaborators associated with the repository.
	Collaborators *RepositoryCollaboratorConnection `json:"collaborators"`
	// A list of commit comments associated with the repository.
	CommitComments *CommitCommentConnection `json:"commitComments"`
	// Returns a list of contact links associated to the repository
	ContactLinks []*RepositoryContactLink `json:"contactLinks"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	// Identifies the primary key from the database.
	DatabaseID *int `json:"databaseId"`
	// The Ref associated with the repository's default branch.
	DefaultBranchRef *Ref `json:"defaultBranchRef"`
	// Whether or not branches are automatically deleted when merged in this repository.
	DeleteBranchOnMerge bool `json:"deleteBranchOnMerge"`
	// A list of deploy keys that are on this repository.
	DeployKeys *DeployKeyConnection `json:"deployKeys"`
	// Deployments associated with the repository
	Deployments *DeploymentConnection `json:"deployments"`
	// The description of the repository.
	Description *string `json:"description"`
	// The description of the repository rendered to HTML.
	DescriptionHTML string `json:"descriptionHTML"`
	// The number of kilobytes this repository occupies on disk.
	DiskUsage *int `json:"diskUsage"`
	// Returns how many forks there are of this repository in the whole network.
	ForkCount int `json:"forkCount"`
	// A list of direct forked repositories.
	Forks *RepositoryConnection `json:"forks"`
	// The funding links for this repository
	FundingLinks []*FundingLink `json:"fundingLinks"`
	// Indicates if the repository has issues feature enabled.
	HasIssuesEnabled bool `json:"hasIssuesEnabled"`
	// Indicates if the repository has the Projects feature enabled.
	HasProjectsEnabled bool `json:"hasProjectsEnabled"`
	// Indicates if the repository has wiki feature enabled.
	HasWikiEnabled bool `json:"hasWikiEnabled"`
	// The repository's URL.
	HomepageURL *string `json:"homepageUrl"`
	ID          string  `json:"id"`
	// Indicates if the repository is unmaintained.
	IsArchived bool `json:"isArchived"`
	// Returns true if blank issue creation is allowed
	IsBlankIssuesEnabled bool `json:"isBlankIssuesEnabled"`
	// Returns whether or not this repository disabled.
	IsDisabled bool `json:"isDisabled"`
	// Returns whether or not this repository is empty.
	IsEmpty bool `json:"isEmpty"`
	// Identifies if the repository is a fork.
	IsFork bool `json:"isFork"`
	// Indicates if a repository is either owned by an organization, or is a private fork of an organization repository.
	IsInOrganization bool `json:"isInOrganization"`
	// Indicates if the repository has been locked or not.
	IsLocked bool `json:"isLocked"`
	// Identifies if the repository is a mirror.
	IsMirror bool `json:"isMirror"`
	// Identifies if the repository is private.
	IsPrivate bool `json:"isPrivate"`
	// Returns true if this repository has a security policy
	IsSecurityPolicyEnabled *bool `json:"isSecurityPolicyEnabled"`
	// Identifies if the repository is a template that can be used to generate new repositories.
	IsTemplate bool `json:"isTemplate"`
	// Is this repository a user configuration repository?
	IsUserConfigurationRepository bool `json:"isUserConfigurationRepository"`
	// Returns a single issue from the current repository by number.
	Issue *Issue `json:"issue"`
	// Returns a single issue-like object from the current repository by number.
	IssueOrPullRequest IssueOrPullRequest `json:"issueOrPullRequest"`
	// Returns a list of issue templates associated to the repository
	IssueTemplates []*IssueTemplate `json:"issueTemplates"`
	// A list of issues that have been opened in the repository.
	Issues *IssueConnection `json:"issues"`
	// Returns a single label by name
	Label *Label `json:"label"`
	// A list of labels associated with the repository.
	Labels *LabelConnection `json:"labels"`
	// A list containing a breakdown of the language composition of the repository.
	Languages *LanguageConnection `json:"languages"`
	// The license associated with the repository
	LicenseInfo *License `json:"licenseInfo"`
	// The reason the repository has been locked.
	LockReason *RepositoryLockReason `json:"lockReason"`
	// A list of Users that can be mentioned in the context of the repository.
	MentionableUsers *UserConnection `json:"mentionableUsers"`
	// Whether or not PRs are merged with a merge commit on this repository.
	MergeCommitAllowed bool `json:"mergeCommitAllowed"`
	// Returns a single milestone from the current repository by number.
	Milestone *Milestone `json:"milestone"`
	// A list of milestones associated with the repository.
	Milestones *MilestoneConnection `json:"milestones"`
	// The repository's original mirror URL.
	MirrorURL *string `json:"mirrorUrl"`
	// The name of the repository.
	Name string `json:"name"`
	// The repository's name with owner.
	NameWithOwner string `json:"nameWithOwner"`
	// A Git object in the repository
	Object GitObject `json:"object"`
	// The image used to represent this repository in Open Graph data.
	OpenGraphImageURL string `json:"openGraphImageUrl"`
	// The User owner of the repository.
	Owner RepositoryOwner `json:"owner"`
	// A list of packages under the owner.
	Packages *PackageConnection `json:"packages"`
	// The repository parent, if this is a fork.
	Parent *Repository `json:"parent"`
	// The primary language of the repository's code.
	PrimaryLanguage *Language `json:"primaryLanguage"`
	// Find project by number.
	Project *Project `json:"project"`
	// A list of projects under the owner.
	Projects *ProjectConnection `json:"projects"`
	// The HTTP path listing the repository's projects
	ProjectsResourcePath string `json:"projectsResourcePath"`
	// The HTTP URL listing the repository's projects
	ProjectsURL string `json:"projectsUrl"`
	// Returns a single pull request from the current repository by number.
	PullRequest *PullRequest `json:"pullRequest"`
	// A list of pull requests that have been opened in the repository.
	PullRequests *PullRequestConnection `json:"pullRequests"`
	// Identifies when the repository was last pushed to.
	PushedAt *string `json:"pushedAt"`
	// Whether or not rebase-merging is enabled on this repository.
	RebaseMergeAllowed bool `json:"rebaseMergeAllowed"`
	// Fetch a given ref from the repository
	Ref *Ref `json:"ref"`
	// Fetch a list of refs from the repository
	Refs *RefConnection `json:"refs"`
	// Lookup a single release given various criteria.
	Release *Release `json:"release"`
	// List of releases which are dependent on this repository.
	Releases *ReleaseConnection `json:"releases"`
	// A list of applied repository-topic associations for this repository.
	RepositoryTopics *RepositoryTopicConnection `json:"repositoryTopics"`
	// The HTTP path for this repository
	ResourcePath string `json:"resourcePath"`
	// The security policy URL.
	SecurityPolicyURL *string `json:"securityPolicyUrl"`
	// A description of the repository, rendered to HTML without any links in it.
	ShortDescriptionHTML string `json:"shortDescriptionHTML"`
	// Whether or not squash-merging is enabled on this repository.
	SquashMergeAllowed bool `json:"squashMergeAllowed"`
	// The SSH URL to clone this repository
	SSHURL string `json:"sshUrl"`
	// Returns a count of how many stargazers there are on this object
	//
	StargazerCount int `json:"stargazerCount"`
	// A list of users who have starred this starrable.
	Stargazers *StargazerConnection `json:"stargazers"`
	// Returns a list of all submodules in this repository parsed from the .gitmodules file as of the default branch's HEAD commit.
	Submodules *SubmoduleConnection `json:"submodules"`
	// Temporary authentication token for cloning this repository.
	TempCloneToken *string `json:"tempCloneToken"`
	// The repository from which this repository was generated, if any.
	TemplateRepository *Repository `json:"templateRepository"`
	// Identifies the date and time when the object was last updated.
	UpdatedAt string `json:"updatedAt"`
	// The HTTP URL for this repository
	URL string `json:"url"`
	// Whether this repository has a custom image to use with Open Graph as opposed to being represented by the owner's avatar.
	UsesCustomOpenGraphImage bool `json:"usesCustomOpenGraphImage"`
	// Indicates whether the viewer has admin permissions on this repository.
	ViewerCanAdminister bool `json:"viewerCanAdminister"`
	// Can the current viewer create new projects on this owner.
	ViewerCanCreateProjects bool `json:"viewerCanCreateProjects"`
	// Check if the viewer is able to change their subscription status for the repository.
	ViewerCanSubscribe bool `json:"viewerCanSubscribe"`
	// Indicates whether the viewer can update the topics of this repository.
	ViewerCanUpdateTopics bool `json:"viewerCanUpdateTopics"`
	// The last commit email for the viewer.
	ViewerDefaultCommitEmail *string `json:"viewerDefaultCommitEmail"`
	// The last used merge method by the viewer or the default for the repository.
	ViewerDefaultMergeMethod PullRequestMergeMethod `json:"viewerDefaultMergeMethod"`
	// Returns a boolean indicating whether the viewing user has starred this starrable.
	ViewerHasStarred bool `json:"viewerHasStarred"`
	// The users permission level on the repository. Will return null if authenticated as an GitHub App.
	ViewerPermission *RepositoryPermission `json:"viewerPermission"`
	// A list of emails this viewer can commit with.
	ViewerPossibleCommitEmails []string `json:"viewerPossibleCommitEmails"`
	// Identifies if the viewer is watching, not watching, or ignoring the subscribable entity.
	ViewerSubscription *SubscriptionState `json:"viewerSubscription"`
	// A list of vulnerability alerts that are on this repository.
	VulnerabilityAlerts *RepositoryVulnerabilityAlertConnection `json:"vulnerabilityAlerts"`
	// A list of users watching the repository.
	Watchers *UserConnection `json:"watchers"`
}

func (Repository) IsSearchResultItem()         {}
func (Repository) IsPermissionGranter()        {}
func (Repository) IsNode()                     {}
func (Repository) IsProjectOwner()             {}
func (Repository) IsPackageOwner()             {}
func (Repository) IsSubscribable()             {}
func (Repository) IsStarrable()                {}
func (Repository) IsUniformResourceLocatable() {}
func (Repository) IsRepositoryInfo()           {}
func (Repository) IsPinnableItem()             {}

// The connection type for User.
type RepositoryCollaboratorConnection struct {
	// A list of edges.
	Edges []*RepositoryCollaboratorEdge `json:"edges"`
	// A list of nodes.
	Nodes []*User `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// Represents a user who is a collaborator of a repository.
type RepositoryCollaboratorEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	Node   *User  `json:"node"`
	// The permission the user has on the repository.
	//
	// **Upcoming Change on 2020-10-01 UTC**
	// **Description:** Type for `permission` will change from `RepositoryPermission!` to `String`.
	// **Reason:** This field may return additional values
	//
	Permission RepositoryPermission `json:"permission"`
	// A list of sources for the user's access to the repository.
	PermissionSources []*PermissionSource `json:"permissionSources"`
}

// A list of repositories owned by the subject.
type RepositoryConnection struct {
	// A list of edges.
	Edges []*RepositoryEdge `json:"edges"`
	// A list of nodes.
	Nodes []*Repository `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
	// The total size in kilobytes of all repositories in the connection.
	TotalDiskUsage int `json:"totalDiskUsage"`
}

// A repository contact link.
type RepositoryContactLink struct {
	// The contact link purpose.
	About string `json:"about"`
	// The contact link name.
	Name string `json:"name"`
	// The contact link URL.
	URL string `json:"url"`
}

// An edge in a connection.
type RepositoryEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *Repository `json:"node"`
}

// An invitation for a user to be added to a repository.
type RepositoryInvitation struct {
	// The email address that received the invitation.
	Email *string `json:"email"`
	ID    string  `json:"id"`
	// The user who received the invitation.
	Invitee *User `json:"invitee"`
	// The user who created the invitation.
	Inviter *User `json:"inviter"`
	// The permalink for this repository invitation.
	Permalink string `json:"permalink"`
	// The permission granted on this repository by this invitation.
	//
	// **Upcoming Change on 2020-10-01 UTC**
	// **Description:** Type for `permission` will change from `RepositoryPermission!` to `String`.
	// **Reason:** This field may return additional values
	//
	Permission RepositoryPermission `json:"permission"`
	// The Repository the user is invited to.
	Repository RepositoryInfo `json:"repository"`
}

func (RepositoryInvitation) IsNode() {}

// The connection type for RepositoryInvitation.
type RepositoryInvitationConnection struct {
	// A list of edges.
	Edges []*RepositoryInvitationEdge `json:"edges"`
	// A list of nodes.
	Nodes []*RepositoryInvitation `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type RepositoryInvitationEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *RepositoryInvitation `json:"node"`
}

// Ordering options for repository invitation connections.
type RepositoryInvitationOrder struct {
	// The field to order repository invitations by.
	Field RepositoryInvitationOrderField `json:"field"`
	// The ordering direction.
	Direction OrderDirection `json:"direction"`
}

// Ordering options for repository connections
type RepositoryOrder struct {
	// The field to order repositories by.
	Field RepositoryOrderField `json:"field"`
	// The ordering direction.
	Direction OrderDirection `json:"direction"`
}

// A repository-topic connects a repository to a topic.
type RepositoryTopic struct {
	ID string `json:"id"`
	// The HTTP path for this repository-topic.
	ResourcePath string `json:"resourcePath"`
	// The topic.
	Topic *Topic `json:"topic"`
	// The HTTP URL for this repository-topic.
	URL string `json:"url"`
}

func (RepositoryTopic) IsNode()                     {}
func (RepositoryTopic) IsUniformResourceLocatable() {}

// The connection type for RepositoryTopic.
type RepositoryTopicConnection struct {
	// A list of edges.
	Edges []*RepositoryTopicEdge `json:"edges"`
	// A list of nodes.
	Nodes []*RepositoryTopic `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type RepositoryTopicEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *RepositoryTopic `json:"node"`
}

// Audit log entry for a repository_visibility_change.disable event.
type RepositoryVisibilityChangeDisableAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	// The HTTP path for this enterprise.
	EnterpriseResourcePath *string `json:"enterpriseResourcePath"`
	// The slug of the enterprise.
	EnterpriseSlug *string `json:"enterpriseSlug"`
	// The HTTP URL for this enterprise.
	EnterpriseURL *string `json:"enterpriseUrl"`
	ID            string  `json:"id"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
}

func (RepositoryVisibilityChangeDisableAuditEntry) IsOrganizationAuditEntry()     {}
func (RepositoryVisibilityChangeDisableAuditEntry) IsNode()                       {}
func (RepositoryVisibilityChangeDisableAuditEntry) IsAuditEntry()                 {}
func (RepositoryVisibilityChangeDisableAuditEntry) IsEnterpriseAuditEntryData()   {}
func (RepositoryVisibilityChangeDisableAuditEntry) IsOrganizationAuditEntryData() {}

// Audit log entry for a repository_visibility_change.enable event.
type RepositoryVisibilityChangeEnableAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	// The HTTP path for this enterprise.
	EnterpriseResourcePath *string `json:"enterpriseResourcePath"`
	// The slug of the enterprise.
	EnterpriseSlug *string `json:"enterpriseSlug"`
	// The HTTP URL for this enterprise.
	EnterpriseURL *string `json:"enterpriseUrl"`
	ID            string  `json:"id"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
}

func (RepositoryVisibilityChangeEnableAuditEntry) IsOrganizationAuditEntry()     {}
func (RepositoryVisibilityChangeEnableAuditEntry) IsNode()                       {}
func (RepositoryVisibilityChangeEnableAuditEntry) IsAuditEntry()                 {}
func (RepositoryVisibilityChangeEnableAuditEntry) IsEnterpriseAuditEntryData()   {}
func (RepositoryVisibilityChangeEnableAuditEntry) IsOrganizationAuditEntryData() {}

// A alert for a repository with an affected vulnerability.
type RepositoryVulnerabilityAlert struct {
	// When was the alert created?
	CreatedAt string `json:"createdAt"`
	// The reason the alert was dismissed
	DismissReason *string `json:"dismissReason"`
	// When was the alert dimissed?
	DismissedAt *string `json:"dismissedAt"`
	// The user who dismissed the alert
	Dismisser *User  `json:"dismisser"`
	ID        string `json:"id"`
	// The associated repository
	Repository *Repository `json:"repository"`
	// The associated security advisory
	SecurityAdvisory *SecurityAdvisory `json:"securityAdvisory"`
	// The associated security vulnerablity
	SecurityVulnerability *SecurityVulnerability `json:"securityVulnerability"`
	// The vulnerable manifest filename
	VulnerableManifestFilename string `json:"vulnerableManifestFilename"`
	// The vulnerable manifest path
	VulnerableManifestPath string `json:"vulnerableManifestPath"`
	// The vulnerable requirements
	VulnerableRequirements *string `json:"vulnerableRequirements"`
}

func (RepositoryVulnerabilityAlert) IsNode()           {}
func (RepositoryVulnerabilityAlert) IsRepositoryNode() {}

// The connection type for RepositoryVulnerabilityAlert.
type RepositoryVulnerabilityAlertConnection struct {
	// A list of edges.
	Edges []*RepositoryVulnerabilityAlertEdge `json:"edges"`
	// A list of nodes.
	Nodes []*RepositoryVulnerabilityAlert `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type RepositoryVulnerabilityAlertEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *RepositoryVulnerabilityAlert `json:"node"`
}

// Autogenerated input type of RequestReviews
type RequestReviewsInput struct {
	// The Node ID of the pull request to modify.
	PullRequestID string `json:"pullRequestId"`
	// The Node IDs of the user to request.
	UserIds []string `json:"userIds"`
	// The Node IDs of the team to request.
	TeamIds []string `json:"teamIds"`
	// Add users to the set rather than replace.
	Union *bool `json:"union"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of RequestReviews
type RequestReviewsPayload struct {
	// Identifies the actor who performed the event.
	Actor Actor `json:"actor"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The pull request that is getting requests.
	PullRequest *PullRequest `json:"pullRequest"`
	// The edge from the pull request to the requested reviewers.
	RequestedReviewersEdge *UserEdge `json:"requestedReviewersEdge"`
}

// Autogenerated input type of RerequestCheckSuite
type RerequestCheckSuiteInput struct {
	// The Node ID of the repository.
	RepositoryID string `json:"repositoryId"`
	// The Node ID of the check suite.
	CheckSuiteID string `json:"checkSuiteId"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of RerequestCheckSuite
type RerequestCheckSuitePayload struct {
	// The requested check suite.
	CheckSuite *CheckSuite `json:"checkSuite"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated input type of ResolveReviewThread
type ResolveReviewThreadInput struct {
	// The ID of the thread to resolve
	ThreadID string `json:"threadId"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of ResolveReviewThread
type ResolveReviewThreadPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The thread to resolve.
	Thread *PullRequestReviewThread `json:"thread"`
}

// Represents a private contribution a user made on GitHub.
type RestrictedContribution struct {
	// Whether this contribution is associated with a record you do not have access to. For
	// example, your own 'first issue' contribution may have been made on a repository you can no
	// longer access.
	//
	IsRestricted bool `json:"isRestricted"`
	// When this contribution was made.
	OccurredAt string `json:"occurredAt"`
	// The HTTP path for this contribution.
	ResourcePath string `json:"resourcePath"`
	// The HTTP URL for this contribution.
	URL string `json:"url"`
	// The user who made this contribution.
	//
	User *User `json:"user"`
}

func (RestrictedContribution) IsCreatedIssueOrRestrictedContribution()       {}
func (RestrictedContribution) IsContribution()                               {}
func (RestrictedContribution) IsCreatedPullRequestOrRestrictedContribution() {}
func (RestrictedContribution) IsCreatedRepositoryOrRestrictedContribution()  {}

// A team or user who has the ability to dismiss a review on a protected branch.
type ReviewDismissalAllowance struct {
	// The actor that can dismiss.
	Actor ReviewDismissalAllowanceActor `json:"actor"`
	// Identifies the branch protection rule associated with the allowed user or team.
	BranchProtectionRule *BranchProtectionRule `json:"branchProtectionRule"`
	ID                   string                `json:"id"`
}

func (ReviewDismissalAllowance) IsNode() {}

// The connection type for ReviewDismissalAllowance.
type ReviewDismissalAllowanceConnection struct {
	// A list of edges.
	Edges []*ReviewDismissalAllowanceEdge `json:"edges"`
	// A list of nodes.
	Nodes []*ReviewDismissalAllowance `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type ReviewDismissalAllowanceEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *ReviewDismissalAllowance `json:"node"`
}

// Represents a 'review_dismissed' event on a given issue or pull request.
type ReviewDismissedEvent struct {
	// Identifies the actor who performed the event.
	Actor Actor `json:"actor"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	// Identifies the primary key from the database.
	DatabaseID *int `json:"databaseId"`
	// Identifies the optional message associated with the 'review_dismissed' event.
	DismissalMessage *string `json:"dismissalMessage"`
	// Identifies the optional message associated with the event, rendered to HTML.
	DismissalMessageHTML *string `json:"dismissalMessageHTML"`
	ID                   string  `json:"id"`
	// Identifies the previous state of the review with the 'review_dismissed' event.
	PreviousReviewState PullRequestReviewState `json:"previousReviewState"`
	// PullRequest referenced by event.
	PullRequest *PullRequest `json:"pullRequest"`
	// Identifies the commit which caused the review to become stale.
	PullRequestCommit *PullRequestCommit `json:"pullRequestCommit"`
	// The HTTP path for this review dismissed event.
	ResourcePath string `json:"resourcePath"`
	// Identifies the review associated with the 'review_dismissed' event.
	Review *PullRequestReview `json:"review"`
	// The HTTP URL for this review dismissed event.
	URL string `json:"url"`
}

func (ReviewDismissedEvent) IsPullRequestTimelineItems() {}
func (ReviewDismissedEvent) IsPullRequestTimelineItem()  {}
func (ReviewDismissedEvent) IsNode()                     {}
func (ReviewDismissedEvent) IsUniformResourceLocatable() {}

// A request for a user to review a pull request.
type ReviewRequest struct {
	// Whether this request was created for a code owner
	AsCodeOwner bool `json:"asCodeOwner"`
	// Identifies the primary key from the database.
	DatabaseID *int   `json:"databaseId"`
	ID         string `json:"id"`
	// Identifies the pull request associated with this review request.
	PullRequest *PullRequest `json:"pullRequest"`
	// The reviewer that is requested.
	RequestedReviewer RequestedReviewer `json:"requestedReviewer"`
}

func (ReviewRequest) IsNode() {}

// The connection type for ReviewRequest.
type ReviewRequestConnection struct {
	// A list of edges.
	Edges []*ReviewRequestEdge `json:"edges"`
	// A list of nodes.
	Nodes []*ReviewRequest `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type ReviewRequestEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *ReviewRequest `json:"node"`
}

// Represents an 'review_request_removed' event on a given pull request.
type ReviewRequestRemovedEvent struct {
	// Identifies the actor who performed the event.
	Actor Actor `json:"actor"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// PullRequest referenced by event.
	PullRequest *PullRequest `json:"pullRequest"`
	// Identifies the reviewer whose review request was removed.
	RequestedReviewer RequestedReviewer `json:"requestedReviewer"`
}

func (ReviewRequestRemovedEvent) IsPullRequestTimelineItems() {}
func (ReviewRequestRemovedEvent) IsPullRequestTimelineItem()  {}
func (ReviewRequestRemovedEvent) IsNode()                     {}

// Represents an 'review_requested' event on a given pull request.
type ReviewRequestedEvent struct {
	// Identifies the actor who performed the event.
	Actor Actor `json:"actor"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// PullRequest referenced by event.
	PullRequest *PullRequest `json:"pullRequest"`
	// Identifies the reviewer whose review was requested.
	RequestedReviewer RequestedReviewer `json:"requestedReviewer"`
}

func (ReviewRequestedEvent) IsPullRequestTimelineItems() {}
func (ReviewRequestedEvent) IsPullRequestTimelineItem()  {}
func (ReviewRequestedEvent) IsNode()                     {}

// A hovercard context with a message describing the current code review state of the pull
// request.
//
type ReviewStatusHovercardContext struct {
	// A string describing this context
	Message string `json:"message"`
	// An octicon to accompany this context
	Octicon string `json:"octicon"`
	// The current status of the pull request with respect to code review.
	ReviewDecision *PullRequestReviewDecision `json:"reviewDecision"`
}

func (ReviewStatusHovercardContext) IsHovercardContext() {}

// A Saved Reply is text a user can use to reply quickly.
type SavedReply struct {
	// The body of the saved reply.
	Body string `json:"body"`
	// The saved reply body rendered to HTML.
	BodyHTML string `json:"bodyHTML"`
	// Identifies the primary key from the database.
	DatabaseID *int   `json:"databaseId"`
	ID         string `json:"id"`
	// The title of the saved reply.
	Title string `json:"title"`
	// The user that saved this reply.
	User Actor `json:"user"`
}

func (SavedReply) IsNode() {}

// The connection type for SavedReply.
type SavedReplyConnection struct {
	// A list of edges.
	Edges []*SavedReplyEdge `json:"edges"`
	// A list of nodes.
	Nodes []*SavedReply `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type SavedReplyEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *SavedReply `json:"node"`
}

// Ordering options for saved reply connections.
type SavedReplyOrder struct {
	// The field to order saved replies by.
	Field SavedReplyOrderField `json:"field"`
	// The ordering direction.
	Direction OrderDirection `json:"direction"`
}

// A list of results that matched against a search query.
type SearchResultItemConnection struct {
	// The number of pieces of code that matched the search query.
	CodeCount int `json:"codeCount"`
	// A list of edges.
	Edges []*SearchResultItemEdge `json:"edges"`
	// The number of issues that matched the search query.
	IssueCount int `json:"issueCount"`
	// A list of nodes.
	Nodes []SearchResultItem `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// The number of repositories that matched the search query.
	RepositoryCount int `json:"repositoryCount"`
	// The number of users that matched the search query.
	UserCount int `json:"userCount"`
	// The number of wiki pages that matched the search query.
	WikiCount int `json:"wikiCount"`
}

// An edge in a connection.
type SearchResultItemEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node SearchResultItem `json:"node"`
	// Text matches on the result found.
	TextMatches []*TextMatch `json:"textMatches"`
}

// A GitHub Security Advisory
type SecurityAdvisory struct {
	// Identifies the primary key from the database.
	DatabaseID *int `json:"databaseId"`
	// This is a long plaintext description of the advisory
	Description string `json:"description"`
	// The GitHub Security Advisory ID
	GhsaID string `json:"ghsaId"`
	ID     string `json:"id"`
	// A list of identifiers for this advisory
	Identifiers []*SecurityAdvisoryIdentifier `json:"identifiers"`
	// The organization that originated the advisory
	Origin string `json:"origin"`
	// The permalink for the advisory
	Permalink *string `json:"permalink"`
	// When the advisory was published
	PublishedAt string `json:"publishedAt"`
	// A list of references for this advisory
	References []*SecurityAdvisoryReference `json:"references"`
	// The severity of the advisory
	Severity SecurityAdvisorySeverity `json:"severity"`
	// A short plaintext summary of the advisory
	Summary string `json:"summary"`
	// When the advisory was last updated
	UpdatedAt string `json:"updatedAt"`
	// Vulnerabilities associated with this Advisory
	Vulnerabilities *SecurityVulnerabilityConnection `json:"vulnerabilities"`
	// When the advisory was withdrawn, if it has been withdrawn
	WithdrawnAt *string `json:"withdrawnAt"`
}

func (SecurityAdvisory) IsNode() {}

// The connection type for SecurityAdvisory.
type SecurityAdvisoryConnection struct {
	// A list of edges.
	Edges []*SecurityAdvisoryEdge `json:"edges"`
	// A list of nodes.
	Nodes []*SecurityAdvisory `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type SecurityAdvisoryEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *SecurityAdvisory `json:"node"`
}

// A GitHub Security Advisory Identifier
type SecurityAdvisoryIdentifier struct {
	// The identifier type, e.g. GHSA, CVE
	Type string `json:"type"`
	// The identifier
	Value string `json:"value"`
}

// An advisory identifier to filter results on.
type SecurityAdvisoryIdentifierFilter struct {
	// The identifier type.
	Type SecurityAdvisoryIdentifierType `json:"type"`
	// The identifier string. Supports exact or partial matching.
	Value string `json:"value"`
}

// Ordering options for security advisory connections
type SecurityAdvisoryOrder struct {
	// The field to order security advisories by.
	Field SecurityAdvisoryOrderField `json:"field"`
	// The ordering direction.
	Direction OrderDirection `json:"direction"`
}

// An individual package
type SecurityAdvisoryPackage struct {
	// The ecosystem the package belongs to, e.g. RUBYGEMS, NPM
	Ecosystem SecurityAdvisoryEcosystem `json:"ecosystem"`
	// The package name
	Name string `json:"name"`
}

// An individual package version
type SecurityAdvisoryPackageVersion struct {
	// The package name or version
	Identifier string `json:"identifier"`
}

// A GitHub Security Advisory Reference
type SecurityAdvisoryReference struct {
	// A publicly accessible reference
	URL string `json:"url"`
}

// An individual vulnerability within an Advisory
type SecurityVulnerability struct {
	// The Advisory associated with this Vulnerability
	Advisory *SecurityAdvisory `json:"advisory"`
	// The first version containing a fix for the vulnerability
	FirstPatchedVersion *SecurityAdvisoryPackageVersion `json:"firstPatchedVersion"`
	// A description of the vulnerable package
	Package *SecurityAdvisoryPackage `json:"package"`
	// The severity of the vulnerability within this package
	Severity SecurityAdvisorySeverity `json:"severity"`
	// When the vulnerability was last updated
	UpdatedAt string `json:"updatedAt"`
	// A string that describes the vulnerable package versions.
	// This string follows a basic syntax with a few forms.
	// + `= 0.2.0` denotes a single vulnerable version.
	// + `<= 1.0.8` denotes a version range up to and including the specified version
	// + `< 0.1.11` denotes a version range up to, but excluding, the specified version
	// + `>= 4.3.0, < 4.3.5` denotes a version range with a known minimum and maximum version.
	// + `>= 0.0.1` denotes a version range with a known minimum, but no known maximum
	//
	VulnerableVersionRange string `json:"vulnerableVersionRange"`
}

// The connection type for SecurityVulnerability.
type SecurityVulnerabilityConnection struct {
	// A list of edges.
	Edges []*SecurityVulnerabilityEdge `json:"edges"`
	// A list of nodes.
	Nodes []*SecurityVulnerability `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type SecurityVulnerabilityEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *SecurityVulnerability `json:"node"`
}

// Ordering options for security vulnerability connections
type SecurityVulnerabilityOrder struct {
	// The field to order security vulnerabilities by.
	Field SecurityVulnerabilityOrderField `json:"field"`
	// The ordering direction.
	Direction OrderDirection `json:"direction"`
}

// Autogenerated input type of SetEnterpriseIdentityProvider
type SetEnterpriseIdentityProviderInput struct {
	// The ID of the enterprise on which to set an identity provider.
	EnterpriseID string `json:"enterpriseId"`
	// The URL endpoint for the identity provider's SAML SSO.
	SsoURL string `json:"ssoUrl"`
	// The Issuer Entity ID for the SAML identity provider
	Issuer *string `json:"issuer"`
	// The x509 certificate used by the identity provider to sign assertions and responses.
	IdpCertificate string `json:"idpCertificate"`
	// The signature algorithm used to sign SAML requests for the identity provider.
	SignatureMethod SamlSignatureAlgorithm `json:"signatureMethod"`
	// The digest algorithm used to sign SAML requests for the identity provider.
	DigestMethod SamlDigestAlgorithm `json:"digestMethod"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of SetEnterpriseIdentityProvider
type SetEnterpriseIdentityProviderPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The identity provider for the enterprise.
	IdentityProvider *EnterpriseIdentityProvider `json:"identityProvider"`
}

// Represents an S/MIME signature on a Commit or Tag.
type SmimeSignature struct {
	// Email used to sign this object.
	Email string `json:"email"`
	// True if the signature is valid and verified by GitHub.
	IsValid bool `json:"isValid"`
	// Payload for GPG signing object. Raw ODB object without the signature header.
	Payload string `json:"payload"`
	// ASCII-armored signature header from object.
	Signature string `json:"signature"`
	// GitHub user corresponding to the email signing this commit.
	Signer *User `json:"signer"`
	// The state of this signature. `VALID` if signature is valid and verified by GitHub, otherwise represents reason why signature is considered invalid.
	State GitSignatureState `json:"state"`
	// True if the signature was made with GitHub's signing key.
	WasSignedByGitHub bool `json:"wasSignedByGitHub"`
}

func (SmimeSignature) IsGitSignature() {}

// A GitHub Sponsors listing.
type SponsorsListing struct {
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	// The full description of the listing.
	FullDescription string `json:"fullDescription"`
	// The full description of the listing rendered to HTML.
	FullDescriptionHTML string `json:"fullDescriptionHTML"`
	ID                  string `json:"id"`
	// The listing's full name.
	Name string `json:"name"`
	// The short description of the listing.
	ShortDescription string `json:"shortDescription"`
	// The short name of the listing.
	Slug string `json:"slug"`
	// The published tiers for this GitHub Sponsors listing.
	Tiers *SponsorsTierConnection `json:"tiers"`
}

func (SponsorsListing) IsNode() {}

// A GitHub Sponsors tier associated with a GitHub Sponsors listing.
type SponsorsTier struct {
	// SponsorsTier information only visible to users that can administer the associated Sponsors listing.
	AdminInfo *SponsorsTierAdminInfo `json:"adminInfo"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	// The description of the tier.
	Description string `json:"description"`
	// The tier description rendered to HTML
	DescriptionHTML string `json:"descriptionHTML"`
	ID              string `json:"id"`
	// How much this tier costs per month in cents.
	MonthlyPriceInCents int `json:"monthlyPriceInCents"`
	// How much this tier costs per month in dollars.
	MonthlyPriceInDollars int `json:"monthlyPriceInDollars"`
	// The name of the tier.
	Name string `json:"name"`
	// The sponsors listing that this tier belongs to.
	SponsorsListing *SponsorsListing `json:"sponsorsListing"`
	// Identifies the date and time when the object was last updated.
	UpdatedAt string `json:"updatedAt"`
}

func (SponsorsTier) IsNode() {}

// SponsorsTier information only visible to users that can administer the associated Sponsors listing.
type SponsorsTierAdminInfo struct {
	// The sponsorships associated with this tier.
	Sponsorships *SponsorshipConnection `json:"sponsorships"`
}

// The connection type for SponsorsTier.
type SponsorsTierConnection struct {
	// A list of edges.
	Edges []*SponsorsTierEdge `json:"edges"`
	// A list of nodes.
	Nodes []*SponsorsTier `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type SponsorsTierEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *SponsorsTier `json:"node"`
}

// Ordering options for Sponsors tiers connections.
type SponsorsTierOrder struct {
	// The field to order tiers by.
	Field SponsorsTierOrderField `json:"field"`
	// The ordering direction.
	Direction OrderDirection `json:"direction"`
}

// A sponsorship relationship between a sponsor and a maintainer
type Sponsorship struct {
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// The entity that is being sponsored
	Maintainer *User `json:"maintainer"`
	// The privacy level for this sponsorship.
	PrivacyLevel SponsorshipPrivacy `json:"privacyLevel"`
	// The user that is sponsoring. Returns null if the sponsorship is private or if sponsor is not a user.
	Sponsor *User `json:"sponsor"`
	// The user or organization that is sponsoring. Returns null if the sponsorship is private.
	SponsorEntity Sponsor `json:"sponsorEntity"`
	// The entity that is being sponsored
	Sponsorable Sponsorable `json:"sponsorable"`
	// The associated sponsorship tier
	Tier *SponsorsTier `json:"tier"`
}

func (Sponsorship) IsNode() {}

// The connection type for Sponsorship.
type SponsorshipConnection struct {
	// A list of edges.
	Edges []*SponsorshipEdge `json:"edges"`
	// A list of nodes.
	Nodes []*Sponsorship `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type SponsorshipEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *Sponsorship `json:"node"`
}

// Ordering options for sponsorship connections.
type SponsorshipOrder struct {
	// The field to order sponsorship by.
	Field SponsorshipOrderField `json:"field"`
	// The ordering direction.
	Direction OrderDirection `json:"direction"`
}

// Ways in which star connections can be ordered.
type StarOrder struct {
	// The field in which to order nodes by.
	Field StarOrderField `json:"field"`
	// The direction in which to order nodes.
	Direction OrderDirection `json:"direction"`
}

// The connection type for User.
type StargazerConnection struct {
	// A list of edges.
	Edges []*StargazerEdge `json:"edges"`
	// A list of nodes.
	Nodes []*User `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// Represents a user that's starred a repository.
type StargazerEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	Node   *User  `json:"node"`
	// Identifies when the item was starred.
	StarredAt string `json:"starredAt"`
}

// The connection type for Repository.
type StarredRepositoryConnection struct {
	// A list of edges.
	Edges []*StarredRepositoryEdge `json:"edges"`
	// Is the list of stars for this user truncated? This is true for users that have many stars.
	IsOverLimit bool `json:"isOverLimit"`
	// A list of nodes.
	Nodes []*Repository `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// Represents a starred repository.
type StarredRepositoryEdge struct {
	// A cursor for use in pagination.
	Cursor string      `json:"cursor"`
	Node   *Repository `json:"node"`
	// Identifies when the item was starred.
	StarredAt string `json:"starredAt"`
}

// Represents a commit status.
type Status struct {
	// A list of status contexts and check runs for this commit.
	CombinedContexts *StatusCheckRollupContextConnection `json:"combinedContexts"`
	// The commit this status is attached to.
	Commit *Commit `json:"commit"`
	// Looks up an individual status context by context name.
	Context *StatusContext `json:"context"`
	// The individual status contexts for this commit.
	Contexts []*StatusContext `json:"contexts"`
	ID       string           `json:"id"`
	// The combined commit status.
	State StatusState `json:"state"`
}

func (Status) IsNode() {}

// Represents the rollup for both the check runs and status for a commit.
type StatusCheckRollup struct {
	// The commit the status and check runs are attached to.
	Commit *Commit `json:"commit"`
	// A list of status contexts and check runs for this commit.
	Contexts *StatusCheckRollupContextConnection `json:"contexts"`
	ID       string                              `json:"id"`
	// The combined status for the commit.
	State StatusState `json:"state"`
}

func (StatusCheckRollup) IsNode() {}

// The connection type for StatusCheckRollupContext.
type StatusCheckRollupContextConnection struct {
	// A list of edges.
	Edges []*StatusCheckRollupContextEdge `json:"edges"`
	// A list of nodes.
	Nodes []StatusCheckRollupContext `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type StatusCheckRollupContextEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node StatusCheckRollupContext `json:"node"`
}

// Represents an individual commit status context
type StatusContext struct {
	// The avatar of the OAuth application or the user that created the status
	AvatarURL *string `json:"avatarUrl"`
	// This commit this status context is attached to.
	Commit *Commit `json:"commit"`
	// The name of this status context.
	Context string `json:"context"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	// The actor who created this status context.
	Creator Actor `json:"creator"`
	// The description for this status context.
	Description *string `json:"description"`
	ID          string  `json:"id"`
	// The state of this status context.
	State StatusState `json:"state"`
	// The URL for this status context.
	TargetURL *string `json:"targetUrl"`
}

func (StatusContext) IsStatusCheckRollupContext() {}
func (StatusContext) IsNode()                     {}

// Autogenerated input type of SubmitPullRequestReview
type SubmitPullRequestReviewInput struct {
	// The Pull Request ID to submit any pending reviews.
	PullRequestID *string `json:"pullRequestId"`
	// The Pull Request Review ID to submit.
	PullRequestReviewID *string `json:"pullRequestReviewId"`
	// The event to send to the Pull Request Review.
	Event PullRequestReviewEvent `json:"event"`
	// The text field to set on the Pull Request Review.
	Body *string `json:"body"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of SubmitPullRequestReview
type SubmitPullRequestReviewPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The submitted pull request review.
	PullRequestReview *PullRequestReview `json:"pullRequestReview"`
}

// A pointer to a repository at a specific revision embedded inside another repository.
type Submodule struct {
	// The branch of the upstream submodule for tracking updates
	Branch *string `json:"branch"`
	// The git URL of the submodule repository
	GitURL string `json:"gitUrl"`
	// The name of the submodule in .gitmodules
	Name string `json:"name"`
	// The path in the superproject that this submodule is located in
	Path string `json:"path"`
	// The commit revision of the subproject repository being tracked by the submodule
	SubprojectCommitOid *string `json:"subprojectCommitOid"`
}

// The connection type for Submodule.
type SubmoduleConnection struct {
	// A list of edges.
	Edges []*SubmoduleEdge `json:"edges"`
	// A list of nodes.
	Nodes []*Submodule `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type SubmoduleEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *Submodule `json:"node"`
}

// Represents a 'subscribed' event on a given `Subscribable`.
type SubscribedEvent struct {
	// Identifies the actor who performed the event.
	Actor Actor `json:"actor"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// Object referenced by event.
	Subscribable Subscribable `json:"subscribable"`
}

func (SubscribedEvent) IsPullRequestTimelineItems() {}
func (SubscribedEvent) IsIssueTimelineItems()       {}
func (SubscribedEvent) IsPullRequestTimelineItem()  {}
func (SubscribedEvent) IsNode()                     {}
func (SubscribedEvent) IsIssueTimelineItem()        {}

// A suggestion to review a pull request based on a user's commit history and review comments.
type SuggestedReviewer struct {
	// Is this suggestion based on past commits?
	IsAuthor bool `json:"isAuthor"`
	// Is this suggestion based on past review comments?
	IsCommenter bool `json:"isCommenter"`
	// Identifies the user suggested to review the pull request.
	Reviewer *User `json:"reviewer"`
}

// Represents a Git tag.
type Tag struct {
	// An abbreviated version of the Git object ID
	AbbreviatedOid string `json:"abbreviatedOid"`
	// The HTTP path for this Git object
	CommitResourcePath string `json:"commitResourcePath"`
	// The HTTP URL for this Git object
	CommitURL string `json:"commitUrl"`
	ID        string `json:"id"`
	// The Git tag message.
	Message *string `json:"message"`
	// The Git tag name.
	Name string `json:"name"`
	// The Git object ID
	Oid string `json:"oid"`
	// The Repository the Git object belongs to
	Repository *Repository `json:"repository"`
	// Details about the tag author.
	Tagger *GitActor `json:"tagger"`
	// The Git object the tag points to.
	Target GitObject `json:"target"`
}

func (Tag) IsNode()      {}
func (Tag) IsGitObject() {}

// A team of users in an organization.
type Team struct {
	// A list of teams that are ancestors of this team.
	Ancestors *TeamConnection `json:"ancestors"`
	// A URL pointing to the team's avatar.
	AvatarURL *string `json:"avatarUrl"`
	// List of child teams belonging to this team
	ChildTeams *TeamConnection `json:"childTeams"`
	// The slug corresponding to the organization and team.
	CombinedSlug string `json:"combinedSlug"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	// Identifies the primary key from the database.
	DatabaseID *int `json:"databaseId"`
	// The description of the team.
	Description *string `json:"description"`
	// Find a team discussion by its number.
	Discussion *TeamDiscussion `json:"discussion"`
	// A list of team discussions.
	Discussions *TeamDiscussionConnection `json:"discussions"`
	// The HTTP path for team discussions
	DiscussionsResourcePath string `json:"discussionsResourcePath"`
	// The HTTP URL for team discussions
	DiscussionsURL string `json:"discussionsUrl"`
	// The HTTP path for editing this team
	EditTeamResourcePath string `json:"editTeamResourcePath"`
	// The HTTP URL for editing this team
	EditTeamURL string `json:"editTeamUrl"`
	ID          string `json:"id"`
	// A list of pending invitations for users to this team
	Invitations *OrganizationInvitationConnection `json:"invitations"`
	// Get the status messages members of this entity have set that are either public or visible only to the organization.
	MemberStatuses *UserStatusConnection `json:"memberStatuses"`
	// A list of users who are members of this team.
	Members *TeamMemberConnection `json:"members"`
	// The HTTP path for the team' members
	MembersResourcePath string `json:"membersResourcePath"`
	// The HTTP URL for the team' members
	MembersURL string `json:"membersUrl"`
	// The name of the team.
	Name string `json:"name"`
	// The HTTP path creating a new team
	NewTeamResourcePath string `json:"newTeamResourcePath"`
	// The HTTP URL creating a new team
	NewTeamURL string `json:"newTeamUrl"`
	// The organization that owns this team.
	Organization *Organization `json:"organization"`
	// The parent team of the team.
	ParentTeam *Team `json:"parentTeam"`
	// The level of privacy the team has.
	Privacy TeamPrivacy `json:"privacy"`
	// A list of repositories this team has access to.
	Repositories *TeamRepositoryConnection `json:"repositories"`
	// The HTTP path for this team's repositories
	RepositoriesResourcePath string `json:"repositoriesResourcePath"`
	// The HTTP URL for this team's repositories
	RepositoriesURL string `json:"repositoriesUrl"`
	// The HTTP path for this team
	ResourcePath string `json:"resourcePath"`
	// The slug corresponding to the team.
	Slug string `json:"slug"`
	// The HTTP path for this team's teams
	TeamsResourcePath string `json:"teamsResourcePath"`
	// The HTTP URL for this team's teams
	TeamsURL string `json:"teamsUrl"`
	// Identifies the date and time when the object was last updated.
	UpdatedAt string `json:"updatedAt"`
	// The HTTP URL for this team
	URL string `json:"url"`
	// Team is adminable by the viewer.
	ViewerCanAdminister bool `json:"viewerCanAdminister"`
	// Check if the viewer is able to change their subscription status for the repository.
	ViewerCanSubscribe bool `json:"viewerCanSubscribe"`
	// Identifies if the viewer is watching, not watching, or ignoring the subscribable entity.
	ViewerSubscription *SubscriptionState `json:"viewerSubscription"`
}

func (Team) IsPermissionGranter()             {}
func (Team) IsRequestedReviewer()             {}
func (Team) IsReviewDismissalAllowanceActor() {}
func (Team) IsPushAllowanceActor()            {}
func (Team) IsNode()                          {}
func (Team) IsSubscribable()                  {}
func (Team) IsMemberStatusable()              {}

// Audit log entry for a team.add_member event.
type TeamAddMemberAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// Whether the team was mapped to an LDAP Group.
	IsLdapMapped *bool `json:"isLdapMapped"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The team associated with the action
	Team *Team `json:"team"`
	// The name of the team
	TeamName *string `json:"teamName"`
	// The HTTP path for this team
	TeamResourcePath *string `json:"teamResourcePath"`
	// The HTTP URL for this team
	TeamURL *string `json:"teamUrl"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
}

func (TeamAddMemberAuditEntry) IsNode()                       {}
func (TeamAddMemberAuditEntry) IsAuditEntry()                 {}
func (TeamAddMemberAuditEntry) IsOrganizationAuditEntryData() {}
func (TeamAddMemberAuditEntry) IsTeamAuditEntryData()         {}
func (TeamAddMemberAuditEntry) IsOrganizationAuditEntry()     {}

// Audit log entry for a team.add_repository event.
type TeamAddRepositoryAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// Whether the team was mapped to an LDAP Group.
	IsLdapMapped *bool `json:"isLdapMapped"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The repository associated with the action
	Repository *Repository `json:"repository"`
	// The name of the repository
	RepositoryName *string `json:"repositoryName"`
	// The HTTP path for the repository
	RepositoryResourcePath *string `json:"repositoryResourcePath"`
	// The HTTP URL for the repository
	RepositoryURL *string `json:"repositoryUrl"`
	// The team associated with the action
	Team *Team `json:"team"`
	// The name of the team
	TeamName *string `json:"teamName"`
	// The HTTP path for this team
	TeamResourcePath *string `json:"teamResourcePath"`
	// The HTTP URL for this team
	TeamURL *string `json:"teamUrl"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
}

func (TeamAddRepositoryAuditEntry) IsNode()                       {}
func (TeamAddRepositoryAuditEntry) IsAuditEntry()                 {}
func (TeamAddRepositoryAuditEntry) IsOrganizationAuditEntryData() {}
func (TeamAddRepositoryAuditEntry) IsRepositoryAuditEntryData()   {}
func (TeamAddRepositoryAuditEntry) IsTeamAuditEntryData()         {}
func (TeamAddRepositoryAuditEntry) IsOrganizationAuditEntry()     {}

// Audit log entry for a team.change_parent_team event.
type TeamChangeParentTeamAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// Whether the team was mapped to an LDAP Group.
	IsLdapMapped *bool `json:"isLdapMapped"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The new parent team.
	ParentTeam *Team `json:"parentTeam"`
	// The name of the new parent team
	ParentTeamName *string `json:"parentTeamName"`
	// The name of the former parent team
	ParentTeamNameWas *string `json:"parentTeamNameWas"`
	// The HTTP path for the parent team
	ParentTeamResourcePath *string `json:"parentTeamResourcePath"`
	// The HTTP URL for the parent team
	ParentTeamURL *string `json:"parentTeamUrl"`
	// The former parent team.
	ParentTeamWas *Team `json:"parentTeamWas"`
	// The HTTP path for the previous parent team
	ParentTeamWasResourcePath *string `json:"parentTeamWasResourcePath"`
	// The HTTP URL for the previous parent team
	ParentTeamWasURL *string `json:"parentTeamWasUrl"`
	// The team associated with the action
	Team *Team `json:"team"`
	// The name of the team
	TeamName *string `json:"teamName"`
	// The HTTP path for this team
	TeamResourcePath *string `json:"teamResourcePath"`
	// The HTTP URL for this team
	TeamURL *string `json:"teamUrl"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
}

func (TeamChangeParentTeamAuditEntry) IsNode()                       {}
func (TeamChangeParentTeamAuditEntry) IsAuditEntry()                 {}
func (TeamChangeParentTeamAuditEntry) IsOrganizationAuditEntryData() {}
func (TeamChangeParentTeamAuditEntry) IsTeamAuditEntryData()         {}
func (TeamChangeParentTeamAuditEntry) IsOrganizationAuditEntry()     {}

// The connection type for Team.
type TeamConnection struct {
	// A list of edges.
	Edges []*TeamEdge `json:"edges"`
	// A list of nodes.
	Nodes []*Team `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// A team discussion.
type TeamDiscussion struct {
	// The actor who authored the comment.
	Author Actor `json:"author"`
	// Author's association with the discussion's team.
	AuthorAssociation CommentAuthorAssociation `json:"authorAssociation"`
	// The body as Markdown.
	Body string `json:"body"`
	// The body rendered to HTML.
	BodyHTML string `json:"bodyHTML"`
	// The body rendered to text.
	BodyText string `json:"bodyText"`
	// Identifies the discussion body hash.
	BodyVersion string `json:"bodyVersion"`
	// A list of comments on this discussion.
	Comments *TeamDiscussionCommentConnection `json:"comments"`
	// The HTTP path for discussion comments
	CommentsResourcePath string `json:"commentsResourcePath"`
	// The HTTP URL for discussion comments
	CommentsURL string `json:"commentsUrl"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	// Check if this comment was created via an email reply.
	CreatedViaEmail bool `json:"createdViaEmail"`
	// Identifies the primary key from the database.
	DatabaseID *int `json:"databaseId"`
	// The actor who edited the comment.
	Editor Actor  `json:"editor"`
	ID     string `json:"id"`
	// Check if this comment was edited and includes an edit with the creation data
	IncludesCreatedEdit bool `json:"includesCreatedEdit"`
	// Whether or not the discussion is pinned.
	IsPinned bool `json:"isPinned"`
	// Whether or not the discussion is only visible to team members and org admins.
	IsPrivate bool `json:"isPrivate"`
	// The moment the editor made the last edit
	LastEditedAt *string `json:"lastEditedAt"`
	// Identifies the discussion within its team.
	Number int `json:"number"`
	// Identifies when the comment was published at.
	PublishedAt *string `json:"publishedAt"`
	// A list of reactions grouped by content left on the subject.
	ReactionGroups []*ReactionGroup `json:"reactionGroups"`
	// A list of Reactions left on the Issue.
	Reactions *ReactionConnection `json:"reactions"`
	// The HTTP path for this discussion
	ResourcePath string `json:"resourcePath"`
	// The team that defines the context of this discussion.
	Team *Team `json:"team"`
	// The title of the discussion
	Title string `json:"title"`
	// Identifies the date and time when the object was last updated.
	UpdatedAt string `json:"updatedAt"`
	// The HTTP URL for this discussion
	URL string `json:"url"`
	// A list of edits to this content.
	UserContentEdits *UserContentEditConnection `json:"userContentEdits"`
	// Check if the current viewer can delete this object.
	ViewerCanDelete bool `json:"viewerCanDelete"`
	// Whether or not the current viewer can pin this discussion.
	ViewerCanPin bool `json:"viewerCanPin"`
	// Can user react to this subject
	ViewerCanReact bool `json:"viewerCanReact"`
	// Check if the viewer is able to change their subscription status for the repository.
	ViewerCanSubscribe bool `json:"viewerCanSubscribe"`
	// Check if the current viewer can update this object.
	ViewerCanUpdate bool `json:"viewerCanUpdate"`
	// Reasons why the current viewer can not update this comment.
	ViewerCannotUpdateReasons []CommentCannotUpdateReason `json:"viewerCannotUpdateReasons"`
	// Did the viewer author this comment.
	ViewerDidAuthor bool `json:"viewerDidAuthor"`
	// Identifies if the viewer is watching, not watching, or ignoring the subscribable entity.
	ViewerSubscription *SubscriptionState `json:"viewerSubscription"`
}

func (TeamDiscussion) IsNode()                     {}
func (TeamDiscussion) IsComment()                  {}
func (TeamDiscussion) IsDeletable()                {}
func (TeamDiscussion) IsReactable()                {}
func (TeamDiscussion) IsSubscribable()             {}
func (TeamDiscussion) IsUniformResourceLocatable() {}
func (TeamDiscussion) IsUpdatable()                {}
func (TeamDiscussion) IsUpdatableComment()         {}

// A comment on a team discussion.
type TeamDiscussionComment struct {
	// The actor who authored the comment.
	Author Actor `json:"author"`
	// Author's association with the comment's team.
	AuthorAssociation CommentAuthorAssociation `json:"authorAssociation"`
	// The body as Markdown.
	Body string `json:"body"`
	// The body rendered to HTML.
	BodyHTML string `json:"bodyHTML"`
	// The body rendered to text.
	BodyText string `json:"bodyText"`
	// The current version of the body content.
	BodyVersion string `json:"bodyVersion"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	// Check if this comment was created via an email reply.
	CreatedViaEmail bool `json:"createdViaEmail"`
	// Identifies the primary key from the database.
	DatabaseID *int `json:"databaseId"`
	// The discussion this comment is about.
	Discussion *TeamDiscussion `json:"discussion"`
	// The actor who edited the comment.
	Editor Actor  `json:"editor"`
	ID     string `json:"id"`
	// Check if this comment was edited and includes an edit with the creation data
	IncludesCreatedEdit bool `json:"includesCreatedEdit"`
	// The moment the editor made the last edit
	LastEditedAt *string `json:"lastEditedAt"`
	// Identifies the comment number.
	Number int `json:"number"`
	// Identifies when the comment was published at.
	PublishedAt *string `json:"publishedAt"`
	// A list of reactions grouped by content left on the subject.
	ReactionGroups []*ReactionGroup `json:"reactionGroups"`
	// A list of Reactions left on the Issue.
	Reactions *ReactionConnection `json:"reactions"`
	// The HTTP path for this comment
	ResourcePath string `json:"resourcePath"`
	// Identifies the date and time when the object was last updated.
	UpdatedAt string `json:"updatedAt"`
	// The HTTP URL for this comment
	URL string `json:"url"`
	// A list of edits to this content.
	UserContentEdits *UserContentEditConnection `json:"userContentEdits"`
	// Check if the current viewer can delete this object.
	ViewerCanDelete bool `json:"viewerCanDelete"`
	// Can user react to this subject
	ViewerCanReact bool `json:"viewerCanReact"`
	// Check if the current viewer can update this object.
	ViewerCanUpdate bool `json:"viewerCanUpdate"`
	// Reasons why the current viewer can not update this comment.
	ViewerCannotUpdateReasons []CommentCannotUpdateReason `json:"viewerCannotUpdateReasons"`
	// Did the viewer author this comment.
	ViewerDidAuthor bool `json:"viewerDidAuthor"`
}

func (TeamDiscussionComment) IsNode()                     {}
func (TeamDiscussionComment) IsComment()                  {}
func (TeamDiscussionComment) IsDeletable()                {}
func (TeamDiscussionComment) IsReactable()                {}
func (TeamDiscussionComment) IsUniformResourceLocatable() {}
func (TeamDiscussionComment) IsUpdatable()                {}
func (TeamDiscussionComment) IsUpdatableComment()         {}

// The connection type for TeamDiscussionComment.
type TeamDiscussionCommentConnection struct {
	// A list of edges.
	Edges []*TeamDiscussionCommentEdge `json:"edges"`
	// A list of nodes.
	Nodes []*TeamDiscussionComment `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type TeamDiscussionCommentEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *TeamDiscussionComment `json:"node"`
}

// Ways in which team discussion comment connections can be ordered.
type TeamDiscussionCommentOrder struct {
	// The field by which to order nodes.
	Field TeamDiscussionCommentOrderField `json:"field"`
	// The direction in which to order nodes.
	Direction OrderDirection `json:"direction"`
}

// The connection type for TeamDiscussion.
type TeamDiscussionConnection struct {
	// A list of edges.
	Edges []*TeamDiscussionEdge `json:"edges"`
	// A list of nodes.
	Nodes []*TeamDiscussion `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type TeamDiscussionEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *TeamDiscussion `json:"node"`
}

// Ways in which team discussion connections can be ordered.
type TeamDiscussionOrder struct {
	// The field by which to order nodes.
	Field TeamDiscussionOrderField `json:"field"`
	// The direction in which to order nodes.
	Direction OrderDirection `json:"direction"`
}

// An edge in a connection.
type TeamEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *Team `json:"node"`
}

// The connection type for User.
type TeamMemberConnection struct {
	// A list of edges.
	Edges []*TeamMemberEdge `json:"edges"`
	// A list of nodes.
	Nodes []*User `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// Represents a user who is a member of a team.
type TeamMemberEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The HTTP path to the organization's member access page.
	MemberAccessResourcePath string `json:"memberAccessResourcePath"`
	// The HTTP URL to the organization's member access page.
	MemberAccessURL string `json:"memberAccessUrl"`
	Node            *User  `json:"node"`
	// The role the member has on the team.
	Role TeamMemberRole `json:"role"`
}

// Ordering options for team member connections
type TeamMemberOrder struct {
	// The field to order team members by.
	Field TeamMemberOrderField `json:"field"`
	// The ordering direction.
	Direction OrderDirection `json:"direction"`
}

// Ways in which team connections can be ordered.
type TeamOrder struct {
	// The field in which to order nodes by.
	Field TeamOrderField `json:"field"`
	// The direction in which to order nodes.
	Direction OrderDirection `json:"direction"`
}

// Audit log entry for a team.remove_member event.
type TeamRemoveMemberAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// Whether the team was mapped to an LDAP Group.
	IsLdapMapped *bool `json:"isLdapMapped"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The team associated with the action
	Team *Team `json:"team"`
	// The name of the team
	TeamName *string `json:"teamName"`
	// The HTTP path for this team
	TeamResourcePath *string `json:"teamResourcePath"`
	// The HTTP URL for this team
	TeamURL *string `json:"teamUrl"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
}

func (TeamRemoveMemberAuditEntry) IsOrganizationAuditEntry()     {}
func (TeamRemoveMemberAuditEntry) IsNode()                       {}
func (TeamRemoveMemberAuditEntry) IsAuditEntry()                 {}
func (TeamRemoveMemberAuditEntry) IsOrganizationAuditEntryData() {}
func (TeamRemoveMemberAuditEntry) IsTeamAuditEntryData()         {}

// Audit log entry for a team.remove_repository event.
type TeamRemoveRepositoryAuditEntry struct {
	// The action name
	Action string `json:"action"`
	// The user who initiated the action
	Actor AuditEntryActor `json:"actor"`
	// The IP address of the actor
	ActorIP *string `json:"actorIp"`
	// A readable representation of the actor's location
	ActorLocation *ActorLocation `json:"actorLocation"`
	// The username of the user who initiated the action
	ActorLogin *string `json:"actorLogin"`
	// The HTTP path for the actor.
	ActorResourcePath *string `json:"actorResourcePath"`
	// The HTTP URL for the actor.
	ActorURL *string `json:"actorUrl"`
	// The time the action was initiated
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// Whether the team was mapped to an LDAP Group.
	IsLdapMapped *bool `json:"isLdapMapped"`
	// The corresponding operation type for the action
	OperationType *OperationType `json:"operationType"`
	// The Organization associated with the Audit Entry.
	Organization *Organization `json:"organization"`
	// The name of the Organization.
	OrganizationName *string `json:"organizationName"`
	// The HTTP path for the organization
	OrganizationResourcePath *string `json:"organizationResourcePath"`
	// The HTTP URL for the organization
	OrganizationURL *string `json:"organizationUrl"`
	// The repository associated with the action
	Repository *Repository `json:"repository"`
	// The name of the repository
	RepositoryName *string `json:"repositoryName"`
	// The HTTP path for the repository
	RepositoryResourcePath *string `json:"repositoryResourcePath"`
	// The HTTP URL for the repository
	RepositoryURL *string `json:"repositoryUrl"`
	// The team associated with the action
	Team *Team `json:"team"`
	// The name of the team
	TeamName *string `json:"teamName"`
	// The HTTP path for this team
	TeamResourcePath *string `json:"teamResourcePath"`
	// The HTTP URL for this team
	TeamURL *string `json:"teamUrl"`
	// The user affected by the action
	User *User `json:"user"`
	// For actions involving two users, the actor is the initiator and the user is the affected user.
	UserLogin *string `json:"userLogin"`
	// The HTTP path for the user.
	UserResourcePath *string `json:"userResourcePath"`
	// The HTTP URL for the user.
	UserURL *string `json:"userUrl"`
}

func (TeamRemoveRepositoryAuditEntry) IsNode()                       {}
func (TeamRemoveRepositoryAuditEntry) IsAuditEntry()                 {}
func (TeamRemoveRepositoryAuditEntry) IsOrganizationAuditEntryData() {}
func (TeamRemoveRepositoryAuditEntry) IsRepositoryAuditEntryData()   {}
func (TeamRemoveRepositoryAuditEntry) IsTeamAuditEntryData()         {}
func (TeamRemoveRepositoryAuditEntry) IsOrganizationAuditEntry()     {}

// The connection type for Repository.
type TeamRepositoryConnection struct {
	// A list of edges.
	Edges []*TeamRepositoryEdge `json:"edges"`
	// A list of nodes.
	Nodes []*Repository `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// Represents a team repository.
type TeamRepositoryEdge struct {
	// A cursor for use in pagination.
	Cursor string      `json:"cursor"`
	Node   *Repository `json:"node"`
	// The permission level the team has on the repository
	//
	// **Upcoming Change on 2020-10-01 UTC**
	// **Description:** Type for `permission` will change from `RepositoryPermission!` to `String`.
	// **Reason:** This field may return additional values
	//
	Permission RepositoryPermission `json:"permission"`
}

// Ordering options for team repository connections
type TeamRepositoryOrder struct {
	// The field to order repositories by.
	Field TeamRepositoryOrderField `json:"field"`
	// The ordering direction.
	Direction OrderDirection `json:"direction"`
}

// A text match within a search result.
type TextMatch struct {
	// The specific text fragment within the property matched on.
	Fragment string `json:"fragment"`
	// Highlights within the matched fragment.
	Highlights []*TextMatchHighlight `json:"highlights"`
	// The property matched on.
	Property string `json:"property"`
}

// Represents a single highlight in a search result match.
type TextMatchHighlight struct {
	// The indice in the fragment where the matched text begins.
	BeginIndice int `json:"beginIndice"`
	// The indice in the fragment where the matched text ends.
	EndIndice int `json:"endIndice"`
	// The text matched.
	Text string `json:"text"`
}

// A topic aggregates entities that are related to a subject.
type Topic struct {
	ID string `json:"id"`
	// The topic's name.
	Name string `json:"name"`
	// A list of related topics, including aliases of this topic, sorted with the most relevant
	// first. Returns up to 10 Topics.
	//
	RelatedTopics []*Topic `json:"relatedTopics"`
	// Returns a count of how many stargazers there are on this object
	//
	StargazerCount int `json:"stargazerCount"`
	// A list of users who have starred this starrable.
	Stargazers *StargazerConnection `json:"stargazers"`
	// Returns a boolean indicating whether the viewing user has starred this starrable.
	ViewerHasStarred bool `json:"viewerHasStarred"`
}

func (Topic) IsNode()      {}
func (Topic) IsStarrable() {}

// Autogenerated input type of TransferIssue
type TransferIssueInput struct {
	// The Node ID of the issue to be transferred
	IssueID string `json:"issueId"`
	// The Node ID of the repository the issue should be transferred to
	RepositoryID string `json:"repositoryId"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of TransferIssue
type TransferIssuePayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The issue that was transferred
	Issue *Issue `json:"issue"`
}

// Represents a 'transferred' event on a given issue or pull request.
type TransferredEvent struct {
	// Identifies the actor who performed the event.
	Actor Actor `json:"actor"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	// The repository this came from
	FromRepository *Repository `json:"fromRepository"`
	ID             string      `json:"id"`
	// Identifies the issue associated with the event.
	Issue *Issue `json:"issue"`
}

func (TransferredEvent) IsPullRequestTimelineItems() {}
func (TransferredEvent) IsIssueTimelineItems()       {}
func (TransferredEvent) IsIssueTimelineItem()        {}
func (TransferredEvent) IsNode()                     {}

// Represents a Git tree.
type Tree struct {
	// An abbreviated version of the Git object ID
	AbbreviatedOid string `json:"abbreviatedOid"`
	// The HTTP path for this Git object
	CommitResourcePath string `json:"commitResourcePath"`
	// The HTTP URL for this Git object
	CommitURL string `json:"commitUrl"`
	// A list of tree entries.
	Entries []*TreeEntry `json:"entries"`
	ID      string       `json:"id"`
	// The Git object ID
	Oid string `json:"oid"`
	// The Repository the Git object belongs to
	Repository *Repository `json:"repository"`
}

func (Tree) IsNode()      {}
func (Tree) IsGitObject() {}

// Represents a Git tree entry.
type TreeEntry struct {
	// The extension of the file
	Extension *string `json:"extension"`
	// Whether or not this tree entry is generated
	IsGenerated bool `json:"isGenerated"`
	// Entry file mode.
	Mode int `json:"mode"`
	// Entry file name.
	Name string `json:"name"`
	// Entry file object.
	Object GitObject `json:"object"`
	// Entry file Git object ID.
	Oid string `json:"oid"`
	// The full path of the file.
	Path *string `json:"path"`
	// The Repository the tree entry belongs to
	Repository *Repository `json:"repository"`
	// If the TreeEntry is for a directory occupied by a submodule project, this returns the corresponding submodule
	Submodule *Submodule `json:"submodule"`
	// Entry file type.
	Type string `json:"type"`
}

// Autogenerated input type of UnarchiveRepository
type UnarchiveRepositoryInput struct {
	// The ID of the repository to unarchive.
	RepositoryID string `json:"repositoryId"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of UnarchiveRepository
type UnarchiveRepositoryPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The repository that was unarchived.
	Repository *Repository `json:"repository"`
}

// Represents an 'unassigned' event on any assignable object.
type UnassignedEvent struct {
	// Identifies the actor who performed the event.
	Actor Actor `json:"actor"`
	// Identifies the assignable associated with the event.
	Assignable Assignable `json:"assignable"`
	// Identifies the user or mannequin that was unassigned.
	Assignee Assignee `json:"assignee"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// Identifies the subject (user) who was unassigned.
	User *User `json:"user"`
}

func (UnassignedEvent) IsNode()                     {}
func (UnassignedEvent) IsPullRequestTimelineItems() {}
func (UnassignedEvent) IsIssueTimelineItems()       {}
func (UnassignedEvent) IsPullRequestTimelineItem()  {}
func (UnassignedEvent) IsIssueTimelineItem()        {}

// Autogenerated input type of UnfollowUser
type UnfollowUserInput struct {
	// ID of the user to unfollow.
	UserID string `json:"userId"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of UnfollowUser
type UnfollowUserPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The user that was unfollowed.
	User *User `json:"user"`
}

// Represents an unknown signature on a Commit or Tag.
type UnknownSignature struct {
	// Email used to sign this object.
	Email string `json:"email"`
	// True if the signature is valid and verified by GitHub.
	IsValid bool `json:"isValid"`
	// Payload for GPG signing object. Raw ODB object without the signature header.
	Payload string `json:"payload"`
	// ASCII-armored signature header from object.
	Signature string `json:"signature"`
	// GitHub user corresponding to the email signing this commit.
	Signer *User `json:"signer"`
	// The state of this signature. `VALID` if signature is valid and verified by GitHub, otherwise represents reason why signature is considered invalid.
	State GitSignatureState `json:"state"`
	// True if the signature was made with GitHub's signing key.
	WasSignedByGitHub bool `json:"wasSignedByGitHub"`
}

func (UnknownSignature) IsGitSignature() {}

// Represents an 'unlabeled' event on a given issue or pull request.
type UnlabeledEvent struct {
	// Identifies the actor who performed the event.
	Actor Actor `json:"actor"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// Identifies the label associated with the 'unlabeled' event.
	Label *Label `json:"label"`
	// Identifies the `Labelable` associated with the event.
	Labelable Labelable `json:"labelable"`
}

func (UnlabeledEvent) IsPullRequestTimelineItems() {}
func (UnlabeledEvent) IsIssueTimelineItems()       {}
func (UnlabeledEvent) IsPullRequestTimelineItem()  {}
func (UnlabeledEvent) IsIssueTimelineItem()        {}
func (UnlabeledEvent) IsNode()                     {}

// Autogenerated input type of UnlinkRepositoryFromProject
type UnlinkRepositoryFromProjectInput struct {
	// The ID of the Project linked to the Repository.
	ProjectID string `json:"projectId"`
	// The ID of the Repository linked to the Project.
	RepositoryID string `json:"repositoryId"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of UnlinkRepositoryFromProject
type UnlinkRepositoryFromProjectPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The linked Project.
	Project *Project `json:"project"`
	// The linked Repository.
	Repository *Repository `json:"repository"`
}

// Autogenerated input type of UnlockLockable
type UnlockLockableInput struct {
	// ID of the issue or pull request to be unlocked.
	LockableID string `json:"lockableId"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of UnlockLockable
type UnlockLockablePayload struct {
	// Identifies the actor who performed the event.
	Actor Actor `json:"actor"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The item that was unlocked.
	UnlockedRecord Lockable `json:"unlockedRecord"`
}

// Represents an 'unlocked' event on a given issue or pull request.
type UnlockedEvent struct {
	// Identifies the actor who performed the event.
	Actor Actor `json:"actor"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// Object that was unlocked.
	Lockable Lockable `json:"lockable"`
}

func (UnlockedEvent) IsPullRequestTimelineItems() {}
func (UnlockedEvent) IsIssueTimelineItems()       {}
func (UnlockedEvent) IsPullRequestTimelineItem()  {}
func (UnlockedEvent) IsIssueTimelineItem()        {}
func (UnlockedEvent) IsNode()                     {}

// Autogenerated input type of UnmarkFileAsViewed
type UnmarkFileAsViewedInput struct {
	// The Node ID of the pull request.
	PullRequestID string `json:"pullRequestId"`
	// The path of the file to mark as unviewed
	Path string `json:"path"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of UnmarkFileAsViewed
type UnmarkFileAsViewedPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The updated pull request.
	PullRequest *PullRequest `json:"pullRequest"`
}

// Autogenerated input type of UnmarkIssueAsDuplicate
type UnmarkIssueAsDuplicateInput struct {
	// ID of the issue or pull request currently marked as a duplicate.
	DuplicateID string `json:"duplicateId"`
	// ID of the issue or pull request currently considered canonical/authoritative/original.
	CanonicalID string `json:"canonicalId"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of UnmarkIssueAsDuplicate
type UnmarkIssueAsDuplicatePayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The issue or pull request that was marked as a duplicate.
	Duplicate IssueOrPullRequest `json:"duplicate"`
}

// Represents an 'unmarked_as_duplicate' event on a given issue or pull request.
type UnmarkedAsDuplicateEvent struct {
	// Identifies the actor who performed the event.
	Actor Actor `json:"actor"`
	// The authoritative issue or pull request which has been duplicated by another.
	Canonical IssueOrPullRequest `json:"canonical"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	// The issue or pull request which has been marked as a duplicate of another.
	Duplicate IssueOrPullRequest `json:"duplicate"`
	ID        string             `json:"id"`
	// Canonical and duplicate belong to different repositories.
	IsCrossRepository bool `json:"isCrossRepository"`
}

func (UnmarkedAsDuplicateEvent) IsPullRequestTimelineItems() {}
func (UnmarkedAsDuplicateEvent) IsIssueTimelineItems()       {}
func (UnmarkedAsDuplicateEvent) IsNode()                     {}

// Autogenerated input type of UnminimizeComment
type UnminimizeCommentInput struct {
	// The Node ID of the subject to modify.
	SubjectID string `json:"subjectId"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of UnminimizeComment
type UnminimizeCommentPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The comment that was unminimized.
	UnminimizedComment Minimizable `json:"unminimizedComment"`
}

// Represents an 'unpinned' event on a given issue or pull request.
type UnpinnedEvent struct {
	// Identifies the actor who performed the event.
	Actor Actor `json:"actor"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// Identifies the issue associated with the event.
	Issue *Issue `json:"issue"`
}

func (UnpinnedEvent) IsPullRequestTimelineItems() {}
func (UnpinnedEvent) IsIssueTimelineItems()       {}
func (UnpinnedEvent) IsNode()                     {}

// Autogenerated input type of UnresolveReviewThread
type UnresolveReviewThreadInput struct {
	// The ID of the thread to unresolve
	ThreadID string `json:"threadId"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of UnresolveReviewThread
type UnresolveReviewThreadPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The thread to resolve.
	Thread *PullRequestReviewThread `json:"thread"`
}

// Represents an 'unsubscribed' event on a given `Subscribable`.
type UnsubscribedEvent struct {
	// Identifies the actor who performed the event.
	Actor Actor `json:"actor"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// Object referenced by event.
	Subscribable Subscribable `json:"subscribable"`
}

func (UnsubscribedEvent) IsPullRequestTimelineItems() {}
func (UnsubscribedEvent) IsIssueTimelineItems()       {}
func (UnsubscribedEvent) IsPullRequestTimelineItem()  {}
func (UnsubscribedEvent) IsIssueTimelineItem()        {}
func (UnsubscribedEvent) IsNode()                     {}

// Autogenerated input type of UpdateBranchProtectionRule
type UpdateBranchProtectionRuleInput struct {
	// The global relay id of the branch protection rule to be updated.
	BranchProtectionRuleID string `json:"branchProtectionRuleId"`
	// The glob-like pattern used to determine matching branches.
	Pattern *string `json:"pattern"`
	// Are approving reviews required to update matching branches.
	RequiresApprovingReviews *bool `json:"requiresApprovingReviews"`
	// Number of approving reviews required to update matching branches.
	RequiredApprovingReviewCount *int `json:"requiredApprovingReviewCount"`
	// Are commits required to be signed.
	RequiresCommitSignatures *bool `json:"requiresCommitSignatures"`
	// Can admins overwrite branch protection.
	IsAdminEnforced *bool `json:"isAdminEnforced"`
	// Are status checks required to update matching branches.
	RequiresStatusChecks *bool `json:"requiresStatusChecks"`
	// Are branches required to be up to date before merging.
	RequiresStrictStatusChecks *bool `json:"requiresStrictStatusChecks"`
	// Are reviews from code owners required to update matching branches.
	RequiresCodeOwnerReviews *bool `json:"requiresCodeOwnerReviews"`
	// Will new commits pushed to matching branches dismiss pull request review approvals.
	DismissesStaleReviews *bool `json:"dismissesStaleReviews"`
	// Is dismissal of pull request reviews restricted.
	RestrictsReviewDismissals *bool `json:"restrictsReviewDismissals"`
	// A list of User or Team IDs allowed to dismiss reviews on pull requests targeting matching branches.
	ReviewDismissalActorIds []string `json:"reviewDismissalActorIds"`
	// Is pushing to matching branches restricted.
	RestrictsPushes *bool `json:"restrictsPushes"`
	// A list of User, Team or App IDs allowed to push to matching branches.
	PushActorIds []string `json:"pushActorIds"`
	// List of required status check contexts that must pass for commits to be accepted to matching branches.
	RequiredStatusCheckContexts []string `json:"requiredStatusCheckContexts"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of UpdateBranchProtectionRule
type UpdateBranchProtectionRulePayload struct {
	// The newly created BranchProtectionRule.
	BranchProtectionRule *BranchProtectionRule `json:"branchProtectionRule"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated input type of UpdateCheckRun
type UpdateCheckRunInput struct {
	// The node ID of the repository.
	RepositoryID string `json:"repositoryId"`
	// The node of the check.
	CheckRunID string `json:"checkRunId"`
	// The name of the check.
	Name *string `json:"name"`
	// The URL of the integrator's site that has the full details of the check.
	DetailsURL *string `json:"detailsUrl"`
	// A reference for the run on the integrator's system.
	ExternalID *string `json:"externalId"`
	// The current status.
	Status *RequestableCheckStatusState `json:"status"`
	// The time that the check run began.
	StartedAt *string `json:"startedAt"`
	// The final conclusion of the check.
	Conclusion *CheckConclusionState `json:"conclusion"`
	// The time that the check run finished.
	CompletedAt *string `json:"completedAt"`
	// Descriptive details about the run.
	Output *CheckRunOutput `json:"output"`
	// Possible further actions the integrator can perform, which a user may trigger.
	Actions []*CheckRunAction `json:"actions"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of UpdateCheckRun
type UpdateCheckRunPayload struct {
	// The updated check run.
	CheckRun *CheckRun `json:"checkRun"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated input type of UpdateCheckSuitePreferences
type UpdateCheckSuitePreferencesInput struct {
	// The Node ID of the repository.
	RepositoryID string `json:"repositoryId"`
	// The check suite preferences to modify.
	AutoTriggerPreferences []*CheckSuiteAutoTriggerPreference `json:"autoTriggerPreferences"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of UpdateCheckSuitePreferences
type UpdateCheckSuitePreferencesPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The updated repository.
	Repository *Repository `json:"repository"`
}

// Autogenerated input type of UpdateEnterpriseActionExecutionCapabilitySetting
type UpdateEnterpriseActionExecutionCapabilitySettingInput struct {
	// The ID of the enterprise on which to set the members can create repositories setting.
	EnterpriseID string `json:"enterpriseId"`
	// The value for the action execution capability setting on the enterprise.
	Capability ActionExecutionCapabilitySetting `json:"capability"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of UpdateEnterpriseActionExecutionCapabilitySetting
type UpdateEnterpriseActionExecutionCapabilitySettingPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The enterprise with the updated action execution capability setting.
	Enterprise *Enterprise `json:"enterprise"`
	// A message confirming the result of updating the action execution capability setting.
	Message *string `json:"message"`
}

// Autogenerated input type of UpdateEnterpriseAdministratorRole
type UpdateEnterpriseAdministratorRoleInput struct {
	// The ID of the Enterprise which the admin belongs to.
	EnterpriseID string `json:"enterpriseId"`
	// The login of a administrator whose role is being changed.
	Login string `json:"login"`
	// The new role for the Enterprise administrator.
	Role EnterpriseAdministratorRole `json:"role"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of UpdateEnterpriseAdministratorRole
type UpdateEnterpriseAdministratorRolePayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// A message confirming the result of changing the administrator's role.
	Message *string `json:"message"`
}

// Autogenerated input type of UpdateEnterpriseAllowPrivateRepositoryForkingSetting
type UpdateEnterpriseAllowPrivateRepositoryForkingSettingInput struct {
	// The ID of the enterprise on which to set the allow private repository forking setting.
	EnterpriseID string `json:"enterpriseId"`
	// The value for the allow private repository forking setting on the enterprise.
	SettingValue EnterpriseEnabledDisabledSettingValue `json:"settingValue"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of UpdateEnterpriseAllowPrivateRepositoryForkingSetting
type UpdateEnterpriseAllowPrivateRepositoryForkingSettingPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The enterprise with the updated allow private repository forking setting.
	Enterprise *Enterprise `json:"enterprise"`
	// A message confirming the result of updating the allow private repository forking setting.
	Message *string `json:"message"`
}

// Autogenerated input type of UpdateEnterpriseDefaultRepositoryPermissionSetting
type UpdateEnterpriseDefaultRepositoryPermissionSettingInput struct {
	// The ID of the enterprise on which to set the default repository permission setting.
	EnterpriseID string `json:"enterpriseId"`
	// The value for the default repository permission setting on the enterprise.
	SettingValue EnterpriseDefaultRepositoryPermissionSettingValue `json:"settingValue"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of UpdateEnterpriseDefaultRepositoryPermissionSetting
type UpdateEnterpriseDefaultRepositoryPermissionSettingPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The enterprise with the updated default repository permission setting.
	Enterprise *Enterprise `json:"enterprise"`
	// A message confirming the result of updating the default repository permission setting.
	Message *string `json:"message"`
}

// Autogenerated input type of UpdateEnterpriseMembersCanChangeRepositoryVisibilitySetting
type UpdateEnterpriseMembersCanChangeRepositoryVisibilitySettingInput struct {
	// The ID of the enterprise on which to set the members can change repository visibility setting.
	EnterpriseID string `json:"enterpriseId"`
	// The value for the members can change repository visibility setting on the enterprise.
	SettingValue EnterpriseEnabledDisabledSettingValue `json:"settingValue"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of UpdateEnterpriseMembersCanChangeRepositoryVisibilitySetting
type UpdateEnterpriseMembersCanChangeRepositoryVisibilitySettingPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The enterprise with the updated members can change repository visibility setting.
	Enterprise *Enterprise `json:"enterprise"`
	// A message confirming the result of updating the members can change repository visibility setting.
	Message *string `json:"message"`
}

// Autogenerated input type of UpdateEnterpriseMembersCanCreateRepositoriesSetting
type UpdateEnterpriseMembersCanCreateRepositoriesSettingInput struct {
	// The ID of the enterprise on which to set the members can create repositories setting.
	EnterpriseID string `json:"enterpriseId"`
	// Value for the members can create repositories setting on the enterprise. This or the granular public/private/internal allowed fields (but not both) must be provided.
	SettingValue *EnterpriseMembersCanCreateRepositoriesSettingValue `json:"settingValue"`
	// When false, allow member organizations to set their own repository creation member privileges.
	MembersCanCreateRepositoriesPolicyEnabled *bool `json:"membersCanCreateRepositoriesPolicyEnabled"`
	// Allow members to create public repositories. Defaults to current value.
	MembersCanCreatePublicRepositories *bool `json:"membersCanCreatePublicRepositories"`
	// Allow members to create private repositories. Defaults to current value.
	MembersCanCreatePrivateRepositories *bool `json:"membersCanCreatePrivateRepositories"`
	// Allow members to create internal repositories. Defaults to current value.
	MembersCanCreateInternalRepositories *bool `json:"membersCanCreateInternalRepositories"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of UpdateEnterpriseMembersCanCreateRepositoriesSetting
type UpdateEnterpriseMembersCanCreateRepositoriesSettingPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The enterprise with the updated members can create repositories setting.
	Enterprise *Enterprise `json:"enterprise"`
	// A message confirming the result of updating the members can create repositories setting.
	Message *string `json:"message"`
}

// Autogenerated input type of UpdateEnterpriseMembersCanDeleteIssuesSetting
type UpdateEnterpriseMembersCanDeleteIssuesSettingInput struct {
	// The ID of the enterprise on which to set the members can delete issues setting.
	EnterpriseID string `json:"enterpriseId"`
	// The value for the members can delete issues setting on the enterprise.
	SettingValue EnterpriseEnabledDisabledSettingValue `json:"settingValue"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of UpdateEnterpriseMembersCanDeleteIssuesSetting
type UpdateEnterpriseMembersCanDeleteIssuesSettingPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The enterprise with the updated members can delete issues setting.
	Enterprise *Enterprise `json:"enterprise"`
	// A message confirming the result of updating the members can delete issues setting.
	Message *string `json:"message"`
}

// Autogenerated input type of UpdateEnterpriseMembersCanDeleteRepositoriesSetting
type UpdateEnterpriseMembersCanDeleteRepositoriesSettingInput struct {
	// The ID of the enterprise on which to set the members can delete repositories setting.
	EnterpriseID string `json:"enterpriseId"`
	// The value for the members can delete repositories setting on the enterprise.
	SettingValue EnterpriseEnabledDisabledSettingValue `json:"settingValue"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of UpdateEnterpriseMembersCanDeleteRepositoriesSetting
type UpdateEnterpriseMembersCanDeleteRepositoriesSettingPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The enterprise with the updated members can delete repositories setting.
	Enterprise *Enterprise `json:"enterprise"`
	// A message confirming the result of updating the members can delete repositories setting.
	Message *string `json:"message"`
}

// Autogenerated input type of UpdateEnterpriseMembersCanInviteCollaboratorsSetting
type UpdateEnterpriseMembersCanInviteCollaboratorsSettingInput struct {
	// The ID of the enterprise on which to set the members can invite collaborators setting.
	EnterpriseID string `json:"enterpriseId"`
	// The value for the members can invite collaborators setting on the enterprise.
	SettingValue EnterpriseEnabledDisabledSettingValue `json:"settingValue"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of UpdateEnterpriseMembersCanInviteCollaboratorsSetting
type UpdateEnterpriseMembersCanInviteCollaboratorsSettingPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The enterprise with the updated members can invite collaborators setting.
	Enterprise *Enterprise `json:"enterprise"`
	// A message confirming the result of updating the members can invite collaborators setting.
	Message *string `json:"message"`
}

// Autogenerated input type of UpdateEnterpriseMembersCanMakePurchasesSetting
type UpdateEnterpriseMembersCanMakePurchasesSettingInput struct {
	// The ID of the enterprise on which to set the members can make purchases setting.
	EnterpriseID string `json:"enterpriseId"`
	// The value for the members can make purchases setting on the enterprise.
	SettingValue EnterpriseMembersCanMakePurchasesSettingValue `json:"settingValue"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of UpdateEnterpriseMembersCanMakePurchasesSetting
type UpdateEnterpriseMembersCanMakePurchasesSettingPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The enterprise with the updated members can make purchases setting.
	Enterprise *Enterprise `json:"enterprise"`
	// A message confirming the result of updating the members can make purchases setting.
	Message *string `json:"message"`
}

// Autogenerated input type of UpdateEnterpriseMembersCanUpdateProtectedBranchesSetting
type UpdateEnterpriseMembersCanUpdateProtectedBranchesSettingInput struct {
	// The ID of the enterprise on which to set the members can update protected branches setting.
	EnterpriseID string `json:"enterpriseId"`
	// The value for the members can update protected branches setting on the enterprise.
	SettingValue EnterpriseEnabledDisabledSettingValue `json:"settingValue"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of UpdateEnterpriseMembersCanUpdateProtectedBranchesSetting
type UpdateEnterpriseMembersCanUpdateProtectedBranchesSettingPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The enterprise with the updated members can update protected branches setting.
	Enterprise *Enterprise `json:"enterprise"`
	// A message confirming the result of updating the members can update protected branches setting.
	Message *string `json:"message"`
}

// Autogenerated input type of UpdateEnterpriseMembersCanViewDependencyInsightsSetting
type UpdateEnterpriseMembersCanViewDependencyInsightsSettingInput struct {
	// The ID of the enterprise on which to set the members can view dependency insights setting.
	EnterpriseID string `json:"enterpriseId"`
	// The value for the members can view dependency insights setting on the enterprise.
	SettingValue EnterpriseEnabledDisabledSettingValue `json:"settingValue"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of UpdateEnterpriseMembersCanViewDependencyInsightsSetting
type UpdateEnterpriseMembersCanViewDependencyInsightsSettingPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The enterprise with the updated members can view dependency insights setting.
	Enterprise *Enterprise `json:"enterprise"`
	// A message confirming the result of updating the members can view dependency insights setting.
	Message *string `json:"message"`
}

// Autogenerated input type of UpdateEnterpriseOrganizationProjectsSetting
type UpdateEnterpriseOrganizationProjectsSettingInput struct {
	// The ID of the enterprise on which to set the organization projects setting.
	EnterpriseID string `json:"enterpriseId"`
	// The value for the organization projects setting on the enterprise.
	SettingValue EnterpriseEnabledDisabledSettingValue `json:"settingValue"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of UpdateEnterpriseOrganizationProjectsSetting
type UpdateEnterpriseOrganizationProjectsSettingPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The enterprise with the updated organization projects setting.
	Enterprise *Enterprise `json:"enterprise"`
	// A message confirming the result of updating the organization projects setting.
	Message *string `json:"message"`
}

// Autogenerated input type of UpdateEnterpriseProfile
type UpdateEnterpriseProfileInput struct {
	// The Enterprise ID to update.
	EnterpriseID string `json:"enterpriseId"`
	// The name of the enterprise.
	Name *string `json:"name"`
	// The description of the enterprise.
	Description *string `json:"description"`
	// The URL of the enterprise's website.
	WebsiteURL *string `json:"websiteUrl"`
	// The location of the enterprise.
	Location *string `json:"location"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of UpdateEnterpriseProfile
type UpdateEnterpriseProfilePayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The updated enterprise.
	Enterprise *Enterprise `json:"enterprise"`
}

// Autogenerated input type of UpdateEnterpriseRepositoryProjectsSetting
type UpdateEnterpriseRepositoryProjectsSettingInput struct {
	// The ID of the enterprise on which to set the repository projects setting.
	EnterpriseID string `json:"enterpriseId"`
	// The value for the repository projects setting on the enterprise.
	SettingValue EnterpriseEnabledDisabledSettingValue `json:"settingValue"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of UpdateEnterpriseRepositoryProjectsSetting
type UpdateEnterpriseRepositoryProjectsSettingPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The enterprise with the updated repository projects setting.
	Enterprise *Enterprise `json:"enterprise"`
	// A message confirming the result of updating the repository projects setting.
	Message *string `json:"message"`
}

// Autogenerated input type of UpdateEnterpriseTeamDiscussionsSetting
type UpdateEnterpriseTeamDiscussionsSettingInput struct {
	// The ID of the enterprise on which to set the team discussions setting.
	EnterpriseID string `json:"enterpriseId"`
	// The value for the team discussions setting on the enterprise.
	SettingValue EnterpriseEnabledDisabledSettingValue `json:"settingValue"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of UpdateEnterpriseTeamDiscussionsSetting
type UpdateEnterpriseTeamDiscussionsSettingPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The enterprise with the updated team discussions setting.
	Enterprise *Enterprise `json:"enterprise"`
	// A message confirming the result of updating the team discussions setting.
	Message *string `json:"message"`
}

// Autogenerated input type of UpdateEnterpriseTwoFactorAuthenticationRequiredSetting
type UpdateEnterpriseTwoFactorAuthenticationRequiredSettingInput struct {
	// The ID of the enterprise on which to set the two factor authentication required setting.
	EnterpriseID string `json:"enterpriseId"`
	// The value for the two factor authentication required setting on the enterprise.
	SettingValue EnterpriseEnabledSettingValue `json:"settingValue"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of UpdateEnterpriseTwoFactorAuthenticationRequiredSetting
type UpdateEnterpriseTwoFactorAuthenticationRequiredSettingPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The enterprise with the updated two factor authentication required setting.
	Enterprise *Enterprise `json:"enterprise"`
	// A message confirming the result of updating the two factor authentication required setting.
	Message *string `json:"message"`
}

// Autogenerated input type of UpdateIpAllowListEnabledSetting
type UpdateIPAllowListEnabledSettingInput struct {
	// The ID of the owner on which to set the IP allow list enabled setting.
	OwnerID string `json:"ownerId"`
	// The value for the IP allow list enabled setting.
	SettingValue IPAllowListEnabledSettingValue `json:"settingValue"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of UpdateIpAllowListEnabledSetting
type UpdateIPAllowListEnabledSettingPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The IP allow list owner on which the setting was updated.
	Owner IPAllowListOwner `json:"owner"`
}

// Autogenerated input type of UpdateIpAllowListEntry
type UpdateIPAllowListEntryInput struct {
	// The ID of the IP allow list entry to update.
	IPAllowListEntryID string `json:"ipAllowListEntryId"`
	// An IP address or range of addresses in CIDR notation.
	AllowListValue string `json:"allowListValue"`
	// An optional name for the IP allow list entry.
	Name *string `json:"name"`
	// Whether the IP allow list entry is active when an IP allow list is enabled.
	IsActive bool `json:"isActive"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of UpdateIpAllowListEntry
type UpdateIPAllowListEntryPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The IP allow list entry that was updated.
	IPAllowListEntry *IPAllowListEntry `json:"ipAllowListEntry"`
}

// Autogenerated input type of UpdateIssueComment
type UpdateIssueCommentInput struct {
	// The ID of the IssueComment to modify.
	ID string `json:"id"`
	// The updated text of the comment.
	Body string `json:"body"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of UpdateIssueComment
type UpdateIssueCommentPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The updated comment.
	IssueComment *IssueComment `json:"issueComment"`
}

// Autogenerated input type of UpdateIssue
type UpdateIssueInput struct {
	// The ID of the Issue to modify.
	ID string `json:"id"`
	// The title for the issue.
	Title *string `json:"title"`
	// The body for the issue description.
	Body *string `json:"body"`
	// An array of Node IDs of users for this issue.
	AssigneeIds []string `json:"assigneeIds"`
	// The Node ID of the milestone for this issue.
	MilestoneID *string `json:"milestoneId"`
	// An array of Node IDs of labels for this issue.
	LabelIds []string `json:"labelIds"`
	// The desired issue state.
	State *IssueState `json:"state"`
	// An array of Node IDs for projects associated with this issue.
	ProjectIds []string `json:"projectIds"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of UpdateIssue
type UpdateIssuePayload struct {
	// Identifies the actor who performed the event.
	Actor Actor `json:"actor"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The issue.
	Issue *Issue `json:"issue"`
}

// Autogenerated input type of UpdateProjectCard
type UpdateProjectCardInput struct {
	// The ProjectCard ID to update.
	ProjectCardID string `json:"projectCardId"`
	// Whether or not the ProjectCard should be archived
	IsArchived *bool `json:"isArchived"`
	// The note of ProjectCard.
	Note *string `json:"note"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of UpdateProjectCard
type UpdateProjectCardPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The updated ProjectCard.
	ProjectCard *ProjectCard `json:"projectCard"`
}

// Autogenerated input type of UpdateProjectColumn
type UpdateProjectColumnInput struct {
	// The ProjectColumn ID to update.
	ProjectColumnID string `json:"projectColumnId"`
	// The name of project column.
	Name string `json:"name"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of UpdateProjectColumn
type UpdateProjectColumnPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The updated project column.
	ProjectColumn *ProjectColumn `json:"projectColumn"`
}

// Autogenerated input type of UpdateProject
type UpdateProjectInput struct {
	// The Project ID to update.
	ProjectID string `json:"projectId"`
	// The name of project.
	Name *string `json:"name"`
	// The description of project.
	Body *string `json:"body"`
	// Whether the project is open or closed.
	State *ProjectState `json:"state"`
	// Whether the project is public or not.
	Public *bool `json:"public"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of UpdateProject
type UpdateProjectPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The updated project.
	Project *Project `json:"project"`
}

// Autogenerated input type of UpdatePullRequest
type UpdatePullRequestInput struct {
	// The Node ID of the pull request.
	PullRequestID string `json:"pullRequestId"`
	// The name of the branch you want your changes pulled into. This should be an existing branch
	// on the current repository.
	//
	BaseRefName *string `json:"baseRefName"`
	// The title of the pull request.
	Title *string `json:"title"`
	// The contents of the pull request.
	Body *string `json:"body"`
	// The target state of the pull request.
	State *PullRequestUpdateState `json:"state"`
	// Indicates whether maintainers can modify the pull request.
	MaintainerCanModify *bool `json:"maintainerCanModify"`
	// An array of Node IDs of users for this pull request.
	AssigneeIds []string `json:"assigneeIds"`
	// The Node ID of the milestone for this pull request.
	MilestoneID *string `json:"milestoneId"`
	// An array of Node IDs of labels for this pull request.
	LabelIds []string `json:"labelIds"`
	// An array of Node IDs for projects associated with this pull request.
	ProjectIds []string `json:"projectIds"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of UpdatePullRequest
type UpdatePullRequestPayload struct {
	// Identifies the actor who performed the event.
	Actor Actor `json:"actor"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The updated pull request.
	PullRequest *PullRequest `json:"pullRequest"`
}

// Autogenerated input type of UpdatePullRequestReviewComment
type UpdatePullRequestReviewCommentInput struct {
	// The Node ID of the comment to modify.
	PullRequestReviewCommentID string `json:"pullRequestReviewCommentId"`
	// The text of the comment.
	Body string `json:"body"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of UpdatePullRequestReviewComment
type UpdatePullRequestReviewCommentPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The updated comment.
	PullRequestReviewComment *PullRequestReviewComment `json:"pullRequestReviewComment"`
}

// Autogenerated input type of UpdatePullRequestReview
type UpdatePullRequestReviewInput struct {
	// The Node ID of the pull request review to modify.
	PullRequestReviewID string `json:"pullRequestReviewId"`
	// The contents of the pull request review body.
	Body string `json:"body"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of UpdatePullRequestReview
type UpdatePullRequestReviewPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The updated pull request review.
	PullRequestReview *PullRequestReview `json:"pullRequestReview"`
}

// Autogenerated input type of UpdateRef
type UpdateRefInput struct {
	// The Node ID of the Ref to be updated.
	RefID string `json:"refId"`
	// The GitObjectID that the Ref shall be updated to target.
	Oid string `json:"oid"`
	// Permit updates of branch Refs that are not fast-forwards?
	Force *bool `json:"force"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of UpdateRef
type UpdateRefPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The updated Ref.
	Ref *Ref `json:"ref"`
}

// Autogenerated input type of UpdateRepository
type UpdateRepositoryInput struct {
	// The ID of the repository to update.
	RepositoryID string `json:"repositoryId"`
	// The new name of the repository.
	Name *string `json:"name"`
	// A new description for the repository. Pass an empty string to erase the existing description.
	Description *string `json:"description"`
	// Whether this repository should be marked as a template such that anyone who can access it can create new repositories with the same files and directory structure.
	Template *bool `json:"template"`
	// The URL for a web page about this repository. Pass an empty string to erase the existing URL.
	HomepageURL *string `json:"homepageUrl"`
	// Indicates if the repository should have the wiki feature enabled.
	HasWikiEnabled *bool `json:"hasWikiEnabled"`
	// Indicates if the repository should have the issues feature enabled.
	HasIssuesEnabled *bool `json:"hasIssuesEnabled"`
	// Indicates if the repository should have the project boards feature enabled.
	HasProjectsEnabled *bool `json:"hasProjectsEnabled"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of UpdateRepository
type UpdateRepositoryPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The updated repository.
	Repository *Repository `json:"repository"`
}

// Autogenerated input type of UpdateSubscription
type UpdateSubscriptionInput struct {
	// The Node ID of the subscribable object to modify.
	SubscribableID string `json:"subscribableId"`
	// The new state of the subscription.
	State SubscriptionState `json:"state"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of UpdateSubscription
type UpdateSubscriptionPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The input subscribable entity.
	Subscribable Subscribable `json:"subscribable"`
}

// Autogenerated input type of UpdateTeamDiscussionComment
type UpdateTeamDiscussionCommentInput struct {
	// The ID of the comment to modify.
	ID string `json:"id"`
	// The updated text of the comment.
	Body string `json:"body"`
	// The current version of the body content.
	BodyVersion *string `json:"bodyVersion"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of UpdateTeamDiscussionComment
type UpdateTeamDiscussionCommentPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The updated comment.
	TeamDiscussionComment *TeamDiscussionComment `json:"teamDiscussionComment"`
}

// Autogenerated input type of UpdateTeamDiscussion
type UpdateTeamDiscussionInput struct {
	// The Node ID of the discussion to modify.
	ID string `json:"id"`
	// The updated title of the discussion.
	Title *string `json:"title"`
	// The updated text of the discussion.
	Body *string `json:"body"`
	// The current version of the body content. If provided, this update operation will be rejected if the given version does not match the latest version on the server.
	BodyVersion *string `json:"bodyVersion"`
	// If provided, sets the pinned state of the updated discussion.
	Pinned *bool `json:"pinned"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of UpdateTeamDiscussion
type UpdateTeamDiscussionPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The updated discussion.
	TeamDiscussion *TeamDiscussion `json:"teamDiscussion"`
}

// Autogenerated input type of UpdateTopics
type UpdateTopicsInput struct {
	// The Node ID of the repository.
	RepositoryID string `json:"repositoryId"`
	// An array of topic names.
	TopicNames []string `json:"topicNames"`
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
}

// Autogenerated return type of UpdateTopics
type UpdateTopicsPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// Names of the provided topics that are not valid.
	InvalidTopicNames []string `json:"invalidTopicNames"`
	// The updated repository.
	Repository *Repository `json:"repository"`
}

// A user is an individual's account on GitHub that owns repositories and can make new content.
type User struct {
	// Determine if this repository owner has any items that can be pinned to their profile.
	AnyPinnableItems bool `json:"anyPinnableItems"`
	// A URL pointing to the user's public avatar.
	AvatarURL string `json:"avatarUrl"`
	// The user's public profile bio.
	Bio *string `json:"bio"`
	// The user's public profile bio as HTML.
	BioHTML string `json:"bioHTML"`
	// A list of commit comments made by this user.
	CommitComments *CommitCommentConnection `json:"commitComments"`
	// The user's public profile company.
	Company *string `json:"company"`
	// The user's public profile company as HTML.
	CompanyHTML string `json:"companyHTML"`
	// The collection of contributions this user has made to different repositories.
	ContributionsCollection *ContributionsCollection `json:"contributionsCollection"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	// Identifies the primary key from the database.
	DatabaseID *int `json:"databaseId"`
	// The user's publicly visible profile email.
	Email string `json:"email"`
	// A list of users the given user is followed by.
	Followers *FollowerConnection `json:"followers"`
	// A list of users the given user is following.
	Following *FollowingConnection `json:"following"`
	// Find gist by repo name.
	Gist *Gist `json:"gist"`
	// A list of gist comments made by this user.
	GistComments *GistCommentConnection `json:"gistComments"`
	// A list of the Gists the user has created.
	Gists *GistConnection `json:"gists"`
	// The hovercard information for this user in a given context
	Hovercard *Hovercard `json:"hovercard"`
	ID        string     `json:"id"`
	// Whether or not this user is a participant in the GitHub Security Bug Bounty.
	IsBountyHunter bool `json:"isBountyHunter"`
	// Whether or not this user is a participant in the GitHub Campus Experts Program.
	IsCampusExpert bool `json:"isCampusExpert"`
	// Whether or not this user is a GitHub Developer Program member.
	IsDeveloperProgramMember bool `json:"isDeveloperProgramMember"`
	// Whether or not this user is a GitHub employee.
	IsEmployee bool `json:"isEmployee"`
	// Whether or not the user has marked themselves as for hire.
	IsHireable bool `json:"isHireable"`
	// Whether or not this user is a site administrator.
	IsSiteAdmin bool `json:"isSiteAdmin"`
	// Whether or not this user is the viewing user.
	IsViewer bool `json:"isViewer"`
	// A list of issue comments made by this user.
	IssueComments *IssueCommentConnection `json:"issueComments"`
	// A list of issues associated with this user.
	Issues *IssueConnection `json:"issues"`
	// Showcases a selection of repositories and gists that the profile owner has either curated or that have been selected automatically based on popularity.
	ItemShowcase *ProfileItemShowcase `json:"itemShowcase"`
	// The user's public profile location.
	Location *string `json:"location"`
	// The username used to login.
	Login string `json:"login"`
	// The user's public profile name.
	Name *string `json:"name"`
	// Find an organization by its login that the user belongs to.
	Organization *Organization `json:"organization"`
	// Verified email addresses that match verified domains for a specified organization the user is a member of.
	OrganizationVerifiedDomainEmails []string `json:"organizationVerifiedDomainEmails"`
	// A list of organizations the user belongs to.
	Organizations *OrganizationConnection `json:"organizations"`
	// A list of packages under the owner.
	Packages *PackageConnection `json:"packages"`
	// A list of repositories and gists this profile owner can pin to their profile.
	PinnableItems *PinnableItemConnection `json:"pinnableItems"`
	// A list of repositories and gists this profile owner has pinned to their profile
	PinnedItems *PinnableItemConnection `json:"pinnedItems"`
	// Returns how many more items this profile owner can pin to their profile.
	PinnedItemsRemaining int `json:"pinnedItemsRemaining"`
	// Find project by number.
	Project *Project `json:"project"`
	// A list of projects under the owner.
	Projects *ProjectConnection `json:"projects"`
	// The HTTP path listing user's projects
	ProjectsResourcePath string `json:"projectsResourcePath"`
	// The HTTP URL listing user's projects
	ProjectsURL string `json:"projectsUrl"`
	// A list of public keys associated with this user.
	PublicKeys *PublicKeyConnection `json:"publicKeys"`
	// A list of pull requests associated with this user.
	PullRequests *PullRequestConnection `json:"pullRequests"`
	// A list of repositories that the user owns.
	Repositories *RepositoryConnection `json:"repositories"`
	// A list of repositories that the user recently contributed to.
	RepositoriesContributedTo *RepositoryConnection `json:"repositoriesContributedTo"`
	// Find Repository.
	Repository *Repository `json:"repository"`
	// The HTTP path for this user
	ResourcePath string `json:"resourcePath"`
	// Replies this user has saved
	SavedReplies *SavedReplyConnection `json:"savedReplies"`
	// The GitHub Sponsors listing for this user.
	SponsorsListing *SponsorsListing `json:"sponsorsListing"`
	// This object's sponsorships as the maintainer.
	SponsorshipsAsMaintainer *SponsorshipConnection `json:"sponsorshipsAsMaintainer"`
	// This object's sponsorships as the sponsor.
	SponsorshipsAsSponsor *SponsorshipConnection `json:"sponsorshipsAsSponsor"`
	// Repositories the user has starred.
	StarredRepositories *StarredRepositoryConnection `json:"starredRepositories"`
	// The user's description of what they're currently doing.
	Status *UserStatus `json:"status"`
	// Repositories the user has contributed to, ordered by contribution rank, plus repositories the user has created
	//
	TopRepositories *RepositoryConnection `json:"topRepositories"`
	// The user's Twitter username.
	TwitterUsername *string `json:"twitterUsername"`
	// Identifies the date and time when the object was last updated.
	UpdatedAt string `json:"updatedAt"`
	// The HTTP URL for this user
	URL string `json:"url"`
	// Can the viewer pin repositories and gists to the profile?
	ViewerCanChangePinnedItems bool `json:"viewerCanChangePinnedItems"`
	// Can the current viewer create new projects on this owner.
	ViewerCanCreateProjects bool `json:"viewerCanCreateProjects"`
	// Whether or not the viewer is able to follow the user.
	ViewerCanFollow bool `json:"viewerCanFollow"`
	// Whether or not this user is followed by the viewer.
	ViewerIsFollowing bool `json:"viewerIsFollowing"`
	// A list of repositories the given user is watching.
	Watching *RepositoryConnection `json:"watching"`
	// A URL pointing to the user's public website/blog.
	WebsiteURL *string `json:"websiteUrl"`
}

func (User) IsNode()                          {}
func (User) IsActor()                         {}
func (User) IsPackageOwner()                  {}
func (User) IsProjectOwner()                  {}
func (User) IsRepositoryOwner()               {}
func (User) IsUniformResourceLocatable()      {}
func (User) IsProfileOwner()                  {}
func (User) IsSponsorable()                   {}
func (User) IsEnterpriseMember()              {}
func (User) IsSearchResultItem()              {}
func (User) IsSponsor()                       {}
func (User) IsRequestedReviewer()             {}
func (User) IsAuditEntryActor()               {}
func (User) IsReviewDismissalAllowanceActor() {}
func (User) IsPushAllowanceActor()            {}
func (User) IsAssignee()                      {}

// Represents a 'user_blocked' event on a given user.
type UserBlockedEvent struct {
	// Identifies the actor who performed the event.
	Actor Actor `json:"actor"`
	// Number of days that the user was blocked for.
	BlockDuration UserBlockDuration `json:"blockDuration"`
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	// The user who was blocked.
	Subject *User `json:"subject"`
}

func (UserBlockedEvent) IsPullRequestTimelineItems() {}
func (UserBlockedEvent) IsIssueTimelineItems()       {}
func (UserBlockedEvent) IsPullRequestTimelineItem()  {}
func (UserBlockedEvent) IsNode()                     {}
func (UserBlockedEvent) IsIssueTimelineItem()        {}

// The connection type for User.
type UserConnection struct {
	// A list of edges.
	Edges []*UserEdge `json:"edges"`
	// A list of nodes.
	Nodes []*User `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edit on user content
type UserContentEdit struct {
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	// Identifies the date and time when the object was deleted.
	DeletedAt *string `json:"deletedAt"`
	// The actor who deleted this content
	DeletedBy Actor `json:"deletedBy"`
	// A summary of the changes for this edit
	Diff *string `json:"diff"`
	// When this content was edited
	EditedAt string `json:"editedAt"`
	// The actor who edited this content
	Editor Actor  `json:"editor"`
	ID     string `json:"id"`
	// Identifies the date and time when the object was last updated.
	UpdatedAt string `json:"updatedAt"`
}

func (UserContentEdit) IsNode() {}

// A list of edits to content.
type UserContentEditConnection struct {
	// A list of edges.
	Edges []*UserContentEditEdge `json:"edges"`
	// A list of nodes.
	Nodes []*UserContentEdit `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type UserContentEditEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *UserContentEdit `json:"node"`
}

// Represents a user.
type UserEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *User `json:"node"`
}

// Email attributes from External Identity
type UserEmailMetadata struct {
	// Boolean to identify primary emails
	Primary *bool `json:"primary"`
	// Type of email
	Type *string `json:"type"`
	// Email id
	Value string `json:"value"`
}

// The user's description of what they're currently doing.
type UserStatus struct {
	// Identifies the date and time when the object was created.
	CreatedAt string `json:"createdAt"`
	// An emoji summarizing the user's status.
	Emoji *string `json:"emoji"`
	// The status emoji as HTML.
	EmojiHTML *string `json:"emojiHTML"`
	// If set, the status will not be shown after this date.
	ExpiresAt *string `json:"expiresAt"`
	// ID of the object.
	ID string `json:"id"`
	// Whether this status indicates the user is not fully available on GitHub.
	IndicatesLimitedAvailability bool `json:"indicatesLimitedAvailability"`
	// A brief message describing what the user is doing.
	Message *string `json:"message"`
	// The organization whose members can see this status. If null, this status is publicly visible.
	Organization *Organization `json:"organization"`
	// Identifies the date and time when the object was last updated.
	UpdatedAt string `json:"updatedAt"`
	// The user who has this status.
	User *User `json:"user"`
}

func (UserStatus) IsNode() {}

// The connection type for UserStatus.
type UserStatusConnection struct {
	// A list of edges.
	Edges []*UserStatusEdge `json:"edges"`
	// A list of nodes.
	Nodes []*UserStatus `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type UserStatusEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *UserStatus `json:"node"`
}

// Ordering options for user status connections.
type UserStatusOrder struct {
	// The field to order user statuses by.
	Field UserStatusOrderField `json:"field"`
	// The ordering direction.
	Direction OrderDirection `json:"direction"`
}

// A hovercard context with a message describing how the viewer is related.
type ViewerHovercardContext struct {
	// A string describing this context
	Message string `json:"message"`
	// An octicon to accompany this context
	Octicon string `json:"octicon"`
	// Identifies the user who is related to this context.
	Viewer *User `json:"viewer"`
}

func (ViewerHovercardContext) IsHovercardContext() {}

// The possible capabilities for action executions setting.
type ActionExecutionCapabilitySetting string

const (
	// All action executions are disabled.
	ActionExecutionCapabilitySettingDisabled ActionExecutionCapabilitySetting = "DISABLED"
	// All action executions are enabled.
	ActionExecutionCapabilitySettingAllActions ActionExecutionCapabilitySetting = "ALL_ACTIONS"
	// Only actions defined within the repo are allowed.
	ActionExecutionCapabilitySettingLocalActionsOnly ActionExecutionCapabilitySetting = "LOCAL_ACTIONS_ONLY"
	// Organization administrators action execution capabilities.
	ActionExecutionCapabilitySettingNoPolicy ActionExecutionCapabilitySetting = "NO_POLICY"
)

var AllActionExecutionCapabilitySetting = []ActionExecutionCapabilitySetting{
	ActionExecutionCapabilitySettingDisabled,
	ActionExecutionCapabilitySettingAllActions,
	ActionExecutionCapabilitySettingLocalActionsOnly,
	ActionExecutionCapabilitySettingNoPolicy,
}

func (e ActionExecutionCapabilitySetting) IsValid() bool {
	switch e {
	case ActionExecutionCapabilitySettingDisabled, ActionExecutionCapabilitySettingAllActions, ActionExecutionCapabilitySettingLocalActionsOnly, ActionExecutionCapabilitySettingNoPolicy:
		return true
	}
	return false
}

func (e ActionExecutionCapabilitySetting) String() string {
	return string(e)
}

func (e *ActionExecutionCapabilitySetting) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ActionExecutionCapabilitySetting(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ActionExecutionCapabilitySetting", str)
	}
	return nil
}

func (e ActionExecutionCapabilitySetting) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Properties by which Audit Log connections can be ordered.
type AuditLogOrderField string

const (
	// Order audit log entries by timestamp
	AuditLogOrderFieldCreatedAt AuditLogOrderField = "CREATED_AT"
)

var AllAuditLogOrderField = []AuditLogOrderField{
	AuditLogOrderFieldCreatedAt,
}

func (e AuditLogOrderField) IsValid() bool {
	switch e {
	case AuditLogOrderFieldCreatedAt:
		return true
	}
	return false
}

func (e AuditLogOrderField) String() string {
	return string(e)
}

func (e *AuditLogOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AuditLogOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AuditLogOrderField", str)
	}
	return nil
}

func (e AuditLogOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Represents an annotation's information level.
type CheckAnnotationLevel string

const (
	// An annotation indicating an inescapable error.
	CheckAnnotationLevelFailure CheckAnnotationLevel = "FAILURE"
	// An annotation indicating some information.
	CheckAnnotationLevelNotice CheckAnnotationLevel = "NOTICE"
	// An annotation indicating an ignorable error.
	CheckAnnotationLevelWarning CheckAnnotationLevel = "WARNING"
)

var AllCheckAnnotationLevel = []CheckAnnotationLevel{
	CheckAnnotationLevelFailure,
	CheckAnnotationLevelNotice,
	CheckAnnotationLevelWarning,
}

func (e CheckAnnotationLevel) IsValid() bool {
	switch e {
	case CheckAnnotationLevelFailure, CheckAnnotationLevelNotice, CheckAnnotationLevelWarning:
		return true
	}
	return false
}

func (e CheckAnnotationLevel) String() string {
	return string(e)
}

func (e *CheckAnnotationLevel) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CheckAnnotationLevel(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CheckAnnotationLevel", str)
	}
	return nil
}

func (e CheckAnnotationLevel) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The possible states for a check suite or run conclusion.
type CheckConclusionState string

const (
	// The check suite or run requires action.
	CheckConclusionStateActionRequired CheckConclusionState = "ACTION_REQUIRED"
	// The check suite or run has timed out.
	CheckConclusionStateTimedOut CheckConclusionState = "TIMED_OUT"
	// The check suite or run has been cancelled.
	CheckConclusionStateCancelled CheckConclusionState = "CANCELLED"
	// The check suite or run has failed.
	CheckConclusionStateFailure CheckConclusionState = "FAILURE"
	// The check suite or run has succeeded.
	CheckConclusionStateSuccess CheckConclusionState = "SUCCESS"
	// The check suite or run was neutral.
	CheckConclusionStateNeutral CheckConclusionState = "NEUTRAL"
	// The check suite or run was skipped.
	CheckConclusionStateSkipped CheckConclusionState = "SKIPPED"
	// The check suite or run was marked stale by GitHub. Only GitHub can use this conclusion.
	CheckConclusionStateStale CheckConclusionState = "STALE"
)

var AllCheckConclusionState = []CheckConclusionState{
	CheckConclusionStateActionRequired,
	CheckConclusionStateTimedOut,
	CheckConclusionStateCancelled,
	CheckConclusionStateFailure,
	CheckConclusionStateSuccess,
	CheckConclusionStateNeutral,
	CheckConclusionStateSkipped,
	CheckConclusionStateStale,
}

func (e CheckConclusionState) IsValid() bool {
	switch e {
	case CheckConclusionStateActionRequired, CheckConclusionStateTimedOut, CheckConclusionStateCancelled, CheckConclusionStateFailure, CheckConclusionStateSuccess, CheckConclusionStateNeutral, CheckConclusionStateSkipped, CheckConclusionStateStale:
		return true
	}
	return false
}

func (e CheckConclusionState) String() string {
	return string(e)
}

func (e *CheckConclusionState) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CheckConclusionState(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CheckConclusionState", str)
	}
	return nil
}

func (e CheckConclusionState) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The possible types of check runs.
type CheckRunType string

const (
	// Every check run available.
	CheckRunTypeAll CheckRunType = "ALL"
	// The latest check run.
	CheckRunTypeLatest CheckRunType = "LATEST"
)

var AllCheckRunType = []CheckRunType{
	CheckRunTypeAll,
	CheckRunTypeLatest,
}

func (e CheckRunType) IsValid() bool {
	switch e {
	case CheckRunTypeAll, CheckRunTypeLatest:
		return true
	}
	return false
}

func (e CheckRunType) String() string {
	return string(e)
}

func (e *CheckRunType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CheckRunType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CheckRunType", str)
	}
	return nil
}

func (e CheckRunType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The possible states for a check suite or run status.
type CheckStatusState string

const (
	// The check suite or run has been queued.
	CheckStatusStateQueued CheckStatusState = "QUEUED"
	// The check suite or run is in progress.
	CheckStatusStateInProgress CheckStatusState = "IN_PROGRESS"
	// The check suite or run has been completed.
	CheckStatusStateCompleted CheckStatusState = "COMPLETED"
	// The check suite or run has been requested.
	CheckStatusStateRequested CheckStatusState = "REQUESTED"
)

var AllCheckStatusState = []CheckStatusState{
	CheckStatusStateQueued,
	CheckStatusStateInProgress,
	CheckStatusStateCompleted,
	CheckStatusStateRequested,
}

func (e CheckStatusState) IsValid() bool {
	switch e {
	case CheckStatusStateQueued, CheckStatusStateInProgress, CheckStatusStateCompleted, CheckStatusStateRequested:
		return true
	}
	return false
}

func (e CheckStatusState) String() string {
	return string(e)
}

func (e *CheckStatusState) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CheckStatusState(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CheckStatusState", str)
	}
	return nil
}

func (e CheckStatusState) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Collaborators affiliation level with a subject.
type CollaboratorAffiliation string

const (
	// All outside collaborators of an organization-owned subject.
	CollaboratorAffiliationOutside CollaboratorAffiliation = "OUTSIDE"
	// All collaborators with permissions to an organization-owned subject, regardless of organization membership status.
	CollaboratorAffiliationDirect CollaboratorAffiliation = "DIRECT"
	// All collaborators the authenticated user can see.
	CollaboratorAffiliationAll CollaboratorAffiliation = "ALL"
)

var AllCollaboratorAffiliation = []CollaboratorAffiliation{
	CollaboratorAffiliationOutside,
	CollaboratorAffiliationDirect,
	CollaboratorAffiliationAll,
}

func (e CollaboratorAffiliation) IsValid() bool {
	switch e {
	case CollaboratorAffiliationOutside, CollaboratorAffiliationDirect, CollaboratorAffiliationAll:
		return true
	}
	return false
}

func (e CollaboratorAffiliation) String() string {
	return string(e)
}

func (e *CollaboratorAffiliation) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CollaboratorAffiliation(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CollaboratorAffiliation", str)
	}
	return nil
}

func (e CollaboratorAffiliation) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// A comment author association with repository.
type CommentAuthorAssociation string

const (
	// Author is a member of the organization that owns the repository.
	CommentAuthorAssociationMember CommentAuthorAssociation = "MEMBER"
	// Author is the owner of the repository.
	CommentAuthorAssociationOwner CommentAuthorAssociation = "OWNER"
	// Author is a placeholder for an unclaimed user.
	CommentAuthorAssociationMannequin CommentAuthorAssociation = "MANNEQUIN"
	// Author has been invited to collaborate on the repository.
	CommentAuthorAssociationCollaborator CommentAuthorAssociation = "COLLABORATOR"
	// Author has previously committed to the repository.
	CommentAuthorAssociationContributor CommentAuthorAssociation = "CONTRIBUTOR"
	// Author has not previously committed to the repository.
	CommentAuthorAssociationFirstTimeContributor CommentAuthorAssociation = "FIRST_TIME_CONTRIBUTOR"
	// Author has not previously committed to GitHub.
	CommentAuthorAssociationFirstTimer CommentAuthorAssociation = "FIRST_TIMER"
	// Author has no association with the repository.
	CommentAuthorAssociationNone CommentAuthorAssociation = "NONE"
)

var AllCommentAuthorAssociation = []CommentAuthorAssociation{
	CommentAuthorAssociationMember,
	CommentAuthorAssociationOwner,
	CommentAuthorAssociationMannequin,
	CommentAuthorAssociationCollaborator,
	CommentAuthorAssociationContributor,
	CommentAuthorAssociationFirstTimeContributor,
	CommentAuthorAssociationFirstTimer,
	CommentAuthorAssociationNone,
}

func (e CommentAuthorAssociation) IsValid() bool {
	switch e {
	case CommentAuthorAssociationMember, CommentAuthorAssociationOwner, CommentAuthorAssociationMannequin, CommentAuthorAssociationCollaborator, CommentAuthorAssociationContributor, CommentAuthorAssociationFirstTimeContributor, CommentAuthorAssociationFirstTimer, CommentAuthorAssociationNone:
		return true
	}
	return false
}

func (e CommentAuthorAssociation) String() string {
	return string(e)
}

func (e *CommentAuthorAssociation) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CommentAuthorAssociation(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CommentAuthorAssociation", str)
	}
	return nil
}

func (e CommentAuthorAssociation) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The possible errors that will prevent a user from updating a comment.
type CommentCannotUpdateReason string

const (
	// Unable to create comment because repository is archived.
	CommentCannotUpdateReasonArchived CommentCannotUpdateReason = "ARCHIVED"
	// You must be the author or have write access to this repository to update this comment.
	CommentCannotUpdateReasonInsufficientAccess CommentCannotUpdateReason = "INSUFFICIENT_ACCESS"
	// Unable to create comment because issue is locked.
	CommentCannotUpdateReasonLocked CommentCannotUpdateReason = "LOCKED"
	// You must be logged in to update this comment.
	CommentCannotUpdateReasonLoginRequired CommentCannotUpdateReason = "LOGIN_REQUIRED"
	// Repository is under maintenance.
	CommentCannotUpdateReasonMaintenance CommentCannotUpdateReason = "MAINTENANCE"
	// At least one email address must be verified to update this comment.
	CommentCannotUpdateReasonVerifiedEmailRequired CommentCannotUpdateReason = "VERIFIED_EMAIL_REQUIRED"
	// You cannot update this comment
	CommentCannotUpdateReasonDenied CommentCannotUpdateReason = "DENIED"
)

var AllCommentCannotUpdateReason = []CommentCannotUpdateReason{
	CommentCannotUpdateReasonArchived,
	CommentCannotUpdateReasonInsufficientAccess,
	CommentCannotUpdateReasonLocked,
	CommentCannotUpdateReasonLoginRequired,
	CommentCannotUpdateReasonMaintenance,
	CommentCannotUpdateReasonVerifiedEmailRequired,
	CommentCannotUpdateReasonDenied,
}

func (e CommentCannotUpdateReason) IsValid() bool {
	switch e {
	case CommentCannotUpdateReasonArchived, CommentCannotUpdateReasonInsufficientAccess, CommentCannotUpdateReasonLocked, CommentCannotUpdateReasonLoginRequired, CommentCannotUpdateReasonMaintenance, CommentCannotUpdateReasonVerifiedEmailRequired, CommentCannotUpdateReasonDenied:
		return true
	}
	return false
}

func (e CommentCannotUpdateReason) String() string {
	return string(e)
}

func (e *CommentCannotUpdateReason) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CommentCannotUpdateReason(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CommentCannotUpdateReason", str)
	}
	return nil
}

func (e CommentCannotUpdateReason) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Properties by which commit contribution connections can be ordered.
type CommitContributionOrderField string

const (
	// Order commit contributions by when they were made.
	CommitContributionOrderFieldOccurredAt CommitContributionOrderField = "OCCURRED_AT"
	// Order commit contributions by how many commits they represent.
	CommitContributionOrderFieldCommitCount CommitContributionOrderField = "COMMIT_COUNT"
)

var AllCommitContributionOrderField = []CommitContributionOrderField{
	CommitContributionOrderFieldOccurredAt,
	CommitContributionOrderFieldCommitCount,
}

func (e CommitContributionOrderField) IsValid() bool {
	switch e {
	case CommitContributionOrderFieldOccurredAt, CommitContributionOrderFieldCommitCount:
		return true
	}
	return false
}

func (e CommitContributionOrderField) String() string {
	return string(e)
}

func (e *CommitContributionOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CommitContributionOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CommitContributionOrderField", str)
	}
	return nil
}

func (e CommitContributionOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The possible default permissions for repositories.
type DefaultRepositoryPermissionField string

const (
	// No access
	DefaultRepositoryPermissionFieldNone DefaultRepositoryPermissionField = "NONE"
	// Can read repos by default
	DefaultRepositoryPermissionFieldRead DefaultRepositoryPermissionField = "READ"
	// Can read and write repos by default
	DefaultRepositoryPermissionFieldWrite DefaultRepositoryPermissionField = "WRITE"
	// Can read, write, and administrate repos by default
	DefaultRepositoryPermissionFieldAdmin DefaultRepositoryPermissionField = "ADMIN"
)

var AllDefaultRepositoryPermissionField = []DefaultRepositoryPermissionField{
	DefaultRepositoryPermissionFieldNone,
	DefaultRepositoryPermissionFieldRead,
	DefaultRepositoryPermissionFieldWrite,
	DefaultRepositoryPermissionFieldAdmin,
}

func (e DefaultRepositoryPermissionField) IsValid() bool {
	switch e {
	case DefaultRepositoryPermissionFieldNone, DefaultRepositoryPermissionFieldRead, DefaultRepositoryPermissionFieldWrite, DefaultRepositoryPermissionFieldAdmin:
		return true
	}
	return false
}

func (e DefaultRepositoryPermissionField) String() string {
	return string(e)
}

func (e *DefaultRepositoryPermissionField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DefaultRepositoryPermissionField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DefaultRepositoryPermissionField", str)
	}
	return nil
}

func (e DefaultRepositoryPermissionField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Properties by which deployment connections can be ordered.
type DeploymentOrderField string

const (
	// Order collection by creation time
	DeploymentOrderFieldCreatedAt DeploymentOrderField = "CREATED_AT"
)

var AllDeploymentOrderField = []DeploymentOrderField{
	DeploymentOrderFieldCreatedAt,
}

func (e DeploymentOrderField) IsValid() bool {
	switch e {
	case DeploymentOrderFieldCreatedAt:
		return true
	}
	return false
}

func (e DeploymentOrderField) String() string {
	return string(e)
}

func (e *DeploymentOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DeploymentOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DeploymentOrderField", str)
	}
	return nil
}

func (e DeploymentOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The possible states in which a deployment can be.
type DeploymentState string

const (
	// The pending deployment was not updated after 30 minutes.
	DeploymentStateAbandoned DeploymentState = "ABANDONED"
	// The deployment is currently active.
	DeploymentStateActive DeploymentState = "ACTIVE"
	// An inactive transient deployment.
	DeploymentStateDestroyed DeploymentState = "DESTROYED"
	// The deployment experienced an error.
	DeploymentStateError DeploymentState = "ERROR"
	// The deployment has failed.
	DeploymentStateFailure DeploymentState = "FAILURE"
	// The deployment is inactive.
	DeploymentStateInactive DeploymentState = "INACTIVE"
	// The deployment is pending.
	DeploymentStatePending DeploymentState = "PENDING"
	// The deployment has queued
	DeploymentStateQueued DeploymentState = "QUEUED"
	// The deployment is in progress.
	DeploymentStateInProgress DeploymentState = "IN_PROGRESS"
)

var AllDeploymentState = []DeploymentState{
	DeploymentStateAbandoned,
	DeploymentStateActive,
	DeploymentStateDestroyed,
	DeploymentStateError,
	DeploymentStateFailure,
	DeploymentStateInactive,
	DeploymentStatePending,
	DeploymentStateQueued,
	DeploymentStateInProgress,
}

func (e DeploymentState) IsValid() bool {
	switch e {
	case DeploymentStateAbandoned, DeploymentStateActive, DeploymentStateDestroyed, DeploymentStateError, DeploymentStateFailure, DeploymentStateInactive, DeploymentStatePending, DeploymentStateQueued, DeploymentStateInProgress:
		return true
	}
	return false
}

func (e DeploymentState) String() string {
	return string(e)
}

func (e *DeploymentState) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DeploymentState(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DeploymentState", str)
	}
	return nil
}

func (e DeploymentState) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The possible states for a deployment status.
type DeploymentStatusState string

const (
	// The deployment is pending.
	DeploymentStatusStatePending DeploymentStatusState = "PENDING"
	// The deployment was successful.
	DeploymentStatusStateSuccess DeploymentStatusState = "SUCCESS"
	// The deployment has failed.
	DeploymentStatusStateFailure DeploymentStatusState = "FAILURE"
	// The deployment is inactive.
	DeploymentStatusStateInactive DeploymentStatusState = "INACTIVE"
	// The deployment experienced an error.
	DeploymentStatusStateError DeploymentStatusState = "ERROR"
	// The deployment is queued
	DeploymentStatusStateQueued DeploymentStatusState = "QUEUED"
	// The deployment is in progress.
	DeploymentStatusStateInProgress DeploymentStatusState = "IN_PROGRESS"
)

var AllDeploymentStatusState = []DeploymentStatusState{
	DeploymentStatusStatePending,
	DeploymentStatusStateSuccess,
	DeploymentStatusStateFailure,
	DeploymentStatusStateInactive,
	DeploymentStatusStateError,
	DeploymentStatusStateQueued,
	DeploymentStatusStateInProgress,
}

func (e DeploymentStatusState) IsValid() bool {
	switch e {
	case DeploymentStatusStatePending, DeploymentStatusStateSuccess, DeploymentStatusStateFailure, DeploymentStatusStateInactive, DeploymentStatusStateError, DeploymentStatusStateQueued, DeploymentStatusStateInProgress:
		return true
	}
	return false
}

func (e DeploymentStatusState) String() string {
	return string(e)
}

func (e *DeploymentStatusState) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DeploymentStatusState(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DeploymentStatusState", str)
	}
	return nil
}

func (e DeploymentStatusState) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The possible sides of a diff.
type DiffSide string

const (
	// The left side of the diff.
	DiffSideLeft DiffSide = "LEFT"
	// The right side of the diff.
	DiffSideRight DiffSide = "RIGHT"
)

var AllDiffSide = []DiffSide{
	DiffSideLeft,
	DiffSideRight,
}

func (e DiffSide) IsValid() bool {
	switch e {
	case DiffSideLeft, DiffSideRight:
		return true
	}
	return false
}

func (e DiffSide) String() string {
	return string(e)
}

func (e *DiffSide) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DiffSide(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DiffSide", str)
	}
	return nil
}

func (e DiffSide) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Properties by which enterprise administrator invitation connections can be ordered.
type EnterpriseAdministratorInvitationOrderField string

const (
	// Order enterprise administrator member invitations by creation time
	EnterpriseAdministratorInvitationOrderFieldCreatedAt EnterpriseAdministratorInvitationOrderField = "CREATED_AT"
)

var AllEnterpriseAdministratorInvitationOrderField = []EnterpriseAdministratorInvitationOrderField{
	EnterpriseAdministratorInvitationOrderFieldCreatedAt,
}

func (e EnterpriseAdministratorInvitationOrderField) IsValid() bool {
	switch e {
	case EnterpriseAdministratorInvitationOrderFieldCreatedAt:
		return true
	}
	return false
}

func (e EnterpriseAdministratorInvitationOrderField) String() string {
	return string(e)
}

func (e *EnterpriseAdministratorInvitationOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EnterpriseAdministratorInvitationOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EnterpriseAdministratorInvitationOrderField", str)
	}
	return nil
}

func (e EnterpriseAdministratorInvitationOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The possible administrator roles in an enterprise account.
type EnterpriseAdministratorRole string

const (
	// Represents an owner of the enterprise account.
	EnterpriseAdministratorRoleOwner EnterpriseAdministratorRole = "OWNER"
	// Represents a billing manager of the enterprise account.
	EnterpriseAdministratorRoleBillingManager EnterpriseAdministratorRole = "BILLING_MANAGER"
)

var AllEnterpriseAdministratorRole = []EnterpriseAdministratorRole{
	EnterpriseAdministratorRoleOwner,
	EnterpriseAdministratorRoleBillingManager,
}

func (e EnterpriseAdministratorRole) IsValid() bool {
	switch e {
	case EnterpriseAdministratorRoleOwner, EnterpriseAdministratorRoleBillingManager:
		return true
	}
	return false
}

func (e EnterpriseAdministratorRole) String() string {
	return string(e)
}

func (e *EnterpriseAdministratorRole) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EnterpriseAdministratorRole(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EnterpriseAdministratorRole", str)
	}
	return nil
}

func (e EnterpriseAdministratorRole) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The possible values for the enterprise default repository permission setting.
type EnterpriseDefaultRepositoryPermissionSettingValue string

const (
	// Organizations in the enterprise choose default repository permissions for their members.
	EnterpriseDefaultRepositoryPermissionSettingValueNoPolicy EnterpriseDefaultRepositoryPermissionSettingValue = "NO_POLICY"
	// Organization members will be able to clone, pull, push, and add new collaborators to all organization repositories.
	EnterpriseDefaultRepositoryPermissionSettingValueAdmin EnterpriseDefaultRepositoryPermissionSettingValue = "ADMIN"
	// Organization members will be able to clone, pull, and push all organization repositories.
	EnterpriseDefaultRepositoryPermissionSettingValueWrite EnterpriseDefaultRepositoryPermissionSettingValue = "WRITE"
	// Organization members will be able to clone and pull all organization repositories.
	EnterpriseDefaultRepositoryPermissionSettingValueRead EnterpriseDefaultRepositoryPermissionSettingValue = "READ"
	// Organization members will only be able to clone and pull public repositories.
	EnterpriseDefaultRepositoryPermissionSettingValueNone EnterpriseDefaultRepositoryPermissionSettingValue = "NONE"
)

var AllEnterpriseDefaultRepositoryPermissionSettingValue = []EnterpriseDefaultRepositoryPermissionSettingValue{
	EnterpriseDefaultRepositoryPermissionSettingValueNoPolicy,
	EnterpriseDefaultRepositoryPermissionSettingValueAdmin,
	EnterpriseDefaultRepositoryPermissionSettingValueWrite,
	EnterpriseDefaultRepositoryPermissionSettingValueRead,
	EnterpriseDefaultRepositoryPermissionSettingValueNone,
}

func (e EnterpriseDefaultRepositoryPermissionSettingValue) IsValid() bool {
	switch e {
	case EnterpriseDefaultRepositoryPermissionSettingValueNoPolicy, EnterpriseDefaultRepositoryPermissionSettingValueAdmin, EnterpriseDefaultRepositoryPermissionSettingValueWrite, EnterpriseDefaultRepositoryPermissionSettingValueRead, EnterpriseDefaultRepositoryPermissionSettingValueNone:
		return true
	}
	return false
}

func (e EnterpriseDefaultRepositoryPermissionSettingValue) String() string {
	return string(e)
}

func (e *EnterpriseDefaultRepositoryPermissionSettingValue) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EnterpriseDefaultRepositoryPermissionSettingValue(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EnterpriseDefaultRepositoryPermissionSettingValue", str)
	}
	return nil
}

func (e EnterpriseDefaultRepositoryPermissionSettingValue) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The possible values for an enabled/disabled enterprise setting.
type EnterpriseEnabledDisabledSettingValue string

const (
	// The setting is enabled for organizations in the enterprise.
	EnterpriseEnabledDisabledSettingValueEnabled EnterpriseEnabledDisabledSettingValue = "ENABLED"
	// The setting is disabled for organizations in the enterprise.
	EnterpriseEnabledDisabledSettingValueDisabled EnterpriseEnabledDisabledSettingValue = "DISABLED"
	// There is no policy set for organizations in the enterprise.
	EnterpriseEnabledDisabledSettingValueNoPolicy EnterpriseEnabledDisabledSettingValue = "NO_POLICY"
)

var AllEnterpriseEnabledDisabledSettingValue = []EnterpriseEnabledDisabledSettingValue{
	EnterpriseEnabledDisabledSettingValueEnabled,
	EnterpriseEnabledDisabledSettingValueDisabled,
	EnterpriseEnabledDisabledSettingValueNoPolicy,
}

func (e EnterpriseEnabledDisabledSettingValue) IsValid() bool {
	switch e {
	case EnterpriseEnabledDisabledSettingValueEnabled, EnterpriseEnabledDisabledSettingValueDisabled, EnterpriseEnabledDisabledSettingValueNoPolicy:
		return true
	}
	return false
}

func (e EnterpriseEnabledDisabledSettingValue) String() string {
	return string(e)
}

func (e *EnterpriseEnabledDisabledSettingValue) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EnterpriseEnabledDisabledSettingValue(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EnterpriseEnabledDisabledSettingValue", str)
	}
	return nil
}

func (e EnterpriseEnabledDisabledSettingValue) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The possible values for an enabled/no policy enterprise setting.
type EnterpriseEnabledSettingValue string

const (
	// The setting is enabled for organizations in the enterprise.
	EnterpriseEnabledSettingValueEnabled EnterpriseEnabledSettingValue = "ENABLED"
	// There is no policy set for organizations in the enterprise.
	EnterpriseEnabledSettingValueNoPolicy EnterpriseEnabledSettingValue = "NO_POLICY"
)

var AllEnterpriseEnabledSettingValue = []EnterpriseEnabledSettingValue{
	EnterpriseEnabledSettingValueEnabled,
	EnterpriseEnabledSettingValueNoPolicy,
}

func (e EnterpriseEnabledSettingValue) IsValid() bool {
	switch e {
	case EnterpriseEnabledSettingValueEnabled, EnterpriseEnabledSettingValueNoPolicy:
		return true
	}
	return false
}

func (e EnterpriseEnabledSettingValue) String() string {
	return string(e)
}

func (e *EnterpriseEnabledSettingValue) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EnterpriseEnabledSettingValue(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EnterpriseEnabledSettingValue", str)
	}
	return nil
}

func (e EnterpriseEnabledSettingValue) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Properties by which enterprise member connections can be ordered.
type EnterpriseMemberOrderField string

const (
	// Order enterprise members by login
	EnterpriseMemberOrderFieldLogin EnterpriseMemberOrderField = "LOGIN"
	// Order enterprise members by creation time
	EnterpriseMemberOrderFieldCreatedAt EnterpriseMemberOrderField = "CREATED_AT"
)

var AllEnterpriseMemberOrderField = []EnterpriseMemberOrderField{
	EnterpriseMemberOrderFieldLogin,
	EnterpriseMemberOrderFieldCreatedAt,
}

func (e EnterpriseMemberOrderField) IsValid() bool {
	switch e {
	case EnterpriseMemberOrderFieldLogin, EnterpriseMemberOrderFieldCreatedAt:
		return true
	}
	return false
}

func (e EnterpriseMemberOrderField) String() string {
	return string(e)
}

func (e *EnterpriseMemberOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EnterpriseMemberOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EnterpriseMemberOrderField", str)
	}
	return nil
}

func (e EnterpriseMemberOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The possible values for the enterprise members can create repositories setting.
type EnterpriseMembersCanCreateRepositoriesSettingValue string

const (
	// Organization administrators choose whether to allow members to create repositories.
	EnterpriseMembersCanCreateRepositoriesSettingValueNoPolicy EnterpriseMembersCanCreateRepositoriesSettingValue = "NO_POLICY"
	// Members will be able to create public and private repositories.
	EnterpriseMembersCanCreateRepositoriesSettingValueAll EnterpriseMembersCanCreateRepositoriesSettingValue = "ALL"
	// Members will be able to create only public repositories.
	EnterpriseMembersCanCreateRepositoriesSettingValuePublic EnterpriseMembersCanCreateRepositoriesSettingValue = "PUBLIC"
	// Members will be able to create only private repositories.
	EnterpriseMembersCanCreateRepositoriesSettingValuePrivate EnterpriseMembersCanCreateRepositoriesSettingValue = "PRIVATE"
	// Members will not be able to create public or private repositories.
	EnterpriseMembersCanCreateRepositoriesSettingValueDisabled EnterpriseMembersCanCreateRepositoriesSettingValue = "DISABLED"
)

var AllEnterpriseMembersCanCreateRepositoriesSettingValue = []EnterpriseMembersCanCreateRepositoriesSettingValue{
	EnterpriseMembersCanCreateRepositoriesSettingValueNoPolicy,
	EnterpriseMembersCanCreateRepositoriesSettingValueAll,
	EnterpriseMembersCanCreateRepositoriesSettingValuePublic,
	EnterpriseMembersCanCreateRepositoriesSettingValuePrivate,
	EnterpriseMembersCanCreateRepositoriesSettingValueDisabled,
}

func (e EnterpriseMembersCanCreateRepositoriesSettingValue) IsValid() bool {
	switch e {
	case EnterpriseMembersCanCreateRepositoriesSettingValueNoPolicy, EnterpriseMembersCanCreateRepositoriesSettingValueAll, EnterpriseMembersCanCreateRepositoriesSettingValuePublic, EnterpriseMembersCanCreateRepositoriesSettingValuePrivate, EnterpriseMembersCanCreateRepositoriesSettingValueDisabled:
		return true
	}
	return false
}

func (e EnterpriseMembersCanCreateRepositoriesSettingValue) String() string {
	return string(e)
}

func (e *EnterpriseMembersCanCreateRepositoriesSettingValue) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EnterpriseMembersCanCreateRepositoriesSettingValue(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EnterpriseMembersCanCreateRepositoriesSettingValue", str)
	}
	return nil
}

func (e EnterpriseMembersCanCreateRepositoriesSettingValue) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The possible values for the members can make purchases setting.
type EnterpriseMembersCanMakePurchasesSettingValue string

const (
	// The setting is enabled for organizations in the enterprise.
	EnterpriseMembersCanMakePurchasesSettingValueEnabled EnterpriseMembersCanMakePurchasesSettingValue = "ENABLED"
	// The setting is disabled for organizations in the enterprise.
	EnterpriseMembersCanMakePurchasesSettingValueDisabled EnterpriseMembersCanMakePurchasesSettingValue = "DISABLED"
)

var AllEnterpriseMembersCanMakePurchasesSettingValue = []EnterpriseMembersCanMakePurchasesSettingValue{
	EnterpriseMembersCanMakePurchasesSettingValueEnabled,
	EnterpriseMembersCanMakePurchasesSettingValueDisabled,
}

func (e EnterpriseMembersCanMakePurchasesSettingValue) IsValid() bool {
	switch e {
	case EnterpriseMembersCanMakePurchasesSettingValueEnabled, EnterpriseMembersCanMakePurchasesSettingValueDisabled:
		return true
	}
	return false
}

func (e EnterpriseMembersCanMakePurchasesSettingValue) String() string {
	return string(e)
}

func (e *EnterpriseMembersCanMakePurchasesSettingValue) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EnterpriseMembersCanMakePurchasesSettingValue(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EnterpriseMembersCanMakePurchasesSettingValue", str)
	}
	return nil
}

func (e EnterpriseMembersCanMakePurchasesSettingValue) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Properties by which Enterprise Server installation connections can be ordered.
type EnterpriseServerInstallationOrderField string

const (
	// Order Enterprise Server installations by host name
	EnterpriseServerInstallationOrderFieldHostName EnterpriseServerInstallationOrderField = "HOST_NAME"
	// Order Enterprise Server installations by customer name
	EnterpriseServerInstallationOrderFieldCustomerName EnterpriseServerInstallationOrderField = "CUSTOMER_NAME"
	// Order Enterprise Server installations by creation time
	EnterpriseServerInstallationOrderFieldCreatedAt EnterpriseServerInstallationOrderField = "CREATED_AT"
)

var AllEnterpriseServerInstallationOrderField = []EnterpriseServerInstallationOrderField{
	EnterpriseServerInstallationOrderFieldHostName,
	EnterpriseServerInstallationOrderFieldCustomerName,
	EnterpriseServerInstallationOrderFieldCreatedAt,
}

func (e EnterpriseServerInstallationOrderField) IsValid() bool {
	switch e {
	case EnterpriseServerInstallationOrderFieldHostName, EnterpriseServerInstallationOrderFieldCustomerName, EnterpriseServerInstallationOrderFieldCreatedAt:
		return true
	}
	return false
}

func (e EnterpriseServerInstallationOrderField) String() string {
	return string(e)
}

func (e *EnterpriseServerInstallationOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EnterpriseServerInstallationOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EnterpriseServerInstallationOrderField", str)
	}
	return nil
}

func (e EnterpriseServerInstallationOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Properties by which Enterprise Server user account email connections can be ordered.
type EnterpriseServerUserAccountEmailOrderField string

const (
	// Order emails by email
	EnterpriseServerUserAccountEmailOrderFieldEmail EnterpriseServerUserAccountEmailOrderField = "EMAIL"
)

var AllEnterpriseServerUserAccountEmailOrderField = []EnterpriseServerUserAccountEmailOrderField{
	EnterpriseServerUserAccountEmailOrderFieldEmail,
}

func (e EnterpriseServerUserAccountEmailOrderField) IsValid() bool {
	switch e {
	case EnterpriseServerUserAccountEmailOrderFieldEmail:
		return true
	}
	return false
}

func (e EnterpriseServerUserAccountEmailOrderField) String() string {
	return string(e)
}

func (e *EnterpriseServerUserAccountEmailOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EnterpriseServerUserAccountEmailOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EnterpriseServerUserAccountEmailOrderField", str)
	}
	return nil
}

func (e EnterpriseServerUserAccountEmailOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Properties by which Enterprise Server user account connections can be ordered.
type EnterpriseServerUserAccountOrderField string

const (
	// Order user accounts by login
	EnterpriseServerUserAccountOrderFieldLogin EnterpriseServerUserAccountOrderField = "LOGIN"
	// Order user accounts by creation time on the Enterprise Server installation
	EnterpriseServerUserAccountOrderFieldRemoteCreatedAt EnterpriseServerUserAccountOrderField = "REMOTE_CREATED_AT"
)

var AllEnterpriseServerUserAccountOrderField = []EnterpriseServerUserAccountOrderField{
	EnterpriseServerUserAccountOrderFieldLogin,
	EnterpriseServerUserAccountOrderFieldRemoteCreatedAt,
}

func (e EnterpriseServerUserAccountOrderField) IsValid() bool {
	switch e {
	case EnterpriseServerUserAccountOrderFieldLogin, EnterpriseServerUserAccountOrderFieldRemoteCreatedAt:
		return true
	}
	return false
}

func (e EnterpriseServerUserAccountOrderField) String() string {
	return string(e)
}

func (e *EnterpriseServerUserAccountOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EnterpriseServerUserAccountOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EnterpriseServerUserAccountOrderField", str)
	}
	return nil
}

func (e EnterpriseServerUserAccountOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Properties by which Enterprise Server user accounts upload connections can be ordered.
type EnterpriseServerUserAccountsUploadOrderField string

const (
	// Order user accounts uploads by creation time
	EnterpriseServerUserAccountsUploadOrderFieldCreatedAt EnterpriseServerUserAccountsUploadOrderField = "CREATED_AT"
)

var AllEnterpriseServerUserAccountsUploadOrderField = []EnterpriseServerUserAccountsUploadOrderField{
	EnterpriseServerUserAccountsUploadOrderFieldCreatedAt,
}

func (e EnterpriseServerUserAccountsUploadOrderField) IsValid() bool {
	switch e {
	case EnterpriseServerUserAccountsUploadOrderFieldCreatedAt:
		return true
	}
	return false
}

func (e EnterpriseServerUserAccountsUploadOrderField) String() string {
	return string(e)
}

func (e *EnterpriseServerUserAccountsUploadOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EnterpriseServerUserAccountsUploadOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EnterpriseServerUserAccountsUploadOrderField", str)
	}
	return nil
}

func (e EnterpriseServerUserAccountsUploadOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Synchronization state of the Enterprise Server user accounts upload
type EnterpriseServerUserAccountsUploadSyncState string

const (
	// The synchronization of the upload is pending.
	EnterpriseServerUserAccountsUploadSyncStatePending EnterpriseServerUserAccountsUploadSyncState = "PENDING"
	// The synchronization of the upload succeeded.
	EnterpriseServerUserAccountsUploadSyncStateSuccess EnterpriseServerUserAccountsUploadSyncState = "SUCCESS"
	// The synchronization of the upload failed.
	EnterpriseServerUserAccountsUploadSyncStateFailure EnterpriseServerUserAccountsUploadSyncState = "FAILURE"
)

var AllEnterpriseServerUserAccountsUploadSyncState = []EnterpriseServerUserAccountsUploadSyncState{
	EnterpriseServerUserAccountsUploadSyncStatePending,
	EnterpriseServerUserAccountsUploadSyncStateSuccess,
	EnterpriseServerUserAccountsUploadSyncStateFailure,
}

func (e EnterpriseServerUserAccountsUploadSyncState) IsValid() bool {
	switch e {
	case EnterpriseServerUserAccountsUploadSyncStatePending, EnterpriseServerUserAccountsUploadSyncStateSuccess, EnterpriseServerUserAccountsUploadSyncStateFailure:
		return true
	}
	return false
}

func (e EnterpriseServerUserAccountsUploadSyncState) String() string {
	return string(e)
}

func (e *EnterpriseServerUserAccountsUploadSyncState) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EnterpriseServerUserAccountsUploadSyncState(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EnterpriseServerUserAccountsUploadSyncState", str)
	}
	return nil
}

func (e EnterpriseServerUserAccountsUploadSyncState) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The possible roles for enterprise membership.
type EnterpriseUserAccountMembershipRole string

const (
	// The user is a member of the enterprise membership.
	EnterpriseUserAccountMembershipRoleMember EnterpriseUserAccountMembershipRole = "MEMBER"
	// The user is an owner of the enterprise membership.
	EnterpriseUserAccountMembershipRoleOwner EnterpriseUserAccountMembershipRole = "OWNER"
)

var AllEnterpriseUserAccountMembershipRole = []EnterpriseUserAccountMembershipRole{
	EnterpriseUserAccountMembershipRoleMember,
	EnterpriseUserAccountMembershipRoleOwner,
}

func (e EnterpriseUserAccountMembershipRole) IsValid() bool {
	switch e {
	case EnterpriseUserAccountMembershipRoleMember, EnterpriseUserAccountMembershipRoleOwner:
		return true
	}
	return false
}

func (e EnterpriseUserAccountMembershipRole) String() string {
	return string(e)
}

func (e *EnterpriseUserAccountMembershipRole) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EnterpriseUserAccountMembershipRole(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EnterpriseUserAccountMembershipRole", str)
	}
	return nil
}

func (e EnterpriseUserAccountMembershipRole) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The possible GitHub Enterprise deployments where this user can exist.
type EnterpriseUserDeployment string

const (
	// The user is part of a GitHub Enterprise Cloud deployment.
	EnterpriseUserDeploymentCloud EnterpriseUserDeployment = "CLOUD"
	// The user is part of a GitHub Enterprise Server deployment.
	EnterpriseUserDeploymentServer EnterpriseUserDeployment = "SERVER"
)

var AllEnterpriseUserDeployment = []EnterpriseUserDeployment{
	EnterpriseUserDeploymentCloud,
	EnterpriseUserDeploymentServer,
}

func (e EnterpriseUserDeployment) IsValid() bool {
	switch e {
	case EnterpriseUserDeploymentCloud, EnterpriseUserDeploymentServer:
		return true
	}
	return false
}

func (e EnterpriseUserDeployment) String() string {
	return string(e)
}

func (e *EnterpriseUserDeployment) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EnterpriseUserDeployment(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EnterpriseUserDeployment", str)
	}
	return nil
}

func (e EnterpriseUserDeployment) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The possible viewed states of a file .
type FileViewedState string

const (
	// The file has new changes since last viewed.
	FileViewedStateDismissed FileViewedState = "DISMISSED"
	// The file has been marked as viewed.
	FileViewedStateViewed FileViewedState = "VIEWED"
	// The file has not been marked as viewed.
	FileViewedStateUnviewed FileViewedState = "UNVIEWED"
)

var AllFileViewedState = []FileViewedState{
	FileViewedStateDismissed,
	FileViewedStateViewed,
	FileViewedStateUnviewed,
}

func (e FileViewedState) IsValid() bool {
	switch e {
	case FileViewedStateDismissed, FileViewedStateViewed, FileViewedStateUnviewed:
		return true
	}
	return false
}

func (e FileViewedState) String() string {
	return string(e)
}

func (e *FileViewedState) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = FileViewedState(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid FileViewedState", str)
	}
	return nil
}

func (e FileViewedState) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The possible funding platforms for repository funding links.
type FundingPlatform string

const (
	// GitHub funding platform.
	FundingPlatformGithub FundingPlatform = "GITHUB"
	// Patreon funding platform.
	FundingPlatformPatreon FundingPlatform = "PATREON"
	// Open Collective funding platform.
	FundingPlatformOpenCollective FundingPlatform = "OPEN_COLLECTIVE"
	// Ko-fi funding platform.
	FundingPlatformKoFi FundingPlatform = "KO_FI"
	// Tidelift funding platform.
	FundingPlatformTidelift FundingPlatform = "TIDELIFT"
	// Community Bridge funding platform.
	FundingPlatformCommunityBridge FundingPlatform = "COMMUNITY_BRIDGE"
	// Liberapay funding platform.
	FundingPlatformLiberapay FundingPlatform = "LIBERAPAY"
	// IssueHunt funding platform.
	FundingPlatformIssuehunt FundingPlatform = "ISSUEHUNT"
	// Otechie funding platform.
	FundingPlatformOtechie FundingPlatform = "OTECHIE"
	// Custom funding platform.
	FundingPlatformCustom FundingPlatform = "CUSTOM"
)

var AllFundingPlatform = []FundingPlatform{
	FundingPlatformGithub,
	FundingPlatformPatreon,
	FundingPlatformOpenCollective,
	FundingPlatformKoFi,
	FundingPlatformTidelift,
	FundingPlatformCommunityBridge,
	FundingPlatformLiberapay,
	FundingPlatformIssuehunt,
	FundingPlatformOtechie,
	FundingPlatformCustom,
}

func (e FundingPlatform) IsValid() bool {
	switch e {
	case FundingPlatformGithub, FundingPlatformPatreon, FundingPlatformOpenCollective, FundingPlatformKoFi, FundingPlatformTidelift, FundingPlatformCommunityBridge, FundingPlatformLiberapay, FundingPlatformIssuehunt, FundingPlatformOtechie, FundingPlatformCustom:
		return true
	}
	return false
}

func (e FundingPlatform) String() string {
	return string(e)
}

func (e *FundingPlatform) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = FundingPlatform(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid FundingPlatform", str)
	}
	return nil
}

func (e FundingPlatform) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Properties by which gist connections can be ordered.
type GistOrderField string

const (
	// Order gists by creation time
	GistOrderFieldCreatedAt GistOrderField = "CREATED_AT"
	// Order gists by update time
	GistOrderFieldUpdatedAt GistOrderField = "UPDATED_AT"
	// Order gists by push time
	GistOrderFieldPushedAt GistOrderField = "PUSHED_AT"
)

var AllGistOrderField = []GistOrderField{
	GistOrderFieldCreatedAt,
	GistOrderFieldUpdatedAt,
	GistOrderFieldPushedAt,
}

func (e GistOrderField) IsValid() bool {
	switch e {
	case GistOrderFieldCreatedAt, GistOrderFieldUpdatedAt, GistOrderFieldPushedAt:
		return true
	}
	return false
}

func (e GistOrderField) String() string {
	return string(e)
}

func (e *GistOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = GistOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid GistOrderField", str)
	}
	return nil
}

func (e GistOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The privacy of a Gist
type GistPrivacy string

const (
	// Public
	GistPrivacyPublic GistPrivacy = "PUBLIC"
	// Secret
	GistPrivacySecret GistPrivacy = "SECRET"
	// Gists that are public and secret
	GistPrivacyAll GistPrivacy = "ALL"
)

var AllGistPrivacy = []GistPrivacy{
	GistPrivacyPublic,
	GistPrivacySecret,
	GistPrivacyAll,
}

func (e GistPrivacy) IsValid() bool {
	switch e {
	case GistPrivacyPublic, GistPrivacySecret, GistPrivacyAll:
		return true
	}
	return false
}

func (e GistPrivacy) String() string {
	return string(e)
}

func (e *GistPrivacy) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = GistPrivacy(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid GistPrivacy", str)
	}
	return nil
}

func (e GistPrivacy) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The state of a Git signature.
type GitSignatureState string

const (
	// Valid signature and verified by GitHub
	GitSignatureStateValid GitSignatureState = "VALID"
	// Invalid signature
	GitSignatureStateInvalid GitSignatureState = "INVALID"
	// Malformed signature
	GitSignatureStateMalformedSig GitSignatureState = "MALFORMED_SIG"
	// Key used for signing not known to GitHub
	GitSignatureStateUnknownKey GitSignatureState = "UNKNOWN_KEY"
	// Invalid email used for signing
	GitSignatureStateBadEmail GitSignatureState = "BAD_EMAIL"
	// Email used for signing unverified on GitHub
	GitSignatureStateUnverifiedEmail GitSignatureState = "UNVERIFIED_EMAIL"
	// Email used for signing not known to GitHub
	GitSignatureStateNoUser GitSignatureState = "NO_USER"
	// Unknown signature type
	GitSignatureStateUnknownSigType GitSignatureState = "UNKNOWN_SIG_TYPE"
	// Unsigned
	GitSignatureStateUnsigned GitSignatureState = "UNSIGNED"
	// Internal error - the GPG verification service is unavailable at the moment
	GitSignatureStateGpgverifyUnavailable GitSignatureState = "GPGVERIFY_UNAVAILABLE"
	// Internal error - the GPG verification service misbehaved
	GitSignatureStateGpgverifyError GitSignatureState = "GPGVERIFY_ERROR"
	// The usage flags for the key that signed this don't allow signing
	GitSignatureStateNotSigningKey GitSignatureState = "NOT_SIGNING_KEY"
	// Signing key expired
	GitSignatureStateExpiredKey GitSignatureState = "EXPIRED_KEY"
	// Valid signature, pending certificate revocation checking
	GitSignatureStateOcspPending GitSignatureState = "OCSP_PENDING"
	// Valid siganture, though certificate revocation check failed
	GitSignatureStateOcspError GitSignatureState = "OCSP_ERROR"
	// The signing certificate or its chain could not be verified
	GitSignatureStateBadCert GitSignatureState = "BAD_CERT"
	// One or more certificates in chain has been revoked
	GitSignatureStateOcspRevoked GitSignatureState = "OCSP_REVOKED"
)

var AllGitSignatureState = []GitSignatureState{
	GitSignatureStateValid,
	GitSignatureStateInvalid,
	GitSignatureStateMalformedSig,
	GitSignatureStateUnknownKey,
	GitSignatureStateBadEmail,
	GitSignatureStateUnverifiedEmail,
	GitSignatureStateNoUser,
	GitSignatureStateUnknownSigType,
	GitSignatureStateUnsigned,
	GitSignatureStateGpgverifyUnavailable,
	GitSignatureStateGpgverifyError,
	GitSignatureStateNotSigningKey,
	GitSignatureStateExpiredKey,
	GitSignatureStateOcspPending,
	GitSignatureStateOcspError,
	GitSignatureStateBadCert,
	GitSignatureStateOcspRevoked,
}

func (e GitSignatureState) IsValid() bool {
	switch e {
	case GitSignatureStateValid, GitSignatureStateInvalid, GitSignatureStateMalformedSig, GitSignatureStateUnknownKey, GitSignatureStateBadEmail, GitSignatureStateUnverifiedEmail, GitSignatureStateNoUser, GitSignatureStateUnknownSigType, GitSignatureStateUnsigned, GitSignatureStateGpgverifyUnavailable, GitSignatureStateGpgverifyError, GitSignatureStateNotSigningKey, GitSignatureStateExpiredKey, GitSignatureStateOcspPending, GitSignatureStateOcspError, GitSignatureStateBadCert, GitSignatureStateOcspRevoked:
		return true
	}
	return false
}

func (e GitSignatureState) String() string {
	return string(e)
}

func (e *GitSignatureState) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = GitSignatureState(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid GitSignatureState", str)
	}
	return nil
}

func (e GitSignatureState) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The possible states in which authentication can be configured with an identity provider.
type IdentityProviderConfigurationState string

const (
	// Authentication with an identity provider is configured and enforced.
	IdentityProviderConfigurationStateEnforced IdentityProviderConfigurationState = "ENFORCED"
	// Authentication with an identity provider is configured but not enforced.
	IdentityProviderConfigurationStateConfigured IdentityProviderConfigurationState = "CONFIGURED"
	// Authentication with an identity provider is not configured.
	IdentityProviderConfigurationStateUnconfigured IdentityProviderConfigurationState = "UNCONFIGURED"
)

var AllIdentityProviderConfigurationState = []IdentityProviderConfigurationState{
	IdentityProviderConfigurationStateEnforced,
	IdentityProviderConfigurationStateConfigured,
	IdentityProviderConfigurationStateUnconfigured,
}

func (e IdentityProviderConfigurationState) IsValid() bool {
	switch e {
	case IdentityProviderConfigurationStateEnforced, IdentityProviderConfigurationStateConfigured, IdentityProviderConfigurationStateUnconfigured:
		return true
	}
	return false
}

func (e IdentityProviderConfigurationState) String() string {
	return string(e)
}

func (e *IdentityProviderConfigurationState) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = IdentityProviderConfigurationState(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid IdentityProviderConfigurationState", str)
	}
	return nil
}

func (e IdentityProviderConfigurationState) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The possible values for the IP allow list enabled setting.
type IPAllowListEnabledSettingValue string

const (
	// The setting is enabled for the owner.
	IPAllowListEnabledSettingValueEnabled IPAllowListEnabledSettingValue = "ENABLED"
	// The setting is disabled for the owner.
	IPAllowListEnabledSettingValueDisabled IPAllowListEnabledSettingValue = "DISABLED"
)

var AllIPAllowListEnabledSettingValue = []IPAllowListEnabledSettingValue{
	IPAllowListEnabledSettingValueEnabled,
	IPAllowListEnabledSettingValueDisabled,
}

func (e IPAllowListEnabledSettingValue) IsValid() bool {
	switch e {
	case IPAllowListEnabledSettingValueEnabled, IPAllowListEnabledSettingValueDisabled:
		return true
	}
	return false
}

func (e IPAllowListEnabledSettingValue) String() string {
	return string(e)
}

func (e *IPAllowListEnabledSettingValue) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = IPAllowListEnabledSettingValue(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid IpAllowListEnabledSettingValue", str)
	}
	return nil
}

func (e IPAllowListEnabledSettingValue) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Properties by which IP allow list entry connections can be ordered.
type IPAllowListEntryOrderField string

const (
	// Order IP allow list entries by creation time.
	IPAllowListEntryOrderFieldCreatedAt IPAllowListEntryOrderField = "CREATED_AT"
	// Order IP allow list entries by the allow list value.
	IPAllowListEntryOrderFieldAllowListValue IPAllowListEntryOrderField = "ALLOW_LIST_VALUE"
)

var AllIPAllowListEntryOrderField = []IPAllowListEntryOrderField{
	IPAllowListEntryOrderFieldCreatedAt,
	IPAllowListEntryOrderFieldAllowListValue,
}

func (e IPAllowListEntryOrderField) IsValid() bool {
	switch e {
	case IPAllowListEntryOrderFieldCreatedAt, IPAllowListEntryOrderFieldAllowListValue:
		return true
	}
	return false
}

func (e IPAllowListEntryOrderField) String() string {
	return string(e)
}

func (e *IPAllowListEntryOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = IPAllowListEntryOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid IpAllowListEntryOrderField", str)
	}
	return nil
}

func (e IPAllowListEntryOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Properties by which issue connections can be ordered.
type IssueOrderField string

const (
	// Order issues by creation time
	IssueOrderFieldCreatedAt IssueOrderField = "CREATED_AT"
	// Order issues by update time
	IssueOrderFieldUpdatedAt IssueOrderField = "UPDATED_AT"
	// Order issues by comment count
	IssueOrderFieldComments IssueOrderField = "COMMENTS"
)

var AllIssueOrderField = []IssueOrderField{
	IssueOrderFieldCreatedAt,
	IssueOrderFieldUpdatedAt,
	IssueOrderFieldComments,
}

func (e IssueOrderField) IsValid() bool {
	switch e {
	case IssueOrderFieldCreatedAt, IssueOrderFieldUpdatedAt, IssueOrderFieldComments:
		return true
	}
	return false
}

func (e IssueOrderField) String() string {
	return string(e)
}

func (e *IssueOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = IssueOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid IssueOrderField", str)
	}
	return nil
}

func (e IssueOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The possible states of an issue.
type IssueState string

const (
	// An issue that is still open
	IssueStateOpen IssueState = "OPEN"
	// An issue that has been closed
	IssueStateClosed IssueState = "CLOSED"
)

var AllIssueState = []IssueState{
	IssueStateOpen,
	IssueStateClosed,
}

func (e IssueState) IsValid() bool {
	switch e {
	case IssueStateOpen, IssueStateClosed:
		return true
	}
	return false
}

func (e IssueState) String() string {
	return string(e)
}

func (e *IssueState) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = IssueState(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid IssueState", str)
	}
	return nil
}

func (e IssueState) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The possible item types found in a timeline.
type IssueTimelineItemsItemType string

const (
	// Represents a comment on an Issue.
	IssueTimelineItemsItemTypeIssueComment IssueTimelineItemsItemType = "ISSUE_COMMENT"
	// Represents a mention made by one issue or pull request to another.
	IssueTimelineItemsItemTypeCrossReferencedEvent IssueTimelineItemsItemType = "CROSS_REFERENCED_EVENT"
	// Represents a 'added_to_project' event on a given issue or pull request.
	IssueTimelineItemsItemTypeAddedToProjectEvent IssueTimelineItemsItemType = "ADDED_TO_PROJECT_EVENT"
	// Represents an 'assigned' event on any assignable object.
	IssueTimelineItemsItemTypeAssignedEvent IssueTimelineItemsItemType = "ASSIGNED_EVENT"
	// Represents a 'closed' event on any `Closable`.
	IssueTimelineItemsItemTypeClosedEvent IssueTimelineItemsItemType = "CLOSED_EVENT"
	// Represents a 'comment_deleted' event on a given issue or pull request.
	IssueTimelineItemsItemTypeCommentDeletedEvent IssueTimelineItemsItemType = "COMMENT_DELETED_EVENT"
	// Represents a 'connected' event on a given issue or pull request.
	IssueTimelineItemsItemTypeConnectedEvent IssueTimelineItemsItemType = "CONNECTED_EVENT"
	// Represents a 'converted_note_to_issue' event on a given issue or pull request.
	IssueTimelineItemsItemTypeConvertedNoteToIssueEvent IssueTimelineItemsItemType = "CONVERTED_NOTE_TO_ISSUE_EVENT"
	// Represents a 'demilestoned' event on a given issue or pull request.
	IssueTimelineItemsItemTypeDemilestonedEvent IssueTimelineItemsItemType = "DEMILESTONED_EVENT"
	// Represents a 'disconnected' event on a given issue or pull request.
	IssueTimelineItemsItemTypeDisconnectedEvent IssueTimelineItemsItemType = "DISCONNECTED_EVENT"
	// Represents a 'labeled' event on a given issue or pull request.
	IssueTimelineItemsItemTypeLabeledEvent IssueTimelineItemsItemType = "LABELED_EVENT"
	// Represents a 'locked' event on a given issue or pull request.
	IssueTimelineItemsItemTypeLockedEvent IssueTimelineItemsItemType = "LOCKED_EVENT"
	// Represents a 'marked_as_duplicate' event on a given issue or pull request.
	IssueTimelineItemsItemTypeMarkedAsDuplicateEvent IssueTimelineItemsItemType = "MARKED_AS_DUPLICATE_EVENT"
	// Represents a 'mentioned' event on a given issue or pull request.
	IssueTimelineItemsItemTypeMentionedEvent IssueTimelineItemsItemType = "MENTIONED_EVENT"
	// Represents a 'milestoned' event on a given issue or pull request.
	IssueTimelineItemsItemTypeMilestonedEvent IssueTimelineItemsItemType = "MILESTONED_EVENT"
	// Represents a 'moved_columns_in_project' event on a given issue or pull request.
	IssueTimelineItemsItemTypeMovedColumnsInProjectEvent IssueTimelineItemsItemType = "MOVED_COLUMNS_IN_PROJECT_EVENT"
	// Represents a 'pinned' event on a given issue or pull request.
	IssueTimelineItemsItemTypePinnedEvent IssueTimelineItemsItemType = "PINNED_EVENT"
	// Represents a 'referenced' event on a given `ReferencedSubject`.
	IssueTimelineItemsItemTypeReferencedEvent IssueTimelineItemsItemType = "REFERENCED_EVENT"
	// Represents a 'removed_from_project' event on a given issue or pull request.
	IssueTimelineItemsItemTypeRemovedFromProjectEvent IssueTimelineItemsItemType = "REMOVED_FROM_PROJECT_EVENT"
	// Represents a 'renamed' event on a given issue or pull request
	IssueTimelineItemsItemTypeRenamedTitleEvent IssueTimelineItemsItemType = "RENAMED_TITLE_EVENT"
	// Represents a 'reopened' event on any `Closable`.
	IssueTimelineItemsItemTypeReopenedEvent IssueTimelineItemsItemType = "REOPENED_EVENT"
	// Represents a 'subscribed' event on a given `Subscribable`.
	IssueTimelineItemsItemTypeSubscribedEvent IssueTimelineItemsItemType = "SUBSCRIBED_EVENT"
	// Represents a 'transferred' event on a given issue or pull request.
	IssueTimelineItemsItemTypeTransferredEvent IssueTimelineItemsItemType = "TRANSFERRED_EVENT"
	// Represents an 'unassigned' event on any assignable object.
	IssueTimelineItemsItemTypeUnassignedEvent IssueTimelineItemsItemType = "UNASSIGNED_EVENT"
	// Represents an 'unlabeled' event on a given issue or pull request.
	IssueTimelineItemsItemTypeUnlabeledEvent IssueTimelineItemsItemType = "UNLABELED_EVENT"
	// Represents an 'unlocked' event on a given issue or pull request.
	IssueTimelineItemsItemTypeUnlockedEvent IssueTimelineItemsItemType = "UNLOCKED_EVENT"
	// Represents a 'user_blocked' event on a given user.
	IssueTimelineItemsItemTypeUserBlockedEvent IssueTimelineItemsItemType = "USER_BLOCKED_EVENT"
	// Represents an 'unmarked_as_duplicate' event on a given issue or pull request.
	IssueTimelineItemsItemTypeUnmarkedAsDuplicateEvent IssueTimelineItemsItemType = "UNMARKED_AS_DUPLICATE_EVENT"
	// Represents an 'unpinned' event on a given issue or pull request.
	IssueTimelineItemsItemTypeUnpinnedEvent IssueTimelineItemsItemType = "UNPINNED_EVENT"
	// Represents an 'unsubscribed' event on a given `Subscribable`.
	IssueTimelineItemsItemTypeUnsubscribedEvent IssueTimelineItemsItemType = "UNSUBSCRIBED_EVENT"
)

var AllIssueTimelineItemsItemType = []IssueTimelineItemsItemType{
	IssueTimelineItemsItemTypeIssueComment,
	IssueTimelineItemsItemTypeCrossReferencedEvent,
	IssueTimelineItemsItemTypeAddedToProjectEvent,
	IssueTimelineItemsItemTypeAssignedEvent,
	IssueTimelineItemsItemTypeClosedEvent,
	IssueTimelineItemsItemTypeCommentDeletedEvent,
	IssueTimelineItemsItemTypeConnectedEvent,
	IssueTimelineItemsItemTypeConvertedNoteToIssueEvent,
	IssueTimelineItemsItemTypeDemilestonedEvent,
	IssueTimelineItemsItemTypeDisconnectedEvent,
	IssueTimelineItemsItemTypeLabeledEvent,
	IssueTimelineItemsItemTypeLockedEvent,
	IssueTimelineItemsItemTypeMarkedAsDuplicateEvent,
	IssueTimelineItemsItemTypeMentionedEvent,
	IssueTimelineItemsItemTypeMilestonedEvent,
	IssueTimelineItemsItemTypeMovedColumnsInProjectEvent,
	IssueTimelineItemsItemTypePinnedEvent,
	IssueTimelineItemsItemTypeReferencedEvent,
	IssueTimelineItemsItemTypeRemovedFromProjectEvent,
	IssueTimelineItemsItemTypeRenamedTitleEvent,
	IssueTimelineItemsItemTypeReopenedEvent,
	IssueTimelineItemsItemTypeSubscribedEvent,
	IssueTimelineItemsItemTypeTransferredEvent,
	IssueTimelineItemsItemTypeUnassignedEvent,
	IssueTimelineItemsItemTypeUnlabeledEvent,
	IssueTimelineItemsItemTypeUnlockedEvent,
	IssueTimelineItemsItemTypeUserBlockedEvent,
	IssueTimelineItemsItemTypeUnmarkedAsDuplicateEvent,
	IssueTimelineItemsItemTypeUnpinnedEvent,
	IssueTimelineItemsItemTypeUnsubscribedEvent,
}

func (e IssueTimelineItemsItemType) IsValid() bool {
	switch e {
	case IssueTimelineItemsItemTypeIssueComment, IssueTimelineItemsItemTypeCrossReferencedEvent, IssueTimelineItemsItemTypeAddedToProjectEvent, IssueTimelineItemsItemTypeAssignedEvent, IssueTimelineItemsItemTypeClosedEvent, IssueTimelineItemsItemTypeCommentDeletedEvent, IssueTimelineItemsItemTypeConnectedEvent, IssueTimelineItemsItemTypeConvertedNoteToIssueEvent, IssueTimelineItemsItemTypeDemilestonedEvent, IssueTimelineItemsItemTypeDisconnectedEvent, IssueTimelineItemsItemTypeLabeledEvent, IssueTimelineItemsItemTypeLockedEvent, IssueTimelineItemsItemTypeMarkedAsDuplicateEvent, IssueTimelineItemsItemTypeMentionedEvent, IssueTimelineItemsItemTypeMilestonedEvent, IssueTimelineItemsItemTypeMovedColumnsInProjectEvent, IssueTimelineItemsItemTypePinnedEvent, IssueTimelineItemsItemTypeReferencedEvent, IssueTimelineItemsItemTypeRemovedFromProjectEvent, IssueTimelineItemsItemTypeRenamedTitleEvent, IssueTimelineItemsItemTypeReopenedEvent, IssueTimelineItemsItemTypeSubscribedEvent, IssueTimelineItemsItemTypeTransferredEvent, IssueTimelineItemsItemTypeUnassignedEvent, IssueTimelineItemsItemTypeUnlabeledEvent, IssueTimelineItemsItemTypeUnlockedEvent, IssueTimelineItemsItemTypeUserBlockedEvent, IssueTimelineItemsItemTypeUnmarkedAsDuplicateEvent, IssueTimelineItemsItemTypeUnpinnedEvent, IssueTimelineItemsItemTypeUnsubscribedEvent:
		return true
	}
	return false
}

func (e IssueTimelineItemsItemType) String() string {
	return string(e)
}

func (e *IssueTimelineItemsItemType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = IssueTimelineItemsItemType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid IssueTimelineItemsItemType", str)
	}
	return nil
}

func (e IssueTimelineItemsItemType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Properties by which label connections can be ordered.
type LabelOrderField string

const (
	// Order labels by name
	LabelOrderFieldName LabelOrderField = "NAME"
	// Order labels by creation time
	LabelOrderFieldCreatedAt LabelOrderField = "CREATED_AT"
)

var AllLabelOrderField = []LabelOrderField{
	LabelOrderFieldName,
	LabelOrderFieldCreatedAt,
}

func (e LabelOrderField) IsValid() bool {
	switch e {
	case LabelOrderFieldName, LabelOrderFieldCreatedAt:
		return true
	}
	return false
}

func (e LabelOrderField) String() string {
	return string(e)
}

func (e *LabelOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = LabelOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid LabelOrderField", str)
	}
	return nil
}

func (e LabelOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Properties by which language connections can be ordered.
type LanguageOrderField string

const (
	// Order languages by the size of all files containing the language
	LanguageOrderFieldSize LanguageOrderField = "SIZE"
)

var AllLanguageOrderField = []LanguageOrderField{
	LanguageOrderFieldSize,
}

func (e LanguageOrderField) IsValid() bool {
	switch e {
	case LanguageOrderFieldSize:
		return true
	}
	return false
}

func (e LanguageOrderField) String() string {
	return string(e)
}

func (e *LanguageOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = LanguageOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid LanguageOrderField", str)
	}
	return nil
}

func (e LanguageOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The possible reasons that an issue or pull request was locked.
type LockReason string

const (
	// The issue or pull request was locked because the conversation was off-topic.
	LockReasonOffTopic LockReason = "OFF_TOPIC"
	// The issue or pull request was locked because the conversation was too heated.
	LockReasonTooHeated LockReason = "TOO_HEATED"
	// The issue or pull request was locked because the conversation was resolved.
	LockReasonResolved LockReason = "RESOLVED"
	// The issue or pull request was locked because the conversation was spam.
	LockReasonSpam LockReason = "SPAM"
)

var AllLockReason = []LockReason{
	LockReasonOffTopic,
	LockReasonTooHeated,
	LockReasonResolved,
	LockReasonSpam,
}

func (e LockReason) IsValid() bool {
	switch e {
	case LockReasonOffTopic, LockReasonTooHeated, LockReasonResolved, LockReasonSpam:
		return true
	}
	return false
}

func (e LockReason) String() string {
	return string(e)
}

func (e *LockReason) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = LockReason(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid LockReason", str)
	}
	return nil
}

func (e LockReason) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Whether or not a PullRequest can be merged.
type MergeableState string

const (
	// The pull request can be merged.
	MergeableStateMergeable MergeableState = "MERGEABLE"
	// The pull request cannot be merged due to merge conflicts.
	MergeableStateConflicting MergeableState = "CONFLICTING"
	// The mergeability of the pull request is still being calculated.
	MergeableStateUnknown MergeableState = "UNKNOWN"
)

var AllMergeableState = []MergeableState{
	MergeableStateMergeable,
	MergeableStateConflicting,
	MergeableStateUnknown,
}

func (e MergeableState) IsValid() bool {
	switch e {
	case MergeableStateMergeable, MergeableStateConflicting, MergeableStateUnknown:
		return true
	}
	return false
}

func (e MergeableState) String() string {
	return string(e)
}

func (e *MergeableState) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = MergeableState(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid MergeableState", str)
	}
	return nil
}

func (e MergeableState) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Properties by which milestone connections can be ordered.
type MilestoneOrderField string

const (
	// Order milestones by when they are due.
	MilestoneOrderFieldDueDate MilestoneOrderField = "DUE_DATE"
	// Order milestones by when they were created.
	MilestoneOrderFieldCreatedAt MilestoneOrderField = "CREATED_AT"
	// Order milestones by when they were last updated.
	MilestoneOrderFieldUpdatedAt MilestoneOrderField = "UPDATED_AT"
	// Order milestones by their number.
	MilestoneOrderFieldNumber MilestoneOrderField = "NUMBER"
)

var AllMilestoneOrderField = []MilestoneOrderField{
	MilestoneOrderFieldDueDate,
	MilestoneOrderFieldCreatedAt,
	MilestoneOrderFieldUpdatedAt,
	MilestoneOrderFieldNumber,
}

func (e MilestoneOrderField) IsValid() bool {
	switch e {
	case MilestoneOrderFieldDueDate, MilestoneOrderFieldCreatedAt, MilestoneOrderFieldUpdatedAt, MilestoneOrderFieldNumber:
		return true
	}
	return false
}

func (e MilestoneOrderField) String() string {
	return string(e)
}

func (e *MilestoneOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = MilestoneOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid MilestoneOrderField", str)
	}
	return nil
}

func (e MilestoneOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The possible states of a milestone.
type MilestoneState string

const (
	// A milestone that is still open.
	MilestoneStateOpen MilestoneState = "OPEN"
	// A milestone that has been closed.
	MilestoneStateClosed MilestoneState = "CLOSED"
)

var AllMilestoneState = []MilestoneState{
	MilestoneStateOpen,
	MilestoneStateClosed,
}

func (e MilestoneState) IsValid() bool {
	switch e {
	case MilestoneStateOpen, MilestoneStateClosed:
		return true
	}
	return false
}

func (e MilestoneState) String() string {
	return string(e)
}

func (e *MilestoneState) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = MilestoneState(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid MilestoneState", str)
	}
	return nil
}

func (e MilestoneState) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The state of an OAuth Application when it was created.
type OauthApplicationCreateAuditEntryState string

const (
	// The OAuth Application was active and allowed to have OAuth Accesses.
	OauthApplicationCreateAuditEntryStateActive OauthApplicationCreateAuditEntryState = "ACTIVE"
	// The OAuth Application was suspended from generating OAuth Accesses due to abuse or security concerns.
	OauthApplicationCreateAuditEntryStateSuspended OauthApplicationCreateAuditEntryState = "SUSPENDED"
	// The OAuth Application was in the process of being deleted.
	OauthApplicationCreateAuditEntryStatePendingDeletion OauthApplicationCreateAuditEntryState = "PENDING_DELETION"
)

var AllOauthApplicationCreateAuditEntryState = []OauthApplicationCreateAuditEntryState{
	OauthApplicationCreateAuditEntryStateActive,
	OauthApplicationCreateAuditEntryStateSuspended,
	OauthApplicationCreateAuditEntryStatePendingDeletion,
}

func (e OauthApplicationCreateAuditEntryState) IsValid() bool {
	switch e {
	case OauthApplicationCreateAuditEntryStateActive, OauthApplicationCreateAuditEntryStateSuspended, OauthApplicationCreateAuditEntryStatePendingDeletion:
		return true
	}
	return false
}

func (e OauthApplicationCreateAuditEntryState) String() string {
	return string(e)
}

func (e *OauthApplicationCreateAuditEntryState) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OauthApplicationCreateAuditEntryState(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OauthApplicationCreateAuditEntryState", str)
	}
	return nil
}

func (e OauthApplicationCreateAuditEntryState) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The corresponding operation type for the action
type OperationType string

const (
	// An existing resource was accessed
	OperationTypeAccess OperationType = "ACCESS"
	// A resource performed an authentication event
	OperationTypeAuthentication OperationType = "AUTHENTICATION"
	// A new resource was created
	OperationTypeCreate OperationType = "CREATE"
	// An existing resource was modified
	OperationTypeModify OperationType = "MODIFY"
	// An existing resource was removed
	OperationTypeRemove OperationType = "REMOVE"
	// An existing resource was restored
	OperationTypeRestore OperationType = "RESTORE"
	// An existing resource was transferred between multiple resources
	OperationTypeTransfer OperationType = "TRANSFER"
)

var AllOperationType = []OperationType{
	OperationTypeAccess,
	OperationTypeAuthentication,
	OperationTypeCreate,
	OperationTypeModify,
	OperationTypeRemove,
	OperationTypeRestore,
	OperationTypeTransfer,
}

func (e OperationType) IsValid() bool {
	switch e {
	case OperationTypeAccess, OperationTypeAuthentication, OperationTypeCreate, OperationTypeModify, OperationTypeRemove, OperationTypeRestore, OperationTypeTransfer:
		return true
	}
	return false
}

func (e OperationType) String() string {
	return string(e)
}

func (e *OperationType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OperationType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OperationType", str)
	}
	return nil
}

func (e OperationType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Possible directions in which to order a list of items when provided an `orderBy` argument.
type OrderDirection string

const (
	// Specifies an ascending order for a given `orderBy` argument.
	OrderDirectionAsc OrderDirection = "ASC"
	// Specifies a descending order for a given `orderBy` argument.
	OrderDirectionDesc OrderDirection = "DESC"
)

var AllOrderDirection = []OrderDirection{
	OrderDirectionAsc,
	OrderDirectionDesc,
}

func (e OrderDirection) IsValid() bool {
	switch e {
	case OrderDirectionAsc, OrderDirectionDesc:
		return true
	}
	return false
}

func (e OrderDirection) String() string {
	return string(e)
}

func (e *OrderDirection) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrderDirection(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrderDirection", str)
	}
	return nil
}

func (e OrderDirection) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The permissions available to members on an Organization.
type OrgAddMemberAuditEntryPermission string

const (
	// Can read and clone repositories.
	OrgAddMemberAuditEntryPermissionRead OrgAddMemberAuditEntryPermission = "READ"
	// Can read, clone, push, and add collaborators to repositories.
	OrgAddMemberAuditEntryPermissionAdmin OrgAddMemberAuditEntryPermission = "ADMIN"
)

var AllOrgAddMemberAuditEntryPermission = []OrgAddMemberAuditEntryPermission{
	OrgAddMemberAuditEntryPermissionRead,
	OrgAddMemberAuditEntryPermissionAdmin,
}

func (e OrgAddMemberAuditEntryPermission) IsValid() bool {
	switch e {
	case OrgAddMemberAuditEntryPermissionRead, OrgAddMemberAuditEntryPermissionAdmin:
		return true
	}
	return false
}

func (e OrgAddMemberAuditEntryPermission) String() string {
	return string(e)
}

func (e *OrgAddMemberAuditEntryPermission) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrgAddMemberAuditEntryPermission(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrgAddMemberAuditEntryPermission", str)
	}
	return nil
}

func (e OrgAddMemberAuditEntryPermission) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The billing plans available for organizations.
type OrgCreateAuditEntryBillingPlan string

const (
	// Free Plan
	OrgCreateAuditEntryBillingPlanFree OrgCreateAuditEntryBillingPlan = "FREE"
	// Team Plan
	OrgCreateAuditEntryBillingPlanBusiness OrgCreateAuditEntryBillingPlan = "BUSINESS"
	// Enterprise Cloud Plan
	OrgCreateAuditEntryBillingPlanBusinessPlus OrgCreateAuditEntryBillingPlan = "BUSINESS_PLUS"
	// Legacy Unlimited Plan
	OrgCreateAuditEntryBillingPlanUnlimited OrgCreateAuditEntryBillingPlan = "UNLIMITED"
	// Tiered Per Seat Plan
	OrgCreateAuditEntryBillingPlanTieredPerSeat OrgCreateAuditEntryBillingPlan = "TIERED_PER_SEAT"
)

var AllOrgCreateAuditEntryBillingPlan = []OrgCreateAuditEntryBillingPlan{
	OrgCreateAuditEntryBillingPlanFree,
	OrgCreateAuditEntryBillingPlanBusiness,
	OrgCreateAuditEntryBillingPlanBusinessPlus,
	OrgCreateAuditEntryBillingPlanUnlimited,
	OrgCreateAuditEntryBillingPlanTieredPerSeat,
}

func (e OrgCreateAuditEntryBillingPlan) IsValid() bool {
	switch e {
	case OrgCreateAuditEntryBillingPlanFree, OrgCreateAuditEntryBillingPlanBusiness, OrgCreateAuditEntryBillingPlanBusinessPlus, OrgCreateAuditEntryBillingPlanUnlimited, OrgCreateAuditEntryBillingPlanTieredPerSeat:
		return true
	}
	return false
}

func (e OrgCreateAuditEntryBillingPlan) String() string {
	return string(e)
}

func (e *OrgCreateAuditEntryBillingPlan) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrgCreateAuditEntryBillingPlan(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrgCreateAuditEntryBillingPlan", str)
	}
	return nil
}

func (e OrgCreateAuditEntryBillingPlan) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The reason a billing manager was removed from an Organization.
type OrgRemoveBillingManagerAuditEntryReason string

const (
	// The organization required 2FA of its billing managers and this user did not have 2FA enabled.
	OrgRemoveBillingManagerAuditEntryReasonTwoFactorRequirementNonCompliance OrgRemoveBillingManagerAuditEntryReason = "TWO_FACTOR_REQUIREMENT_NON_COMPLIANCE"
	// SAML external identity missing
	OrgRemoveBillingManagerAuditEntryReasonSamlExternalIDEntityMissing OrgRemoveBillingManagerAuditEntryReason = "SAML_EXTERNAL_IDENTITY_MISSING"
	// SAML SSO enforcement requires an external identity
	OrgRemoveBillingManagerAuditEntryReasonSamlSsoEnforcementRequiresExternalIDEntity OrgRemoveBillingManagerAuditEntryReason = "SAML_SSO_ENFORCEMENT_REQUIRES_EXTERNAL_IDENTITY"
)

var AllOrgRemoveBillingManagerAuditEntryReason = []OrgRemoveBillingManagerAuditEntryReason{
	OrgRemoveBillingManagerAuditEntryReasonTwoFactorRequirementNonCompliance,
	OrgRemoveBillingManagerAuditEntryReasonSamlExternalIDEntityMissing,
	OrgRemoveBillingManagerAuditEntryReasonSamlSsoEnforcementRequiresExternalIDEntity,
}

func (e OrgRemoveBillingManagerAuditEntryReason) IsValid() bool {
	switch e {
	case OrgRemoveBillingManagerAuditEntryReasonTwoFactorRequirementNonCompliance, OrgRemoveBillingManagerAuditEntryReasonSamlExternalIDEntityMissing, OrgRemoveBillingManagerAuditEntryReasonSamlSsoEnforcementRequiresExternalIDEntity:
		return true
	}
	return false
}

func (e OrgRemoveBillingManagerAuditEntryReason) String() string {
	return string(e)
}

func (e *OrgRemoveBillingManagerAuditEntryReason) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrgRemoveBillingManagerAuditEntryReason(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrgRemoveBillingManagerAuditEntryReason", str)
	}
	return nil
}

func (e OrgRemoveBillingManagerAuditEntryReason) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The type of membership a user has with an Organization.
type OrgRemoveMemberAuditEntryMembershipType string

const (
	// A direct member is a user that is a member of the Organization.
	OrgRemoveMemberAuditEntryMembershipTypeDirectMember OrgRemoveMemberAuditEntryMembershipType = "DIRECT_MEMBER"
	// Organization administrators have full access and can change several settings, including the names of repositories that belong to the Organization and Owners team membership. In addition, organization admins can delete the organization and all of its repositories.
	OrgRemoveMemberAuditEntryMembershipTypeAdmin OrgRemoveMemberAuditEntryMembershipType = "ADMIN"
	// A billing manager is a user who manages the billing settings for the Organization, such as updating payment information.
	OrgRemoveMemberAuditEntryMembershipTypeBillingManager OrgRemoveMemberAuditEntryMembershipType = "BILLING_MANAGER"
	// An unaffiliated collaborator is a person who is not a member of the Organization and does not have access to any repositories in the Organization.
	OrgRemoveMemberAuditEntryMembershipTypeUnaffiliated OrgRemoveMemberAuditEntryMembershipType = "UNAFFILIATED"
	// An outside collaborator is a person who isn't explicitly a member of the Organization, but who has Read, Write, or Admin permissions to one or more repositories in the organization.
	OrgRemoveMemberAuditEntryMembershipTypeOutsideCollaborator OrgRemoveMemberAuditEntryMembershipType = "OUTSIDE_COLLABORATOR"
)

var AllOrgRemoveMemberAuditEntryMembershipType = []OrgRemoveMemberAuditEntryMembershipType{
	OrgRemoveMemberAuditEntryMembershipTypeDirectMember,
	OrgRemoveMemberAuditEntryMembershipTypeAdmin,
	OrgRemoveMemberAuditEntryMembershipTypeBillingManager,
	OrgRemoveMemberAuditEntryMembershipTypeUnaffiliated,
	OrgRemoveMemberAuditEntryMembershipTypeOutsideCollaborator,
}

func (e OrgRemoveMemberAuditEntryMembershipType) IsValid() bool {
	switch e {
	case OrgRemoveMemberAuditEntryMembershipTypeDirectMember, OrgRemoveMemberAuditEntryMembershipTypeAdmin, OrgRemoveMemberAuditEntryMembershipTypeBillingManager, OrgRemoveMemberAuditEntryMembershipTypeUnaffiliated, OrgRemoveMemberAuditEntryMembershipTypeOutsideCollaborator:
		return true
	}
	return false
}

func (e OrgRemoveMemberAuditEntryMembershipType) String() string {
	return string(e)
}

func (e *OrgRemoveMemberAuditEntryMembershipType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrgRemoveMemberAuditEntryMembershipType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrgRemoveMemberAuditEntryMembershipType", str)
	}
	return nil
}

func (e OrgRemoveMemberAuditEntryMembershipType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The reason a member was removed from an Organization.
type OrgRemoveMemberAuditEntryReason string

const (
	// The organization required 2FA of its billing managers and this user did not have 2FA enabled.
	OrgRemoveMemberAuditEntryReasonTwoFactorRequirementNonCompliance OrgRemoveMemberAuditEntryReason = "TWO_FACTOR_REQUIREMENT_NON_COMPLIANCE"
	// SAML external identity missing
	OrgRemoveMemberAuditEntryReasonSamlExternalIDEntityMissing OrgRemoveMemberAuditEntryReason = "SAML_EXTERNAL_IDENTITY_MISSING"
	// SAML SSO enforcement requires an external identity
	OrgRemoveMemberAuditEntryReasonSamlSsoEnforcementRequiresExternalIDEntity OrgRemoveMemberAuditEntryReason = "SAML_SSO_ENFORCEMENT_REQUIRES_EXTERNAL_IDENTITY"
	// User account has been deleted
	OrgRemoveMemberAuditEntryReasonUserAccountDeleted OrgRemoveMemberAuditEntryReason = "USER_ACCOUNT_DELETED"
	// User was removed from organization during account recovery
	OrgRemoveMemberAuditEntryReasonTwoFactorAccountRecovery OrgRemoveMemberAuditEntryReason = "TWO_FACTOR_ACCOUNT_RECOVERY"
)

var AllOrgRemoveMemberAuditEntryReason = []OrgRemoveMemberAuditEntryReason{
	OrgRemoveMemberAuditEntryReasonTwoFactorRequirementNonCompliance,
	OrgRemoveMemberAuditEntryReasonSamlExternalIDEntityMissing,
	OrgRemoveMemberAuditEntryReasonSamlSsoEnforcementRequiresExternalIDEntity,
	OrgRemoveMemberAuditEntryReasonUserAccountDeleted,
	OrgRemoveMemberAuditEntryReasonTwoFactorAccountRecovery,
}

func (e OrgRemoveMemberAuditEntryReason) IsValid() bool {
	switch e {
	case OrgRemoveMemberAuditEntryReasonTwoFactorRequirementNonCompliance, OrgRemoveMemberAuditEntryReasonSamlExternalIDEntityMissing, OrgRemoveMemberAuditEntryReasonSamlSsoEnforcementRequiresExternalIDEntity, OrgRemoveMemberAuditEntryReasonUserAccountDeleted, OrgRemoveMemberAuditEntryReasonTwoFactorAccountRecovery:
		return true
	}
	return false
}

func (e OrgRemoveMemberAuditEntryReason) String() string {
	return string(e)
}

func (e *OrgRemoveMemberAuditEntryReason) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrgRemoveMemberAuditEntryReason(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrgRemoveMemberAuditEntryReason", str)
	}
	return nil
}

func (e OrgRemoveMemberAuditEntryReason) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The type of membership a user has with an Organization.
type OrgRemoveOutsideCollaboratorAuditEntryMembershipType string

const (
	// An outside collaborator is a person who isn't explicitly a member of the Organization, but who has Read, Write, or Admin permissions to one or more repositories in the organization.
	OrgRemoveOutsideCollaboratorAuditEntryMembershipTypeOutsideCollaborator OrgRemoveOutsideCollaboratorAuditEntryMembershipType = "OUTSIDE_COLLABORATOR"
	// An unaffiliated collaborator is a person who is not a member of the Organization and does not have access to any repositories in the organization.
	OrgRemoveOutsideCollaboratorAuditEntryMembershipTypeUnaffiliated OrgRemoveOutsideCollaboratorAuditEntryMembershipType = "UNAFFILIATED"
	// A billing manager is a user who manages the billing settings for the Organization, such as updating payment information.
	OrgRemoveOutsideCollaboratorAuditEntryMembershipTypeBillingManager OrgRemoveOutsideCollaboratorAuditEntryMembershipType = "BILLING_MANAGER"
)

var AllOrgRemoveOutsideCollaboratorAuditEntryMembershipType = []OrgRemoveOutsideCollaboratorAuditEntryMembershipType{
	OrgRemoveOutsideCollaboratorAuditEntryMembershipTypeOutsideCollaborator,
	OrgRemoveOutsideCollaboratorAuditEntryMembershipTypeUnaffiliated,
	OrgRemoveOutsideCollaboratorAuditEntryMembershipTypeBillingManager,
}

func (e OrgRemoveOutsideCollaboratorAuditEntryMembershipType) IsValid() bool {
	switch e {
	case OrgRemoveOutsideCollaboratorAuditEntryMembershipTypeOutsideCollaborator, OrgRemoveOutsideCollaboratorAuditEntryMembershipTypeUnaffiliated, OrgRemoveOutsideCollaboratorAuditEntryMembershipTypeBillingManager:
		return true
	}
	return false
}

func (e OrgRemoveOutsideCollaboratorAuditEntryMembershipType) String() string {
	return string(e)
}

func (e *OrgRemoveOutsideCollaboratorAuditEntryMembershipType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrgRemoveOutsideCollaboratorAuditEntryMembershipType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrgRemoveOutsideCollaboratorAuditEntryMembershipType", str)
	}
	return nil
}

func (e OrgRemoveOutsideCollaboratorAuditEntryMembershipType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The reason an outside collaborator was removed from an Organization.
type OrgRemoveOutsideCollaboratorAuditEntryReason string

const (
	// The organization required 2FA of its billing managers and this user did not have 2FA enabled.
	OrgRemoveOutsideCollaboratorAuditEntryReasonTwoFactorRequirementNonCompliance OrgRemoveOutsideCollaboratorAuditEntryReason = "TWO_FACTOR_REQUIREMENT_NON_COMPLIANCE"
	// SAML external identity missing
	OrgRemoveOutsideCollaboratorAuditEntryReasonSamlExternalIDEntityMissing OrgRemoveOutsideCollaboratorAuditEntryReason = "SAML_EXTERNAL_IDENTITY_MISSING"
)

var AllOrgRemoveOutsideCollaboratorAuditEntryReason = []OrgRemoveOutsideCollaboratorAuditEntryReason{
	OrgRemoveOutsideCollaboratorAuditEntryReasonTwoFactorRequirementNonCompliance,
	OrgRemoveOutsideCollaboratorAuditEntryReasonSamlExternalIDEntityMissing,
}

func (e OrgRemoveOutsideCollaboratorAuditEntryReason) IsValid() bool {
	switch e {
	case OrgRemoveOutsideCollaboratorAuditEntryReasonTwoFactorRequirementNonCompliance, OrgRemoveOutsideCollaboratorAuditEntryReasonSamlExternalIDEntityMissing:
		return true
	}
	return false
}

func (e OrgRemoveOutsideCollaboratorAuditEntryReason) String() string {
	return string(e)
}

func (e *OrgRemoveOutsideCollaboratorAuditEntryReason) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrgRemoveOutsideCollaboratorAuditEntryReason(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrgRemoveOutsideCollaboratorAuditEntryReason", str)
	}
	return nil
}

func (e OrgRemoveOutsideCollaboratorAuditEntryReason) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The default permission a repository can have in an Organization.
type OrgUpdateDefaultRepositoryPermissionAuditEntryPermission string

const (
	// Can read and clone repositories.
	OrgUpdateDefaultRepositoryPermissionAuditEntryPermissionRead OrgUpdateDefaultRepositoryPermissionAuditEntryPermission = "READ"
	// Can read, clone and push to repositories.
	OrgUpdateDefaultRepositoryPermissionAuditEntryPermissionWrite OrgUpdateDefaultRepositoryPermissionAuditEntryPermission = "WRITE"
	// Can read, clone, push, and add collaborators to repositories.
	OrgUpdateDefaultRepositoryPermissionAuditEntryPermissionAdmin OrgUpdateDefaultRepositoryPermissionAuditEntryPermission = "ADMIN"
	// No default permission value.
	OrgUpdateDefaultRepositoryPermissionAuditEntryPermissionNone OrgUpdateDefaultRepositoryPermissionAuditEntryPermission = "NONE"
)

var AllOrgUpdateDefaultRepositoryPermissionAuditEntryPermission = []OrgUpdateDefaultRepositoryPermissionAuditEntryPermission{
	OrgUpdateDefaultRepositoryPermissionAuditEntryPermissionRead,
	OrgUpdateDefaultRepositoryPermissionAuditEntryPermissionWrite,
	OrgUpdateDefaultRepositoryPermissionAuditEntryPermissionAdmin,
	OrgUpdateDefaultRepositoryPermissionAuditEntryPermissionNone,
}

func (e OrgUpdateDefaultRepositoryPermissionAuditEntryPermission) IsValid() bool {
	switch e {
	case OrgUpdateDefaultRepositoryPermissionAuditEntryPermissionRead, OrgUpdateDefaultRepositoryPermissionAuditEntryPermissionWrite, OrgUpdateDefaultRepositoryPermissionAuditEntryPermissionAdmin, OrgUpdateDefaultRepositoryPermissionAuditEntryPermissionNone:
		return true
	}
	return false
}

func (e OrgUpdateDefaultRepositoryPermissionAuditEntryPermission) String() string {
	return string(e)
}

func (e *OrgUpdateDefaultRepositoryPermissionAuditEntryPermission) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrgUpdateDefaultRepositoryPermissionAuditEntryPermission(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrgUpdateDefaultRepositoryPermissionAuditEntryPermission", str)
	}
	return nil
}

func (e OrgUpdateDefaultRepositoryPermissionAuditEntryPermission) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The permissions available to members on an Organization.
type OrgUpdateMemberAuditEntryPermission string

const (
	// Can read and clone repositories.
	OrgUpdateMemberAuditEntryPermissionRead OrgUpdateMemberAuditEntryPermission = "READ"
	// Can read, clone, push, and add collaborators to repositories.
	OrgUpdateMemberAuditEntryPermissionAdmin OrgUpdateMemberAuditEntryPermission = "ADMIN"
)

var AllOrgUpdateMemberAuditEntryPermission = []OrgUpdateMemberAuditEntryPermission{
	OrgUpdateMemberAuditEntryPermissionRead,
	OrgUpdateMemberAuditEntryPermissionAdmin,
}

func (e OrgUpdateMemberAuditEntryPermission) IsValid() bool {
	switch e {
	case OrgUpdateMemberAuditEntryPermissionRead, OrgUpdateMemberAuditEntryPermissionAdmin:
		return true
	}
	return false
}

func (e OrgUpdateMemberAuditEntryPermission) String() string {
	return string(e)
}

func (e *OrgUpdateMemberAuditEntryPermission) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrgUpdateMemberAuditEntryPermission(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrgUpdateMemberAuditEntryPermission", str)
	}
	return nil
}

func (e OrgUpdateMemberAuditEntryPermission) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The permissions available for repository creation on an Organization.
type OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibility string

const (
	// All organization members are restricted from creating any repositories.
	OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibilityAll OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibility = "ALL"
	// All organization members are restricted from creating public repositories.
	OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibilityPublic OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibility = "PUBLIC"
	// All organization members are allowed to create any repositories.
	OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibilityNone OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibility = "NONE"
	// All organization members are restricted from creating private repositories.
	OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibilityPrivate OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibility = "PRIVATE"
	// All organization members are restricted from creating internal repositories.
	OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibilityInternal OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibility = "INTERNAL"
	// All organization members are restricted from creating public or internal repositories.
	OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibilityPublicInternal OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibility = "PUBLIC_INTERNAL"
	// All organization members are restricted from creating private or internal repositories.
	OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibilityPrivateInternal OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibility = "PRIVATE_INTERNAL"
	// All organization members are restricted from creating public or private repositories.
	OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibilityPublicPrivate OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibility = "PUBLIC_PRIVATE"
)

var AllOrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibility = []OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibility{
	OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibilityAll,
	OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibilityPublic,
	OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibilityNone,
	OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibilityPrivate,
	OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibilityInternal,
	OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibilityPublicInternal,
	OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibilityPrivateInternal,
	OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibilityPublicPrivate,
}

func (e OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibility) IsValid() bool {
	switch e {
	case OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibilityAll, OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibilityPublic, OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibilityNone, OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibilityPrivate, OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibilityInternal, OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibilityPublicInternal, OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibilityPrivateInternal, OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibilityPublicPrivate:
		return true
	}
	return false
}

func (e OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibility) String() string {
	return string(e)
}

func (e *OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibility) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibility(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibility", str)
	}
	return nil
}

func (e OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibility) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The possible organization invitation roles.
type OrganizationInvitationRole string

const (
	// The user is invited to be a direct member of the organization.
	OrganizationInvitationRoleDirectMember OrganizationInvitationRole = "DIRECT_MEMBER"
	// The user is invited to be an admin of the organization.
	OrganizationInvitationRoleAdmin OrganizationInvitationRole = "ADMIN"
	// The user is invited to be a billing manager of the organization.
	OrganizationInvitationRoleBillingManager OrganizationInvitationRole = "BILLING_MANAGER"
	// The user's previous role will be reinstated.
	OrganizationInvitationRoleReinstate OrganizationInvitationRole = "REINSTATE"
)

var AllOrganizationInvitationRole = []OrganizationInvitationRole{
	OrganizationInvitationRoleDirectMember,
	OrganizationInvitationRoleAdmin,
	OrganizationInvitationRoleBillingManager,
	OrganizationInvitationRoleReinstate,
}

func (e OrganizationInvitationRole) IsValid() bool {
	switch e {
	case OrganizationInvitationRoleDirectMember, OrganizationInvitationRoleAdmin, OrganizationInvitationRoleBillingManager, OrganizationInvitationRoleReinstate:
		return true
	}
	return false
}

func (e OrganizationInvitationRole) String() string {
	return string(e)
}

func (e *OrganizationInvitationRole) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrganizationInvitationRole(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrganizationInvitationRole", str)
	}
	return nil
}

func (e OrganizationInvitationRole) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The possible organization invitation types.
type OrganizationInvitationType string

const (
	// The invitation was to an existing user.
	OrganizationInvitationTypeUser OrganizationInvitationType = "USER"
	// The invitation was to an email address.
	OrganizationInvitationTypeEmail OrganizationInvitationType = "EMAIL"
)

var AllOrganizationInvitationType = []OrganizationInvitationType{
	OrganizationInvitationTypeUser,
	OrganizationInvitationTypeEmail,
}

func (e OrganizationInvitationType) IsValid() bool {
	switch e {
	case OrganizationInvitationTypeUser, OrganizationInvitationTypeEmail:
		return true
	}
	return false
}

func (e OrganizationInvitationType) String() string {
	return string(e)
}

func (e *OrganizationInvitationType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrganizationInvitationType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrganizationInvitationType", str)
	}
	return nil
}

func (e OrganizationInvitationType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The possible roles within an organization for its members.
type OrganizationMemberRole string

const (
	// The user is a member of the organization.
	OrganizationMemberRoleMember OrganizationMemberRole = "MEMBER"
	// The user is an administrator of the organization.
	OrganizationMemberRoleAdmin OrganizationMemberRole = "ADMIN"
)

var AllOrganizationMemberRole = []OrganizationMemberRole{
	OrganizationMemberRoleMember,
	OrganizationMemberRoleAdmin,
}

func (e OrganizationMemberRole) IsValid() bool {
	switch e {
	case OrganizationMemberRoleMember, OrganizationMemberRoleAdmin:
		return true
	}
	return false
}

func (e OrganizationMemberRole) String() string {
	return string(e)
}

func (e *OrganizationMemberRole) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrganizationMemberRole(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrganizationMemberRole", str)
	}
	return nil
}

func (e OrganizationMemberRole) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The possible values for the members can create repositories setting on an organization.
type OrganizationMembersCanCreateRepositoriesSettingValue string

const (
	// Members will be able to create public and private repositories.
	OrganizationMembersCanCreateRepositoriesSettingValueAll OrganizationMembersCanCreateRepositoriesSettingValue = "ALL"
	// Members will be able to create only private repositories.
	OrganizationMembersCanCreateRepositoriesSettingValuePrivate OrganizationMembersCanCreateRepositoriesSettingValue = "PRIVATE"
	// Members will not be able to create public or private repositories.
	OrganizationMembersCanCreateRepositoriesSettingValueDisabled OrganizationMembersCanCreateRepositoriesSettingValue = "DISABLED"
)

var AllOrganizationMembersCanCreateRepositoriesSettingValue = []OrganizationMembersCanCreateRepositoriesSettingValue{
	OrganizationMembersCanCreateRepositoriesSettingValueAll,
	OrganizationMembersCanCreateRepositoriesSettingValuePrivate,
	OrganizationMembersCanCreateRepositoriesSettingValueDisabled,
}

func (e OrganizationMembersCanCreateRepositoriesSettingValue) IsValid() bool {
	switch e {
	case OrganizationMembersCanCreateRepositoriesSettingValueAll, OrganizationMembersCanCreateRepositoriesSettingValuePrivate, OrganizationMembersCanCreateRepositoriesSettingValueDisabled:
		return true
	}
	return false
}

func (e OrganizationMembersCanCreateRepositoriesSettingValue) String() string {
	return string(e)
}

func (e *OrganizationMembersCanCreateRepositoriesSettingValue) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrganizationMembersCanCreateRepositoriesSettingValue(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrganizationMembersCanCreateRepositoriesSettingValue", str)
	}
	return nil
}

func (e OrganizationMembersCanCreateRepositoriesSettingValue) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Properties by which organization connections can be ordered.
type OrganizationOrderField string

const (
	// Order organizations by creation time
	OrganizationOrderFieldCreatedAt OrganizationOrderField = "CREATED_AT"
	// Order organizations by login
	OrganizationOrderFieldLogin OrganizationOrderField = "LOGIN"
)

var AllOrganizationOrderField = []OrganizationOrderField{
	OrganizationOrderFieldCreatedAt,
	OrganizationOrderFieldLogin,
}

func (e OrganizationOrderField) IsValid() bool {
	switch e {
	case OrganizationOrderFieldCreatedAt, OrganizationOrderFieldLogin:
		return true
	}
	return false
}

func (e OrganizationOrderField) String() string {
	return string(e)
}

func (e *OrganizationOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrganizationOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrganizationOrderField", str)
	}
	return nil
}

func (e OrganizationOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Properties by which package file connections can be ordered.
type PackageFileOrderField string

const (
	// Order package files by creation time
	PackageFileOrderFieldCreatedAt PackageFileOrderField = "CREATED_AT"
)

var AllPackageFileOrderField = []PackageFileOrderField{
	PackageFileOrderFieldCreatedAt,
}

func (e PackageFileOrderField) IsValid() bool {
	switch e {
	case PackageFileOrderFieldCreatedAt:
		return true
	}
	return false
}

func (e PackageFileOrderField) String() string {
	return string(e)
}

func (e *PackageFileOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PackageFileOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PackageFileOrderField", str)
	}
	return nil
}

func (e PackageFileOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Properties by which package connections can be ordered.
type PackageOrderField string

const (
	// Order packages by creation time
	PackageOrderFieldCreatedAt PackageOrderField = "CREATED_AT"
)

var AllPackageOrderField = []PackageOrderField{
	PackageOrderFieldCreatedAt,
}

func (e PackageOrderField) IsValid() bool {
	switch e {
	case PackageOrderFieldCreatedAt:
		return true
	}
	return false
}

func (e PackageOrderField) String() string {
	return string(e)
}

func (e *PackageOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PackageOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PackageOrderField", str)
	}
	return nil
}

func (e PackageOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The possible types of a package.
type PackageType string

const (
	// An npm package.
	PackageTypeNpm PackageType = "NPM"
	// A rubygems package.
	PackageTypeRubygems PackageType = "RUBYGEMS"
	// A maven package.
	PackageTypeMaven PackageType = "MAVEN"
	// A docker image.
	PackageTypeDocker PackageType = "DOCKER"
	// A debian package.
	PackageTypeDebian PackageType = "DEBIAN"
	// A nuget package.
	PackageTypeNuget PackageType = "NUGET"
	// A python package.
	PackageTypePypi PackageType = "PYPI"
)

var AllPackageType = []PackageType{
	PackageTypeNpm,
	PackageTypeRubygems,
	PackageTypeMaven,
	PackageTypeDocker,
	PackageTypeDebian,
	PackageTypeNuget,
	PackageTypePypi,
}

func (e PackageType) IsValid() bool {
	switch e {
	case PackageTypeNpm, PackageTypeRubygems, PackageTypeMaven, PackageTypeDocker, PackageTypeDebian, PackageTypeNuget, PackageTypePypi:
		return true
	}
	return false
}

func (e PackageType) String() string {
	return string(e)
}

func (e *PackageType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PackageType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PackageType", str)
	}
	return nil
}

func (e PackageType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Properties by which package version connections can be ordered.
type PackageVersionOrderField string

const (
	// Order package versions by creation time
	PackageVersionOrderFieldCreatedAt PackageVersionOrderField = "CREATED_AT"
)

var AllPackageVersionOrderField = []PackageVersionOrderField{
	PackageVersionOrderFieldCreatedAt,
}

func (e PackageVersionOrderField) IsValid() bool {
	switch e {
	case PackageVersionOrderFieldCreatedAt:
		return true
	}
	return false
}

func (e PackageVersionOrderField) String() string {
	return string(e)
}

func (e *PackageVersionOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PackageVersionOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PackageVersionOrderField", str)
	}
	return nil
}

func (e PackageVersionOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Represents items that can be pinned to a profile page or dashboard.
type PinnableItemType string

const (
	// A repository.
	PinnableItemTypeRepository PinnableItemType = "REPOSITORY"
	// A gist.
	PinnableItemTypeGist PinnableItemType = "GIST"
	// An issue.
	PinnableItemTypeIssue PinnableItemType = "ISSUE"
	// A project.
	PinnableItemTypeProject PinnableItemType = "PROJECT"
	// A pull request.
	PinnableItemTypePullRequest PinnableItemType = "PULL_REQUEST"
	// A user.
	PinnableItemTypeUser PinnableItemType = "USER"
	// An organization.
	PinnableItemTypeOrganization PinnableItemType = "ORGANIZATION"
	// A team.
	PinnableItemTypeTeam PinnableItemType = "TEAM"
)

var AllPinnableItemType = []PinnableItemType{
	PinnableItemTypeRepository,
	PinnableItemTypeGist,
	PinnableItemTypeIssue,
	PinnableItemTypeProject,
	PinnableItemTypePullRequest,
	PinnableItemTypeUser,
	PinnableItemTypeOrganization,
	PinnableItemTypeTeam,
}

func (e PinnableItemType) IsValid() bool {
	switch e {
	case PinnableItemTypeRepository, PinnableItemTypeGist, PinnableItemTypeIssue, PinnableItemTypeProject, PinnableItemTypePullRequest, PinnableItemTypeUser, PinnableItemTypeOrganization, PinnableItemTypeTeam:
		return true
	}
	return false
}

func (e PinnableItemType) String() string {
	return string(e)
}

func (e *PinnableItemType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PinnableItemType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PinnableItemType", str)
	}
	return nil
}

func (e PinnableItemType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The possible archived states of a project card.
type ProjectCardArchivedState string

const (
	// A project card that is archived
	ProjectCardArchivedStateArchived ProjectCardArchivedState = "ARCHIVED"
	// A project card that is not archived
	ProjectCardArchivedStateNotArchived ProjectCardArchivedState = "NOT_ARCHIVED"
)

var AllProjectCardArchivedState = []ProjectCardArchivedState{
	ProjectCardArchivedStateArchived,
	ProjectCardArchivedStateNotArchived,
}

func (e ProjectCardArchivedState) IsValid() bool {
	switch e {
	case ProjectCardArchivedStateArchived, ProjectCardArchivedStateNotArchived:
		return true
	}
	return false
}

func (e ProjectCardArchivedState) String() string {
	return string(e)
}

func (e *ProjectCardArchivedState) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ProjectCardArchivedState(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ProjectCardArchivedState", str)
	}
	return nil
}

func (e ProjectCardArchivedState) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Various content states of a ProjectCard
type ProjectCardState string

const (
	// The card has content only.
	ProjectCardStateContentOnly ProjectCardState = "CONTENT_ONLY"
	// The card has a note only.
	ProjectCardStateNoteOnly ProjectCardState = "NOTE_ONLY"
	// The card is redacted.
	ProjectCardStateRedacted ProjectCardState = "REDACTED"
)

var AllProjectCardState = []ProjectCardState{
	ProjectCardStateContentOnly,
	ProjectCardStateNoteOnly,
	ProjectCardStateRedacted,
}

func (e ProjectCardState) IsValid() bool {
	switch e {
	case ProjectCardStateContentOnly, ProjectCardStateNoteOnly, ProjectCardStateRedacted:
		return true
	}
	return false
}

func (e ProjectCardState) String() string {
	return string(e)
}

func (e *ProjectCardState) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ProjectCardState(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ProjectCardState", str)
	}
	return nil
}

func (e ProjectCardState) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The semantic purpose of the column - todo, in progress, or done.
type ProjectColumnPurpose string

const (
	// The column contains cards still to be worked on
	ProjectColumnPurposeTodo ProjectColumnPurpose = "TODO"
	// The column contains cards which are currently being worked on
	ProjectColumnPurposeInProgress ProjectColumnPurpose = "IN_PROGRESS"
	// The column contains cards which are complete
	ProjectColumnPurposeDone ProjectColumnPurpose = "DONE"
)

var AllProjectColumnPurpose = []ProjectColumnPurpose{
	ProjectColumnPurposeTodo,
	ProjectColumnPurposeInProgress,
	ProjectColumnPurposeDone,
}

func (e ProjectColumnPurpose) IsValid() bool {
	switch e {
	case ProjectColumnPurposeTodo, ProjectColumnPurposeInProgress, ProjectColumnPurposeDone:
		return true
	}
	return false
}

func (e ProjectColumnPurpose) String() string {
	return string(e)
}

func (e *ProjectColumnPurpose) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ProjectColumnPurpose(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ProjectColumnPurpose", str)
	}
	return nil
}

func (e ProjectColumnPurpose) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Properties by which project connections can be ordered.
type ProjectOrderField string

const (
	// Order projects by creation time
	ProjectOrderFieldCreatedAt ProjectOrderField = "CREATED_AT"
	// Order projects by update time
	ProjectOrderFieldUpdatedAt ProjectOrderField = "UPDATED_AT"
	// Order projects by name
	ProjectOrderFieldName ProjectOrderField = "NAME"
)

var AllProjectOrderField = []ProjectOrderField{
	ProjectOrderFieldCreatedAt,
	ProjectOrderFieldUpdatedAt,
	ProjectOrderFieldName,
}

func (e ProjectOrderField) IsValid() bool {
	switch e {
	case ProjectOrderFieldCreatedAt, ProjectOrderFieldUpdatedAt, ProjectOrderFieldName:
		return true
	}
	return false
}

func (e ProjectOrderField) String() string {
	return string(e)
}

func (e *ProjectOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ProjectOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ProjectOrderField", str)
	}
	return nil
}

func (e ProjectOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// State of the project; either 'open' or 'closed'
type ProjectState string

const (
	// The project is open.
	ProjectStateOpen ProjectState = "OPEN"
	// The project is closed.
	ProjectStateClosed ProjectState = "CLOSED"
)

var AllProjectState = []ProjectState{
	ProjectStateOpen,
	ProjectStateClosed,
}

func (e ProjectState) IsValid() bool {
	switch e {
	case ProjectStateOpen, ProjectStateClosed:
		return true
	}
	return false
}

func (e ProjectState) String() string {
	return string(e)
}

func (e *ProjectState) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ProjectState(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ProjectState", str)
	}
	return nil
}

func (e ProjectState) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// GitHub-provided templates for Projects
type ProjectTemplate string

const (
	// Create a board with columns for To do, In progress and Done.
	ProjectTemplateBasicKanban ProjectTemplate = "BASIC_KANBAN"
	// Create a board with v2 triggers to automatically move cards across To do, In progress and Done columns.
	ProjectTemplateAutomatedKanbanV2 ProjectTemplate = "AUTOMATED_KANBAN_V2"
	// Create a board with triggers to automatically move cards across columns with review automation.
	ProjectTemplateAutomatedReviewsKanban ProjectTemplate = "AUTOMATED_REVIEWS_KANBAN"
	// Create a board to triage and prioritize bugs with To do, priority, and Done columns.
	ProjectTemplateBugTriage ProjectTemplate = "BUG_TRIAGE"
)

var AllProjectTemplate = []ProjectTemplate{
	ProjectTemplateBasicKanban,
	ProjectTemplateAutomatedKanbanV2,
	ProjectTemplateAutomatedReviewsKanban,
	ProjectTemplateBugTriage,
}

func (e ProjectTemplate) IsValid() bool {
	switch e {
	case ProjectTemplateBasicKanban, ProjectTemplateAutomatedKanbanV2, ProjectTemplateAutomatedReviewsKanban, ProjectTemplateBugTriage:
		return true
	}
	return false
}

func (e ProjectTemplate) String() string {
	return string(e)
}

func (e *ProjectTemplate) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ProjectTemplate(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ProjectTemplate", str)
	}
	return nil
}

func (e ProjectTemplate) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Represents available types of methods to use when merging a pull request.
type PullRequestMergeMethod string

const (
	// Add all commits from the head branch to the base branch with a merge commit.
	PullRequestMergeMethodMerge PullRequestMergeMethod = "MERGE"
	// Combine all commits from the head branch into a single commit in the base branch.
	PullRequestMergeMethodSquash PullRequestMergeMethod = "SQUASH"
	// Add all commits from the head branch onto the base branch individually.
	PullRequestMergeMethodRebase PullRequestMergeMethod = "REBASE"
)

var AllPullRequestMergeMethod = []PullRequestMergeMethod{
	PullRequestMergeMethodMerge,
	PullRequestMergeMethodSquash,
	PullRequestMergeMethodRebase,
}

func (e PullRequestMergeMethod) IsValid() bool {
	switch e {
	case PullRequestMergeMethodMerge, PullRequestMergeMethodSquash, PullRequestMergeMethodRebase:
		return true
	}
	return false
}

func (e PullRequestMergeMethod) String() string {
	return string(e)
}

func (e *PullRequestMergeMethod) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PullRequestMergeMethod(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PullRequestMergeMethod", str)
	}
	return nil
}

func (e PullRequestMergeMethod) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Properties by which pull_requests connections can be ordered.
type PullRequestOrderField string

const (
	// Order pull_requests by creation time
	PullRequestOrderFieldCreatedAt PullRequestOrderField = "CREATED_AT"
	// Order pull_requests by update time
	PullRequestOrderFieldUpdatedAt PullRequestOrderField = "UPDATED_AT"
)

var AllPullRequestOrderField = []PullRequestOrderField{
	PullRequestOrderFieldCreatedAt,
	PullRequestOrderFieldUpdatedAt,
}

func (e PullRequestOrderField) IsValid() bool {
	switch e {
	case PullRequestOrderFieldCreatedAt, PullRequestOrderFieldUpdatedAt:
		return true
	}
	return false
}

func (e PullRequestOrderField) String() string {
	return string(e)
}

func (e *PullRequestOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PullRequestOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PullRequestOrderField", str)
	}
	return nil
}

func (e PullRequestOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The possible states of a pull request review comment.
type PullRequestReviewCommentState string

const (
	// A comment that is part of a pending review
	PullRequestReviewCommentStatePending PullRequestReviewCommentState = "PENDING"
	// A comment that is part of a submitted review
	PullRequestReviewCommentStateSubmitted PullRequestReviewCommentState = "SUBMITTED"
)

var AllPullRequestReviewCommentState = []PullRequestReviewCommentState{
	PullRequestReviewCommentStatePending,
	PullRequestReviewCommentStateSubmitted,
}

func (e PullRequestReviewCommentState) IsValid() bool {
	switch e {
	case PullRequestReviewCommentStatePending, PullRequestReviewCommentStateSubmitted:
		return true
	}
	return false
}

func (e PullRequestReviewCommentState) String() string {
	return string(e)
}

func (e *PullRequestReviewCommentState) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PullRequestReviewCommentState(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PullRequestReviewCommentState", str)
	}
	return nil
}

func (e PullRequestReviewCommentState) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The review status of a pull request.
type PullRequestReviewDecision string

const (
	// Changes have been requested on the pull request.
	PullRequestReviewDecisionChangesRequested PullRequestReviewDecision = "CHANGES_REQUESTED"
	// The pull request has received an approving review.
	PullRequestReviewDecisionApproved PullRequestReviewDecision = "APPROVED"
	// A review is required before the pull request can be merged.
	PullRequestReviewDecisionReviewRequired PullRequestReviewDecision = "REVIEW_REQUIRED"
)

var AllPullRequestReviewDecision = []PullRequestReviewDecision{
	PullRequestReviewDecisionChangesRequested,
	PullRequestReviewDecisionApproved,
	PullRequestReviewDecisionReviewRequired,
}

func (e PullRequestReviewDecision) IsValid() bool {
	switch e {
	case PullRequestReviewDecisionChangesRequested, PullRequestReviewDecisionApproved, PullRequestReviewDecisionReviewRequired:
		return true
	}
	return false
}

func (e PullRequestReviewDecision) String() string {
	return string(e)
}

func (e *PullRequestReviewDecision) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PullRequestReviewDecision(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PullRequestReviewDecision", str)
	}
	return nil
}

func (e PullRequestReviewDecision) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The possible events to perform on a pull request review.
type PullRequestReviewEvent string

const (
	// Submit general feedback without explicit approval.
	PullRequestReviewEventComment PullRequestReviewEvent = "COMMENT"
	// Submit feedback and approve merging these changes.
	PullRequestReviewEventApprove PullRequestReviewEvent = "APPROVE"
	// Submit feedback that must be addressed before merging.
	PullRequestReviewEventRequestChanges PullRequestReviewEvent = "REQUEST_CHANGES"
	// Dismiss review so it now longer effects merging.
	PullRequestReviewEventDismiss PullRequestReviewEvent = "DISMISS"
)

var AllPullRequestReviewEvent = []PullRequestReviewEvent{
	PullRequestReviewEventComment,
	PullRequestReviewEventApprove,
	PullRequestReviewEventRequestChanges,
	PullRequestReviewEventDismiss,
}

func (e PullRequestReviewEvent) IsValid() bool {
	switch e {
	case PullRequestReviewEventComment, PullRequestReviewEventApprove, PullRequestReviewEventRequestChanges, PullRequestReviewEventDismiss:
		return true
	}
	return false
}

func (e PullRequestReviewEvent) String() string {
	return string(e)
}

func (e *PullRequestReviewEvent) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PullRequestReviewEvent(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PullRequestReviewEvent", str)
	}
	return nil
}

func (e PullRequestReviewEvent) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The possible states of a pull request review.
type PullRequestReviewState string

const (
	// A review that has not yet been submitted.
	PullRequestReviewStatePending PullRequestReviewState = "PENDING"
	// An informational review.
	PullRequestReviewStateCommented PullRequestReviewState = "COMMENTED"
	// A review allowing the pull request to merge.
	PullRequestReviewStateApproved PullRequestReviewState = "APPROVED"
	// A review blocking the pull request from merging.
	PullRequestReviewStateChangesRequested PullRequestReviewState = "CHANGES_REQUESTED"
	// A review that has been dismissed.
	PullRequestReviewStateDismissed PullRequestReviewState = "DISMISSED"
)

var AllPullRequestReviewState = []PullRequestReviewState{
	PullRequestReviewStatePending,
	PullRequestReviewStateCommented,
	PullRequestReviewStateApproved,
	PullRequestReviewStateChangesRequested,
	PullRequestReviewStateDismissed,
}

func (e PullRequestReviewState) IsValid() bool {
	switch e {
	case PullRequestReviewStatePending, PullRequestReviewStateCommented, PullRequestReviewStateApproved, PullRequestReviewStateChangesRequested, PullRequestReviewStateDismissed:
		return true
	}
	return false
}

func (e PullRequestReviewState) String() string {
	return string(e)
}

func (e *PullRequestReviewState) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PullRequestReviewState(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PullRequestReviewState", str)
	}
	return nil
}

func (e PullRequestReviewState) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The possible states of a pull request.
type PullRequestState string

const (
	// A pull request that is still open.
	PullRequestStateOpen PullRequestState = "OPEN"
	// A pull request that has been closed without being merged.
	PullRequestStateClosed PullRequestState = "CLOSED"
	// A pull request that has been closed by being merged.
	PullRequestStateMerged PullRequestState = "MERGED"
)

var AllPullRequestState = []PullRequestState{
	PullRequestStateOpen,
	PullRequestStateClosed,
	PullRequestStateMerged,
}

func (e PullRequestState) IsValid() bool {
	switch e {
	case PullRequestStateOpen, PullRequestStateClosed, PullRequestStateMerged:
		return true
	}
	return false
}

func (e PullRequestState) String() string {
	return string(e)
}

func (e *PullRequestState) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PullRequestState(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PullRequestState", str)
	}
	return nil
}

func (e PullRequestState) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The possible item types found in a timeline.
type PullRequestTimelineItemsItemType string

const (
	// Represents a Git commit part of a pull request.
	PullRequestTimelineItemsItemTypePullRequestCommit PullRequestTimelineItemsItemType = "PULL_REQUEST_COMMIT"
	// Represents a commit comment thread part of a pull request.
	PullRequestTimelineItemsItemTypePullRequestCommitCommentThread PullRequestTimelineItemsItemType = "PULL_REQUEST_COMMIT_COMMENT_THREAD"
	// A review object for a given pull request.
	PullRequestTimelineItemsItemTypePullRequestReview PullRequestTimelineItemsItemType = "PULL_REQUEST_REVIEW"
	// A threaded list of comments for a given pull request.
	PullRequestTimelineItemsItemTypePullRequestReviewThread PullRequestTimelineItemsItemType = "PULL_REQUEST_REVIEW_THREAD"
	// Represents the latest point in the pull request timeline for which the viewer has seen the pull request's commits.
	PullRequestTimelineItemsItemTypePullRequestRevisionMarker PullRequestTimelineItemsItemType = "PULL_REQUEST_REVISION_MARKER"
	// Represents a 'automatic_base_change_failed' event on a given pull request.
	PullRequestTimelineItemsItemTypeAutomaticBaseChangeFailedEvent PullRequestTimelineItemsItemType = "AUTOMATIC_BASE_CHANGE_FAILED_EVENT"
	// Represents a 'automatic_base_change_succeeded' event on a given pull request.
	PullRequestTimelineItemsItemTypeAutomaticBaseChangeSucceededEvent PullRequestTimelineItemsItemType = "AUTOMATIC_BASE_CHANGE_SUCCEEDED_EVENT"
	// Represents a 'base_ref_changed' event on a given issue or pull request.
	PullRequestTimelineItemsItemTypeBaseRefChangedEvent PullRequestTimelineItemsItemType = "BASE_REF_CHANGED_EVENT"
	// Represents a 'base_ref_force_pushed' event on a given pull request.
	PullRequestTimelineItemsItemTypeBaseRefForcePushedEvent PullRequestTimelineItemsItemType = "BASE_REF_FORCE_PUSHED_EVENT"
	// Represents a 'base_ref_deleted' event on a given pull request.
	PullRequestTimelineItemsItemTypeBaseRefDeletedEvent PullRequestTimelineItemsItemType = "BASE_REF_DELETED_EVENT"
	// Represents a 'deployed' event on a given pull request.
	PullRequestTimelineItemsItemTypeDeployedEvent PullRequestTimelineItemsItemType = "DEPLOYED_EVENT"
	// Represents a 'deployment_environment_changed' event on a given pull request.
	PullRequestTimelineItemsItemTypeDeploymentEnvironmentChangedEvent PullRequestTimelineItemsItemType = "DEPLOYMENT_ENVIRONMENT_CHANGED_EVENT"
	// Represents a 'head_ref_deleted' event on a given pull request.
	PullRequestTimelineItemsItemTypeHeadRefDeletedEvent PullRequestTimelineItemsItemType = "HEAD_REF_DELETED_EVENT"
	// Represents a 'head_ref_force_pushed' event on a given pull request.
	PullRequestTimelineItemsItemTypeHeadRefForcePushedEvent PullRequestTimelineItemsItemType = "HEAD_REF_FORCE_PUSHED_EVENT"
	// Represents a 'head_ref_restored' event on a given pull request.
	PullRequestTimelineItemsItemTypeHeadRefRestoredEvent PullRequestTimelineItemsItemType = "HEAD_REF_RESTORED_EVENT"
	// Represents a 'merged' event on a given pull request.
	PullRequestTimelineItemsItemTypeMergedEvent PullRequestTimelineItemsItemType = "MERGED_EVENT"
	// Represents a 'review_dismissed' event on a given issue or pull request.
	PullRequestTimelineItemsItemTypeReviewDismissedEvent PullRequestTimelineItemsItemType = "REVIEW_DISMISSED_EVENT"
	// Represents an 'review_requested' event on a given pull request.
	PullRequestTimelineItemsItemTypeReviewRequestedEvent PullRequestTimelineItemsItemType = "REVIEW_REQUESTED_EVENT"
	// Represents an 'review_request_removed' event on a given pull request.
	PullRequestTimelineItemsItemTypeReviewRequestRemovedEvent PullRequestTimelineItemsItemType = "REVIEW_REQUEST_REMOVED_EVENT"
	// Represents a 'ready_for_review' event on a given pull request.
	PullRequestTimelineItemsItemTypeReadyForReviewEvent PullRequestTimelineItemsItemType = "READY_FOR_REVIEW_EVENT"
	// Represents a 'convert_to_draft' event on a given pull request.
	PullRequestTimelineItemsItemTypeConvertToDraftEvent PullRequestTimelineItemsItemType = "CONVERT_TO_DRAFT_EVENT"
	// Represents a comment on an Issue.
	PullRequestTimelineItemsItemTypeIssueComment PullRequestTimelineItemsItemType = "ISSUE_COMMENT"
	// Represents a mention made by one issue or pull request to another.
	PullRequestTimelineItemsItemTypeCrossReferencedEvent PullRequestTimelineItemsItemType = "CROSS_REFERENCED_EVENT"
	// Represents a 'added_to_project' event on a given issue or pull request.
	PullRequestTimelineItemsItemTypeAddedToProjectEvent PullRequestTimelineItemsItemType = "ADDED_TO_PROJECT_EVENT"
	// Represents an 'assigned' event on any assignable object.
	PullRequestTimelineItemsItemTypeAssignedEvent PullRequestTimelineItemsItemType = "ASSIGNED_EVENT"
	// Represents a 'closed' event on any `Closable`.
	PullRequestTimelineItemsItemTypeClosedEvent PullRequestTimelineItemsItemType = "CLOSED_EVENT"
	// Represents a 'comment_deleted' event on a given issue or pull request.
	PullRequestTimelineItemsItemTypeCommentDeletedEvent PullRequestTimelineItemsItemType = "COMMENT_DELETED_EVENT"
	// Represents a 'connected' event on a given issue or pull request.
	PullRequestTimelineItemsItemTypeConnectedEvent PullRequestTimelineItemsItemType = "CONNECTED_EVENT"
	// Represents a 'converted_note_to_issue' event on a given issue or pull request.
	PullRequestTimelineItemsItemTypeConvertedNoteToIssueEvent PullRequestTimelineItemsItemType = "CONVERTED_NOTE_TO_ISSUE_EVENT"
	// Represents a 'demilestoned' event on a given issue or pull request.
	PullRequestTimelineItemsItemTypeDemilestonedEvent PullRequestTimelineItemsItemType = "DEMILESTONED_EVENT"
	// Represents a 'disconnected' event on a given issue or pull request.
	PullRequestTimelineItemsItemTypeDisconnectedEvent PullRequestTimelineItemsItemType = "DISCONNECTED_EVENT"
	// Represents a 'labeled' event on a given issue or pull request.
	PullRequestTimelineItemsItemTypeLabeledEvent PullRequestTimelineItemsItemType = "LABELED_EVENT"
	// Represents a 'locked' event on a given issue or pull request.
	PullRequestTimelineItemsItemTypeLockedEvent PullRequestTimelineItemsItemType = "LOCKED_EVENT"
	// Represents a 'marked_as_duplicate' event on a given issue or pull request.
	PullRequestTimelineItemsItemTypeMarkedAsDuplicateEvent PullRequestTimelineItemsItemType = "MARKED_AS_DUPLICATE_EVENT"
	// Represents a 'mentioned' event on a given issue or pull request.
	PullRequestTimelineItemsItemTypeMentionedEvent PullRequestTimelineItemsItemType = "MENTIONED_EVENT"
	// Represents a 'milestoned' event on a given issue or pull request.
	PullRequestTimelineItemsItemTypeMilestonedEvent PullRequestTimelineItemsItemType = "MILESTONED_EVENT"
	// Represents a 'moved_columns_in_project' event on a given issue or pull request.
	PullRequestTimelineItemsItemTypeMovedColumnsInProjectEvent PullRequestTimelineItemsItemType = "MOVED_COLUMNS_IN_PROJECT_EVENT"
	// Represents a 'pinned' event on a given issue or pull request.
	PullRequestTimelineItemsItemTypePinnedEvent PullRequestTimelineItemsItemType = "PINNED_EVENT"
	// Represents a 'referenced' event on a given `ReferencedSubject`.
	PullRequestTimelineItemsItemTypeReferencedEvent PullRequestTimelineItemsItemType = "REFERENCED_EVENT"
	// Represents a 'removed_from_project' event on a given issue or pull request.
	PullRequestTimelineItemsItemTypeRemovedFromProjectEvent PullRequestTimelineItemsItemType = "REMOVED_FROM_PROJECT_EVENT"
	// Represents a 'renamed' event on a given issue or pull request
	PullRequestTimelineItemsItemTypeRenamedTitleEvent PullRequestTimelineItemsItemType = "RENAMED_TITLE_EVENT"
	// Represents a 'reopened' event on any `Closable`.
	PullRequestTimelineItemsItemTypeReopenedEvent PullRequestTimelineItemsItemType = "REOPENED_EVENT"
	// Represents a 'subscribed' event on a given `Subscribable`.
	PullRequestTimelineItemsItemTypeSubscribedEvent PullRequestTimelineItemsItemType = "SUBSCRIBED_EVENT"
	// Represents a 'transferred' event on a given issue or pull request.
	PullRequestTimelineItemsItemTypeTransferredEvent PullRequestTimelineItemsItemType = "TRANSFERRED_EVENT"
	// Represents an 'unassigned' event on any assignable object.
	PullRequestTimelineItemsItemTypeUnassignedEvent PullRequestTimelineItemsItemType = "UNASSIGNED_EVENT"
	// Represents an 'unlabeled' event on a given issue or pull request.
	PullRequestTimelineItemsItemTypeUnlabeledEvent PullRequestTimelineItemsItemType = "UNLABELED_EVENT"
	// Represents an 'unlocked' event on a given issue or pull request.
	PullRequestTimelineItemsItemTypeUnlockedEvent PullRequestTimelineItemsItemType = "UNLOCKED_EVENT"
	// Represents a 'user_blocked' event on a given user.
	PullRequestTimelineItemsItemTypeUserBlockedEvent PullRequestTimelineItemsItemType = "USER_BLOCKED_EVENT"
	// Represents an 'unmarked_as_duplicate' event on a given issue or pull request.
	PullRequestTimelineItemsItemTypeUnmarkedAsDuplicateEvent PullRequestTimelineItemsItemType = "UNMARKED_AS_DUPLICATE_EVENT"
	// Represents an 'unpinned' event on a given issue or pull request.
	PullRequestTimelineItemsItemTypeUnpinnedEvent PullRequestTimelineItemsItemType = "UNPINNED_EVENT"
	// Represents an 'unsubscribed' event on a given `Subscribable`.
	PullRequestTimelineItemsItemTypeUnsubscribedEvent PullRequestTimelineItemsItemType = "UNSUBSCRIBED_EVENT"
)

var AllPullRequestTimelineItemsItemType = []PullRequestTimelineItemsItemType{
	PullRequestTimelineItemsItemTypePullRequestCommit,
	PullRequestTimelineItemsItemTypePullRequestCommitCommentThread,
	PullRequestTimelineItemsItemTypePullRequestReview,
	PullRequestTimelineItemsItemTypePullRequestReviewThread,
	PullRequestTimelineItemsItemTypePullRequestRevisionMarker,
	PullRequestTimelineItemsItemTypeAutomaticBaseChangeFailedEvent,
	PullRequestTimelineItemsItemTypeAutomaticBaseChangeSucceededEvent,
	PullRequestTimelineItemsItemTypeBaseRefChangedEvent,
	PullRequestTimelineItemsItemTypeBaseRefForcePushedEvent,
	PullRequestTimelineItemsItemTypeBaseRefDeletedEvent,
	PullRequestTimelineItemsItemTypeDeployedEvent,
	PullRequestTimelineItemsItemTypeDeploymentEnvironmentChangedEvent,
	PullRequestTimelineItemsItemTypeHeadRefDeletedEvent,
	PullRequestTimelineItemsItemTypeHeadRefForcePushedEvent,
	PullRequestTimelineItemsItemTypeHeadRefRestoredEvent,
	PullRequestTimelineItemsItemTypeMergedEvent,
	PullRequestTimelineItemsItemTypeReviewDismissedEvent,
	PullRequestTimelineItemsItemTypeReviewRequestedEvent,
	PullRequestTimelineItemsItemTypeReviewRequestRemovedEvent,
	PullRequestTimelineItemsItemTypeReadyForReviewEvent,
	PullRequestTimelineItemsItemTypeConvertToDraftEvent,
	PullRequestTimelineItemsItemTypeIssueComment,
	PullRequestTimelineItemsItemTypeCrossReferencedEvent,
	PullRequestTimelineItemsItemTypeAddedToProjectEvent,
	PullRequestTimelineItemsItemTypeAssignedEvent,
	PullRequestTimelineItemsItemTypeClosedEvent,
	PullRequestTimelineItemsItemTypeCommentDeletedEvent,
	PullRequestTimelineItemsItemTypeConnectedEvent,
	PullRequestTimelineItemsItemTypeConvertedNoteToIssueEvent,
	PullRequestTimelineItemsItemTypeDemilestonedEvent,
	PullRequestTimelineItemsItemTypeDisconnectedEvent,
	PullRequestTimelineItemsItemTypeLabeledEvent,
	PullRequestTimelineItemsItemTypeLockedEvent,
	PullRequestTimelineItemsItemTypeMarkedAsDuplicateEvent,
	PullRequestTimelineItemsItemTypeMentionedEvent,
	PullRequestTimelineItemsItemTypeMilestonedEvent,
	PullRequestTimelineItemsItemTypeMovedColumnsInProjectEvent,
	PullRequestTimelineItemsItemTypePinnedEvent,
	PullRequestTimelineItemsItemTypeReferencedEvent,
	PullRequestTimelineItemsItemTypeRemovedFromProjectEvent,
	PullRequestTimelineItemsItemTypeRenamedTitleEvent,
	PullRequestTimelineItemsItemTypeReopenedEvent,
	PullRequestTimelineItemsItemTypeSubscribedEvent,
	PullRequestTimelineItemsItemTypeTransferredEvent,
	PullRequestTimelineItemsItemTypeUnassignedEvent,
	PullRequestTimelineItemsItemTypeUnlabeledEvent,
	PullRequestTimelineItemsItemTypeUnlockedEvent,
	PullRequestTimelineItemsItemTypeUserBlockedEvent,
	PullRequestTimelineItemsItemTypeUnmarkedAsDuplicateEvent,
	PullRequestTimelineItemsItemTypeUnpinnedEvent,
	PullRequestTimelineItemsItemTypeUnsubscribedEvent,
}

func (e PullRequestTimelineItemsItemType) IsValid() bool {
	switch e {
	case PullRequestTimelineItemsItemTypePullRequestCommit, PullRequestTimelineItemsItemTypePullRequestCommitCommentThread, PullRequestTimelineItemsItemTypePullRequestReview, PullRequestTimelineItemsItemTypePullRequestReviewThread, PullRequestTimelineItemsItemTypePullRequestRevisionMarker, PullRequestTimelineItemsItemTypeAutomaticBaseChangeFailedEvent, PullRequestTimelineItemsItemTypeAutomaticBaseChangeSucceededEvent, PullRequestTimelineItemsItemTypeBaseRefChangedEvent, PullRequestTimelineItemsItemTypeBaseRefForcePushedEvent, PullRequestTimelineItemsItemTypeBaseRefDeletedEvent, PullRequestTimelineItemsItemTypeDeployedEvent, PullRequestTimelineItemsItemTypeDeploymentEnvironmentChangedEvent, PullRequestTimelineItemsItemTypeHeadRefDeletedEvent, PullRequestTimelineItemsItemTypeHeadRefForcePushedEvent, PullRequestTimelineItemsItemTypeHeadRefRestoredEvent, PullRequestTimelineItemsItemTypeMergedEvent, PullRequestTimelineItemsItemTypeReviewDismissedEvent, PullRequestTimelineItemsItemTypeReviewRequestedEvent, PullRequestTimelineItemsItemTypeReviewRequestRemovedEvent, PullRequestTimelineItemsItemTypeReadyForReviewEvent, PullRequestTimelineItemsItemTypeConvertToDraftEvent, PullRequestTimelineItemsItemTypeIssueComment, PullRequestTimelineItemsItemTypeCrossReferencedEvent, PullRequestTimelineItemsItemTypeAddedToProjectEvent, PullRequestTimelineItemsItemTypeAssignedEvent, PullRequestTimelineItemsItemTypeClosedEvent, PullRequestTimelineItemsItemTypeCommentDeletedEvent, PullRequestTimelineItemsItemTypeConnectedEvent, PullRequestTimelineItemsItemTypeConvertedNoteToIssueEvent, PullRequestTimelineItemsItemTypeDemilestonedEvent, PullRequestTimelineItemsItemTypeDisconnectedEvent, PullRequestTimelineItemsItemTypeLabeledEvent, PullRequestTimelineItemsItemTypeLockedEvent, PullRequestTimelineItemsItemTypeMarkedAsDuplicateEvent, PullRequestTimelineItemsItemTypeMentionedEvent, PullRequestTimelineItemsItemTypeMilestonedEvent, PullRequestTimelineItemsItemTypeMovedColumnsInProjectEvent, PullRequestTimelineItemsItemTypePinnedEvent, PullRequestTimelineItemsItemTypeReferencedEvent, PullRequestTimelineItemsItemTypeRemovedFromProjectEvent, PullRequestTimelineItemsItemTypeRenamedTitleEvent, PullRequestTimelineItemsItemTypeReopenedEvent, PullRequestTimelineItemsItemTypeSubscribedEvent, PullRequestTimelineItemsItemTypeTransferredEvent, PullRequestTimelineItemsItemTypeUnassignedEvent, PullRequestTimelineItemsItemTypeUnlabeledEvent, PullRequestTimelineItemsItemTypeUnlockedEvent, PullRequestTimelineItemsItemTypeUserBlockedEvent, PullRequestTimelineItemsItemTypeUnmarkedAsDuplicateEvent, PullRequestTimelineItemsItemTypeUnpinnedEvent, PullRequestTimelineItemsItemTypeUnsubscribedEvent:
		return true
	}
	return false
}

func (e PullRequestTimelineItemsItemType) String() string {
	return string(e)
}

func (e *PullRequestTimelineItemsItemType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PullRequestTimelineItemsItemType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PullRequestTimelineItemsItemType", str)
	}
	return nil
}

func (e PullRequestTimelineItemsItemType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The possible target states when updating a pull request.
type PullRequestUpdateState string

const (
	// A pull request that is still open.
	PullRequestUpdateStateOpen PullRequestUpdateState = "OPEN"
	// A pull request that has been closed without being merged.
	PullRequestUpdateStateClosed PullRequestUpdateState = "CLOSED"
)

var AllPullRequestUpdateState = []PullRequestUpdateState{
	PullRequestUpdateStateOpen,
	PullRequestUpdateStateClosed,
}

func (e PullRequestUpdateState) IsValid() bool {
	switch e {
	case PullRequestUpdateStateOpen, PullRequestUpdateStateClosed:
		return true
	}
	return false
}

func (e PullRequestUpdateState) String() string {
	return string(e)
}

func (e *PullRequestUpdateState) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PullRequestUpdateState(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PullRequestUpdateState", str)
	}
	return nil
}

func (e PullRequestUpdateState) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Emojis that can be attached to Issues, Pull Requests and Comments.
type ReactionContent string

const (
	// Represents the `:+1:` emoji.
	ReactionContentThumbsUp ReactionContent = "THUMBS_UP"
	// Represents the `:-1:` emoji.
	ReactionContentThumbsDown ReactionContent = "THUMBS_DOWN"
	// Represents the `:laugh:` emoji.
	ReactionContentLaugh ReactionContent = "LAUGH"
	// Represents the `:hooray:` emoji.
	ReactionContentHooray ReactionContent = "HOORAY"
	// Represents the `:confused:` emoji.
	ReactionContentConfused ReactionContent = "CONFUSED"
	// Represents the `:heart:` emoji.
	ReactionContentHeart ReactionContent = "HEART"
	// Represents the `:rocket:` emoji.
	ReactionContentRocket ReactionContent = "ROCKET"
	// Represents the `:eyes:` emoji.
	ReactionContentEyes ReactionContent = "EYES"
)

var AllReactionContent = []ReactionContent{
	ReactionContentThumbsUp,
	ReactionContentThumbsDown,
	ReactionContentLaugh,
	ReactionContentHooray,
	ReactionContentConfused,
	ReactionContentHeart,
	ReactionContentRocket,
	ReactionContentEyes,
}

func (e ReactionContent) IsValid() bool {
	switch e {
	case ReactionContentThumbsUp, ReactionContentThumbsDown, ReactionContentLaugh, ReactionContentHooray, ReactionContentConfused, ReactionContentHeart, ReactionContentRocket, ReactionContentEyes:
		return true
	}
	return false
}

func (e ReactionContent) String() string {
	return string(e)
}

func (e *ReactionContent) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ReactionContent(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ReactionContent", str)
	}
	return nil
}

func (e ReactionContent) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// A list of fields that reactions can be ordered by.
type ReactionOrderField string

const (
	// Allows ordering a list of reactions by when they were created.
	ReactionOrderFieldCreatedAt ReactionOrderField = "CREATED_AT"
)

var AllReactionOrderField = []ReactionOrderField{
	ReactionOrderFieldCreatedAt,
}

func (e ReactionOrderField) IsValid() bool {
	switch e {
	case ReactionOrderFieldCreatedAt:
		return true
	}
	return false
}

func (e ReactionOrderField) String() string {
	return string(e)
}

func (e *ReactionOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ReactionOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ReactionOrderField", str)
	}
	return nil
}

func (e ReactionOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Properties by which ref connections can be ordered.
type RefOrderField string

const (
	// Order refs by underlying commit date if the ref prefix is refs/tags/
	RefOrderFieldTagCommitDate RefOrderField = "TAG_COMMIT_DATE"
	// Order refs by their alphanumeric name
	RefOrderFieldAlphabetical RefOrderField = "ALPHABETICAL"
)

var AllRefOrderField = []RefOrderField{
	RefOrderFieldTagCommitDate,
	RefOrderFieldAlphabetical,
}

func (e RefOrderField) IsValid() bool {
	switch e {
	case RefOrderFieldTagCommitDate, RefOrderFieldAlphabetical:
		return true
	}
	return false
}

func (e RefOrderField) String() string {
	return string(e)
}

func (e *RefOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RefOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RefOrderField", str)
	}
	return nil
}

func (e RefOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Properties by which release connections can be ordered.
type ReleaseOrderField string

const (
	// Order releases by creation time
	ReleaseOrderFieldCreatedAt ReleaseOrderField = "CREATED_AT"
	// Order releases alphabetically by name
	ReleaseOrderFieldName ReleaseOrderField = "NAME"
)

var AllReleaseOrderField = []ReleaseOrderField{
	ReleaseOrderFieldCreatedAt,
	ReleaseOrderFieldName,
}

func (e ReleaseOrderField) IsValid() bool {
	switch e {
	case ReleaseOrderFieldCreatedAt, ReleaseOrderFieldName:
		return true
	}
	return false
}

func (e ReleaseOrderField) String() string {
	return string(e)
}

func (e *ReleaseOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ReleaseOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ReleaseOrderField", str)
	}
	return nil
}

func (e ReleaseOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The privacy of a repository
type RepoAccessAuditEntryVisibility string

const (
	// The repository is visible only to users in the same business.
	RepoAccessAuditEntryVisibilityInternal RepoAccessAuditEntryVisibility = "INTERNAL"
	// The repository is visible only to those with explicit access.
	RepoAccessAuditEntryVisibilityPrivate RepoAccessAuditEntryVisibility = "PRIVATE"
	// The repository is visible to everyone.
	RepoAccessAuditEntryVisibilityPublic RepoAccessAuditEntryVisibility = "PUBLIC"
)

var AllRepoAccessAuditEntryVisibility = []RepoAccessAuditEntryVisibility{
	RepoAccessAuditEntryVisibilityInternal,
	RepoAccessAuditEntryVisibilityPrivate,
	RepoAccessAuditEntryVisibilityPublic,
}

func (e RepoAccessAuditEntryVisibility) IsValid() bool {
	switch e {
	case RepoAccessAuditEntryVisibilityInternal, RepoAccessAuditEntryVisibilityPrivate, RepoAccessAuditEntryVisibilityPublic:
		return true
	}
	return false
}

func (e RepoAccessAuditEntryVisibility) String() string {
	return string(e)
}

func (e *RepoAccessAuditEntryVisibility) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RepoAccessAuditEntryVisibility(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RepoAccessAuditEntryVisibility", str)
	}
	return nil
}

func (e RepoAccessAuditEntryVisibility) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The privacy of a repository
type RepoAddMemberAuditEntryVisibility string

const (
	// The repository is visible only to users in the same business.
	RepoAddMemberAuditEntryVisibilityInternal RepoAddMemberAuditEntryVisibility = "INTERNAL"
	// The repository is visible only to those with explicit access.
	RepoAddMemberAuditEntryVisibilityPrivate RepoAddMemberAuditEntryVisibility = "PRIVATE"
	// The repository is visible to everyone.
	RepoAddMemberAuditEntryVisibilityPublic RepoAddMemberAuditEntryVisibility = "PUBLIC"
)

var AllRepoAddMemberAuditEntryVisibility = []RepoAddMemberAuditEntryVisibility{
	RepoAddMemberAuditEntryVisibilityInternal,
	RepoAddMemberAuditEntryVisibilityPrivate,
	RepoAddMemberAuditEntryVisibilityPublic,
}

func (e RepoAddMemberAuditEntryVisibility) IsValid() bool {
	switch e {
	case RepoAddMemberAuditEntryVisibilityInternal, RepoAddMemberAuditEntryVisibilityPrivate, RepoAddMemberAuditEntryVisibilityPublic:
		return true
	}
	return false
}

func (e RepoAddMemberAuditEntryVisibility) String() string {
	return string(e)
}

func (e *RepoAddMemberAuditEntryVisibility) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RepoAddMemberAuditEntryVisibility(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RepoAddMemberAuditEntryVisibility", str)
	}
	return nil
}

func (e RepoAddMemberAuditEntryVisibility) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The privacy of a repository
type RepoArchivedAuditEntryVisibility string

const (
	// The repository is visible only to users in the same business.
	RepoArchivedAuditEntryVisibilityInternal RepoArchivedAuditEntryVisibility = "INTERNAL"
	// The repository is visible only to those with explicit access.
	RepoArchivedAuditEntryVisibilityPrivate RepoArchivedAuditEntryVisibility = "PRIVATE"
	// The repository is visible to everyone.
	RepoArchivedAuditEntryVisibilityPublic RepoArchivedAuditEntryVisibility = "PUBLIC"
)

var AllRepoArchivedAuditEntryVisibility = []RepoArchivedAuditEntryVisibility{
	RepoArchivedAuditEntryVisibilityInternal,
	RepoArchivedAuditEntryVisibilityPrivate,
	RepoArchivedAuditEntryVisibilityPublic,
}

func (e RepoArchivedAuditEntryVisibility) IsValid() bool {
	switch e {
	case RepoArchivedAuditEntryVisibilityInternal, RepoArchivedAuditEntryVisibilityPrivate, RepoArchivedAuditEntryVisibilityPublic:
		return true
	}
	return false
}

func (e RepoArchivedAuditEntryVisibility) String() string {
	return string(e)
}

func (e *RepoArchivedAuditEntryVisibility) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RepoArchivedAuditEntryVisibility(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RepoArchivedAuditEntryVisibility", str)
	}
	return nil
}

func (e RepoArchivedAuditEntryVisibility) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The merge options available for pull requests to this repository.
type RepoChangeMergeSettingAuditEntryMergeType string

const (
	// The pull request is added to the base branch in a merge commit.
	RepoChangeMergeSettingAuditEntryMergeTypeMerge RepoChangeMergeSettingAuditEntryMergeType = "MERGE"
	// Commits from the pull request are added onto the base branch individually without a merge commit.
	RepoChangeMergeSettingAuditEntryMergeTypeRebase RepoChangeMergeSettingAuditEntryMergeType = "REBASE"
	// The pull request's commits are squashed into a single commit before they are merged to the base branch.
	RepoChangeMergeSettingAuditEntryMergeTypeSquash RepoChangeMergeSettingAuditEntryMergeType = "SQUASH"
)

var AllRepoChangeMergeSettingAuditEntryMergeType = []RepoChangeMergeSettingAuditEntryMergeType{
	RepoChangeMergeSettingAuditEntryMergeTypeMerge,
	RepoChangeMergeSettingAuditEntryMergeTypeRebase,
	RepoChangeMergeSettingAuditEntryMergeTypeSquash,
}

func (e RepoChangeMergeSettingAuditEntryMergeType) IsValid() bool {
	switch e {
	case RepoChangeMergeSettingAuditEntryMergeTypeMerge, RepoChangeMergeSettingAuditEntryMergeTypeRebase, RepoChangeMergeSettingAuditEntryMergeTypeSquash:
		return true
	}
	return false
}

func (e RepoChangeMergeSettingAuditEntryMergeType) String() string {
	return string(e)
}

func (e *RepoChangeMergeSettingAuditEntryMergeType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RepoChangeMergeSettingAuditEntryMergeType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RepoChangeMergeSettingAuditEntryMergeType", str)
	}
	return nil
}

func (e RepoChangeMergeSettingAuditEntryMergeType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The privacy of a repository
type RepoCreateAuditEntryVisibility string

const (
	// The repository is visible only to users in the same business.
	RepoCreateAuditEntryVisibilityInternal RepoCreateAuditEntryVisibility = "INTERNAL"
	// The repository is visible only to those with explicit access.
	RepoCreateAuditEntryVisibilityPrivate RepoCreateAuditEntryVisibility = "PRIVATE"
	// The repository is visible to everyone.
	RepoCreateAuditEntryVisibilityPublic RepoCreateAuditEntryVisibility = "PUBLIC"
)

var AllRepoCreateAuditEntryVisibility = []RepoCreateAuditEntryVisibility{
	RepoCreateAuditEntryVisibilityInternal,
	RepoCreateAuditEntryVisibilityPrivate,
	RepoCreateAuditEntryVisibilityPublic,
}

func (e RepoCreateAuditEntryVisibility) IsValid() bool {
	switch e {
	case RepoCreateAuditEntryVisibilityInternal, RepoCreateAuditEntryVisibilityPrivate, RepoCreateAuditEntryVisibilityPublic:
		return true
	}
	return false
}

func (e RepoCreateAuditEntryVisibility) String() string {
	return string(e)
}

func (e *RepoCreateAuditEntryVisibility) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RepoCreateAuditEntryVisibility(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RepoCreateAuditEntryVisibility", str)
	}
	return nil
}

func (e RepoCreateAuditEntryVisibility) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The privacy of a repository
type RepoDestroyAuditEntryVisibility string

const (
	// The repository is visible only to users in the same business.
	RepoDestroyAuditEntryVisibilityInternal RepoDestroyAuditEntryVisibility = "INTERNAL"
	// The repository is visible only to those with explicit access.
	RepoDestroyAuditEntryVisibilityPrivate RepoDestroyAuditEntryVisibility = "PRIVATE"
	// The repository is visible to everyone.
	RepoDestroyAuditEntryVisibilityPublic RepoDestroyAuditEntryVisibility = "PUBLIC"
)

var AllRepoDestroyAuditEntryVisibility = []RepoDestroyAuditEntryVisibility{
	RepoDestroyAuditEntryVisibilityInternal,
	RepoDestroyAuditEntryVisibilityPrivate,
	RepoDestroyAuditEntryVisibilityPublic,
}

func (e RepoDestroyAuditEntryVisibility) IsValid() bool {
	switch e {
	case RepoDestroyAuditEntryVisibilityInternal, RepoDestroyAuditEntryVisibilityPrivate, RepoDestroyAuditEntryVisibilityPublic:
		return true
	}
	return false
}

func (e RepoDestroyAuditEntryVisibility) String() string {
	return string(e)
}

func (e *RepoDestroyAuditEntryVisibility) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RepoDestroyAuditEntryVisibility(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RepoDestroyAuditEntryVisibility", str)
	}
	return nil
}

func (e RepoDestroyAuditEntryVisibility) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The privacy of a repository
type RepoRemoveMemberAuditEntryVisibility string

const (
	// The repository is visible only to users in the same business.
	RepoRemoveMemberAuditEntryVisibilityInternal RepoRemoveMemberAuditEntryVisibility = "INTERNAL"
	// The repository is visible only to those with explicit access.
	RepoRemoveMemberAuditEntryVisibilityPrivate RepoRemoveMemberAuditEntryVisibility = "PRIVATE"
	// The repository is visible to everyone.
	RepoRemoveMemberAuditEntryVisibilityPublic RepoRemoveMemberAuditEntryVisibility = "PUBLIC"
)

var AllRepoRemoveMemberAuditEntryVisibility = []RepoRemoveMemberAuditEntryVisibility{
	RepoRemoveMemberAuditEntryVisibilityInternal,
	RepoRemoveMemberAuditEntryVisibilityPrivate,
	RepoRemoveMemberAuditEntryVisibilityPublic,
}

func (e RepoRemoveMemberAuditEntryVisibility) IsValid() bool {
	switch e {
	case RepoRemoveMemberAuditEntryVisibilityInternal, RepoRemoveMemberAuditEntryVisibilityPrivate, RepoRemoveMemberAuditEntryVisibilityPublic:
		return true
	}
	return false
}

func (e RepoRemoveMemberAuditEntryVisibility) String() string {
	return string(e)
}

func (e *RepoRemoveMemberAuditEntryVisibility) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RepoRemoveMemberAuditEntryVisibility(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RepoRemoveMemberAuditEntryVisibility", str)
	}
	return nil
}

func (e RepoRemoveMemberAuditEntryVisibility) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The reasons a piece of content can be reported or minimized.
type ReportedContentClassifiers string

const (
	// A spammy piece of content
	ReportedContentClassifiersSpam ReportedContentClassifiers = "SPAM"
	// An abusive or harassing piece of content
	ReportedContentClassifiersAbuse ReportedContentClassifiers = "ABUSE"
	// An irrelevant piece of content
	ReportedContentClassifiersOffTopic ReportedContentClassifiers = "OFF_TOPIC"
	// An outdated piece of content
	ReportedContentClassifiersOutdated ReportedContentClassifiers = "OUTDATED"
	// A duplicated piece of content
	ReportedContentClassifiersDuplicate ReportedContentClassifiers = "DUPLICATE"
	// The content has been resolved
	ReportedContentClassifiersResolved ReportedContentClassifiers = "RESOLVED"
)

var AllReportedContentClassifiers = []ReportedContentClassifiers{
	ReportedContentClassifiersSpam,
	ReportedContentClassifiersAbuse,
	ReportedContentClassifiersOffTopic,
	ReportedContentClassifiersOutdated,
	ReportedContentClassifiersDuplicate,
	ReportedContentClassifiersResolved,
}

func (e ReportedContentClassifiers) IsValid() bool {
	switch e {
	case ReportedContentClassifiersSpam, ReportedContentClassifiersAbuse, ReportedContentClassifiersOffTopic, ReportedContentClassifiersOutdated, ReportedContentClassifiersDuplicate, ReportedContentClassifiersResolved:
		return true
	}
	return false
}

func (e ReportedContentClassifiers) String() string {
	return string(e)
}

func (e *ReportedContentClassifiers) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ReportedContentClassifiers(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ReportedContentClassifiers", str)
	}
	return nil
}

func (e ReportedContentClassifiers) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The affiliation of a user to a repository
type RepositoryAffiliation string

const (
	// Repositories that are owned by the authenticated user.
	RepositoryAffiliationOwner RepositoryAffiliation = "OWNER"
	// Repositories that the user has been added to as a collaborator.
	RepositoryAffiliationCollaborator RepositoryAffiliation = "COLLABORATOR"
	// Repositories that the user has access to through being a member of an organization. This includes every repository on every team that the user is on.
	RepositoryAffiliationOrganizationMember RepositoryAffiliation = "ORGANIZATION_MEMBER"
)

var AllRepositoryAffiliation = []RepositoryAffiliation{
	RepositoryAffiliationOwner,
	RepositoryAffiliationCollaborator,
	RepositoryAffiliationOrganizationMember,
}

func (e RepositoryAffiliation) IsValid() bool {
	switch e {
	case RepositoryAffiliationOwner, RepositoryAffiliationCollaborator, RepositoryAffiliationOrganizationMember:
		return true
	}
	return false
}

func (e RepositoryAffiliation) String() string {
	return string(e)
}

func (e *RepositoryAffiliation) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RepositoryAffiliation(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RepositoryAffiliation", str)
	}
	return nil
}

func (e RepositoryAffiliation) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The reason a repository is listed as 'contributed'.
type RepositoryContributionType string

const (
	// Created a commit
	RepositoryContributionTypeCommit RepositoryContributionType = "COMMIT"
	// Created an issue
	RepositoryContributionTypeIssue RepositoryContributionType = "ISSUE"
	// Created a pull request
	RepositoryContributionTypePullRequest RepositoryContributionType = "PULL_REQUEST"
	// Created the repository
	RepositoryContributionTypeRepository RepositoryContributionType = "REPOSITORY"
	// Reviewed a pull request
	RepositoryContributionTypePullRequestReview RepositoryContributionType = "PULL_REQUEST_REVIEW"
)

var AllRepositoryContributionType = []RepositoryContributionType{
	RepositoryContributionTypeCommit,
	RepositoryContributionTypeIssue,
	RepositoryContributionTypePullRequest,
	RepositoryContributionTypeRepository,
	RepositoryContributionTypePullRequestReview,
}

func (e RepositoryContributionType) IsValid() bool {
	switch e {
	case RepositoryContributionTypeCommit, RepositoryContributionTypeIssue, RepositoryContributionTypePullRequest, RepositoryContributionTypeRepository, RepositoryContributionTypePullRequestReview:
		return true
	}
	return false
}

func (e RepositoryContributionType) String() string {
	return string(e)
}

func (e *RepositoryContributionType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RepositoryContributionType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RepositoryContributionType", str)
	}
	return nil
}

func (e RepositoryContributionType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Properties by which repository invitation connections can be ordered.
type RepositoryInvitationOrderField string

const (
	// Order repository invitations by creation time
	RepositoryInvitationOrderFieldCreatedAt RepositoryInvitationOrderField = "CREATED_AT"
	// Order repository invitations by invitee login
	RepositoryInvitationOrderFieldInviteeLogin RepositoryInvitationOrderField = "INVITEE_LOGIN"
)

var AllRepositoryInvitationOrderField = []RepositoryInvitationOrderField{
	RepositoryInvitationOrderFieldCreatedAt,
	RepositoryInvitationOrderFieldInviteeLogin,
}

func (e RepositoryInvitationOrderField) IsValid() bool {
	switch e {
	case RepositoryInvitationOrderFieldCreatedAt, RepositoryInvitationOrderFieldInviteeLogin:
		return true
	}
	return false
}

func (e RepositoryInvitationOrderField) String() string {
	return string(e)
}

func (e *RepositoryInvitationOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RepositoryInvitationOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RepositoryInvitationOrderField", str)
	}
	return nil
}

func (e RepositoryInvitationOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The possible reasons a given repository could be in a locked state.
type RepositoryLockReason string

const (
	// The repository is locked due to a move.
	RepositoryLockReasonMoving RepositoryLockReason = "MOVING"
	// The repository is locked due to a billing related reason.
	RepositoryLockReasonBilling RepositoryLockReason = "BILLING"
	// The repository is locked due to a rename.
	RepositoryLockReasonRename RepositoryLockReason = "RENAME"
	// The repository is locked due to a migration.
	RepositoryLockReasonMigrating RepositoryLockReason = "MIGRATING"
)

var AllRepositoryLockReason = []RepositoryLockReason{
	RepositoryLockReasonMoving,
	RepositoryLockReasonBilling,
	RepositoryLockReasonRename,
	RepositoryLockReasonMigrating,
}

func (e RepositoryLockReason) IsValid() bool {
	switch e {
	case RepositoryLockReasonMoving, RepositoryLockReasonBilling, RepositoryLockReasonRename, RepositoryLockReasonMigrating:
		return true
	}
	return false
}

func (e RepositoryLockReason) String() string {
	return string(e)
}

func (e *RepositoryLockReason) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RepositoryLockReason(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RepositoryLockReason", str)
	}
	return nil
}

func (e RepositoryLockReason) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Properties by which repository connections can be ordered.
type RepositoryOrderField string

const (
	// Order repositories by creation time
	RepositoryOrderFieldCreatedAt RepositoryOrderField = "CREATED_AT"
	// Order repositories by update time
	RepositoryOrderFieldUpdatedAt RepositoryOrderField = "UPDATED_AT"
	// Order repositories by push time
	RepositoryOrderFieldPushedAt RepositoryOrderField = "PUSHED_AT"
	// Order repositories by name
	RepositoryOrderFieldName RepositoryOrderField = "NAME"
	// Order repositories by number of stargazers
	RepositoryOrderFieldStargazers RepositoryOrderField = "STARGAZERS"
)

var AllRepositoryOrderField = []RepositoryOrderField{
	RepositoryOrderFieldCreatedAt,
	RepositoryOrderFieldUpdatedAt,
	RepositoryOrderFieldPushedAt,
	RepositoryOrderFieldName,
	RepositoryOrderFieldStargazers,
}

func (e RepositoryOrderField) IsValid() bool {
	switch e {
	case RepositoryOrderFieldCreatedAt, RepositoryOrderFieldUpdatedAt, RepositoryOrderFieldPushedAt, RepositoryOrderFieldName, RepositoryOrderFieldStargazers:
		return true
	}
	return false
}

func (e RepositoryOrderField) String() string {
	return string(e)
}

func (e *RepositoryOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RepositoryOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RepositoryOrderField", str)
	}
	return nil
}

func (e RepositoryOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The access level to a repository
type RepositoryPermission string

const (
	// Can read, clone, and push to this repository. Can also manage issues, pull requests, and repository settings, including adding collaborators
	RepositoryPermissionAdmin RepositoryPermission = "ADMIN"
	// Can read, clone, and push to this repository. They can also manage issues, pull requests, and some repository settings
	RepositoryPermissionMaintain RepositoryPermission = "MAINTAIN"
	// Can read, clone, and push to this repository. Can also manage issues and pull requests
	RepositoryPermissionWrite RepositoryPermission = "WRITE"
	// Can read and clone this repository. Can also manage issues and pull requests
	RepositoryPermissionTriage RepositoryPermission = "TRIAGE"
	// Can read and clone this repository. Can also open and comment on issues and pull requests
	RepositoryPermissionRead RepositoryPermission = "READ"
)

var AllRepositoryPermission = []RepositoryPermission{
	RepositoryPermissionAdmin,
	RepositoryPermissionMaintain,
	RepositoryPermissionWrite,
	RepositoryPermissionTriage,
	RepositoryPermissionRead,
}

func (e RepositoryPermission) IsValid() bool {
	switch e {
	case RepositoryPermissionAdmin, RepositoryPermissionMaintain, RepositoryPermissionWrite, RepositoryPermissionTriage, RepositoryPermissionRead:
		return true
	}
	return false
}

func (e RepositoryPermission) String() string {
	return string(e)
}

func (e *RepositoryPermission) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RepositoryPermission(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RepositoryPermission", str)
	}
	return nil
}

func (e RepositoryPermission) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The privacy of a repository
type RepositoryPrivacy string

const (
	// Public
	RepositoryPrivacyPublic RepositoryPrivacy = "PUBLIC"
	// Private
	RepositoryPrivacyPrivate RepositoryPrivacy = "PRIVATE"
)

var AllRepositoryPrivacy = []RepositoryPrivacy{
	RepositoryPrivacyPublic,
	RepositoryPrivacyPrivate,
}

func (e RepositoryPrivacy) IsValid() bool {
	switch e {
	case RepositoryPrivacyPublic, RepositoryPrivacyPrivate:
		return true
	}
	return false
}

func (e RepositoryPrivacy) String() string {
	return string(e)
}

func (e *RepositoryPrivacy) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RepositoryPrivacy(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RepositoryPrivacy", str)
	}
	return nil
}

func (e RepositoryPrivacy) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The repository's visibility level.
type RepositoryVisibility string

const (
	// The repository is visible only to those with explicit access.
	RepositoryVisibilityPrivate RepositoryVisibility = "PRIVATE"
	// The repository is visible to everyone.
	RepositoryVisibilityPublic RepositoryVisibility = "PUBLIC"
	// The repository is visible only to users in the same business.
	RepositoryVisibilityInternal RepositoryVisibility = "INTERNAL"
)

var AllRepositoryVisibility = []RepositoryVisibility{
	RepositoryVisibilityPrivate,
	RepositoryVisibilityPublic,
	RepositoryVisibilityInternal,
}

func (e RepositoryVisibility) IsValid() bool {
	switch e {
	case RepositoryVisibilityPrivate, RepositoryVisibilityPublic, RepositoryVisibilityInternal:
		return true
	}
	return false
}

func (e RepositoryVisibility) String() string {
	return string(e)
}

func (e *RepositoryVisibility) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RepositoryVisibility(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RepositoryVisibility", str)
	}
	return nil
}

func (e RepositoryVisibility) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The possible states that can be requested when creating a check run.
type RequestableCheckStatusState string

const (
	// The check suite or run has been queued.
	RequestableCheckStatusStateQueued RequestableCheckStatusState = "QUEUED"
	// The check suite or run is in progress.
	RequestableCheckStatusStateInProgress RequestableCheckStatusState = "IN_PROGRESS"
	// The check suite or run has been completed.
	RequestableCheckStatusStateCompleted RequestableCheckStatusState = "COMPLETED"
)

var AllRequestableCheckStatusState = []RequestableCheckStatusState{
	RequestableCheckStatusStateQueued,
	RequestableCheckStatusStateInProgress,
	RequestableCheckStatusStateCompleted,
}

func (e RequestableCheckStatusState) IsValid() bool {
	switch e {
	case RequestableCheckStatusStateQueued, RequestableCheckStatusStateInProgress, RequestableCheckStatusStateCompleted:
		return true
	}
	return false
}

func (e RequestableCheckStatusState) String() string {
	return string(e)
}

func (e *RequestableCheckStatusState) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RequestableCheckStatusState(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RequestableCheckStatusState", str)
	}
	return nil
}

func (e RequestableCheckStatusState) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The possible digest algorithms used to sign SAML requests for an identity provider.
type SamlDigestAlgorithm string

const (
	// SHA1
	SamlDigestAlgorithmSha1 SamlDigestAlgorithm = "SHA1"
	// SHA256
	SamlDigestAlgorithmSha256 SamlDigestAlgorithm = "SHA256"
	// SHA384
	SamlDigestAlgorithmSha384 SamlDigestAlgorithm = "SHA384"
	// SHA512
	SamlDigestAlgorithmSha512 SamlDigestAlgorithm = "SHA512"
)

var AllSamlDigestAlgorithm = []SamlDigestAlgorithm{
	SamlDigestAlgorithmSha1,
	SamlDigestAlgorithmSha256,
	SamlDigestAlgorithmSha384,
	SamlDigestAlgorithmSha512,
}

func (e SamlDigestAlgorithm) IsValid() bool {
	switch e {
	case SamlDigestAlgorithmSha1, SamlDigestAlgorithmSha256, SamlDigestAlgorithmSha384, SamlDigestAlgorithmSha512:
		return true
	}
	return false
}

func (e SamlDigestAlgorithm) String() string {
	return string(e)
}

func (e *SamlDigestAlgorithm) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SamlDigestAlgorithm(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SamlDigestAlgorithm", str)
	}
	return nil
}

func (e SamlDigestAlgorithm) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The possible signature algorithms used to sign SAML requests for a Identity Provider.
type SamlSignatureAlgorithm string

const (
	// RSA-SHA1
	SamlSignatureAlgorithmRsaSha1 SamlSignatureAlgorithm = "RSA_SHA1"
	// RSA-SHA256
	SamlSignatureAlgorithmRsaSha256 SamlSignatureAlgorithm = "RSA_SHA256"
	// RSA-SHA384
	SamlSignatureAlgorithmRsaSha384 SamlSignatureAlgorithm = "RSA_SHA384"
	// RSA-SHA512
	SamlSignatureAlgorithmRsaSha512 SamlSignatureAlgorithm = "RSA_SHA512"
)

var AllSamlSignatureAlgorithm = []SamlSignatureAlgorithm{
	SamlSignatureAlgorithmRsaSha1,
	SamlSignatureAlgorithmRsaSha256,
	SamlSignatureAlgorithmRsaSha384,
	SamlSignatureAlgorithmRsaSha512,
}

func (e SamlSignatureAlgorithm) IsValid() bool {
	switch e {
	case SamlSignatureAlgorithmRsaSha1, SamlSignatureAlgorithmRsaSha256, SamlSignatureAlgorithmRsaSha384, SamlSignatureAlgorithmRsaSha512:
		return true
	}
	return false
}

func (e SamlSignatureAlgorithm) String() string {
	return string(e)
}

func (e *SamlSignatureAlgorithm) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SamlSignatureAlgorithm(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SamlSignatureAlgorithm", str)
	}
	return nil
}

func (e SamlSignatureAlgorithm) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Properties by which saved reply connections can be ordered.
type SavedReplyOrderField string

const (
	// Order saved reply by when they were updated.
	SavedReplyOrderFieldUpdatedAt SavedReplyOrderField = "UPDATED_AT"
)

var AllSavedReplyOrderField = []SavedReplyOrderField{
	SavedReplyOrderFieldUpdatedAt,
}

func (e SavedReplyOrderField) IsValid() bool {
	switch e {
	case SavedReplyOrderFieldUpdatedAt:
		return true
	}
	return false
}

func (e SavedReplyOrderField) String() string {
	return string(e)
}

func (e *SavedReplyOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SavedReplyOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SavedReplyOrderField", str)
	}
	return nil
}

func (e SavedReplyOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Represents the individual results of a search.
type SearchType string

const (
	// Returns results matching issues in repositories.
	SearchTypeIssue SearchType = "ISSUE"
	// Returns results matching repositories.
	SearchTypeRepository SearchType = "REPOSITORY"
	// Returns results matching users and organizations on GitHub.
	SearchTypeUser SearchType = "USER"
)

var AllSearchType = []SearchType{
	SearchTypeIssue,
	SearchTypeRepository,
	SearchTypeUser,
}

func (e SearchType) IsValid() bool {
	switch e {
	case SearchTypeIssue, SearchTypeRepository, SearchTypeUser:
		return true
	}
	return false
}

func (e SearchType) String() string {
	return string(e)
}

func (e *SearchType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SearchType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SearchType", str)
	}
	return nil
}

func (e SearchType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The possible ecosystems of a security vulnerability's package.
type SecurityAdvisoryEcosystem string

const (
	// Ruby gems hosted at RubyGems.org
	SecurityAdvisoryEcosystemRubygems SecurityAdvisoryEcosystem = "RUBYGEMS"
	// JavaScript packages hosted at npmjs.com
	SecurityAdvisoryEcosystemNpm SecurityAdvisoryEcosystem = "NPM"
	// Python packages hosted at PyPI.org
	SecurityAdvisoryEcosystemPip SecurityAdvisoryEcosystem = "PIP"
	// Java artifacts hosted at the Maven central repository
	SecurityAdvisoryEcosystemMaven SecurityAdvisoryEcosystem = "MAVEN"
	// .NET packages hosted at the NuGet Gallery
	SecurityAdvisoryEcosystemNuget SecurityAdvisoryEcosystem = "NUGET"
	// PHP packages hosted at packagist.org
	SecurityAdvisoryEcosystemComposer SecurityAdvisoryEcosystem = "COMPOSER"
)

var AllSecurityAdvisoryEcosystem = []SecurityAdvisoryEcosystem{
	SecurityAdvisoryEcosystemRubygems,
	SecurityAdvisoryEcosystemNpm,
	SecurityAdvisoryEcosystemPip,
	SecurityAdvisoryEcosystemMaven,
	SecurityAdvisoryEcosystemNuget,
	SecurityAdvisoryEcosystemComposer,
}

func (e SecurityAdvisoryEcosystem) IsValid() bool {
	switch e {
	case SecurityAdvisoryEcosystemRubygems, SecurityAdvisoryEcosystemNpm, SecurityAdvisoryEcosystemPip, SecurityAdvisoryEcosystemMaven, SecurityAdvisoryEcosystemNuget, SecurityAdvisoryEcosystemComposer:
		return true
	}
	return false
}

func (e SecurityAdvisoryEcosystem) String() string {
	return string(e)
}

func (e *SecurityAdvisoryEcosystem) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SecurityAdvisoryEcosystem(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SecurityAdvisoryEcosystem", str)
	}
	return nil
}

func (e SecurityAdvisoryEcosystem) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Identifier formats available for advisories.
type SecurityAdvisoryIdentifierType string

const (
	// Common Vulnerabilities and Exposures Identifier.
	SecurityAdvisoryIdentifierTypeCve SecurityAdvisoryIdentifierType = "CVE"
	// GitHub Security Advisory ID.
	SecurityAdvisoryIdentifierTypeGhsa SecurityAdvisoryIdentifierType = "GHSA"
)

var AllSecurityAdvisoryIdentifierType = []SecurityAdvisoryIdentifierType{
	SecurityAdvisoryIdentifierTypeCve,
	SecurityAdvisoryIdentifierTypeGhsa,
}

func (e SecurityAdvisoryIdentifierType) IsValid() bool {
	switch e {
	case SecurityAdvisoryIdentifierTypeCve, SecurityAdvisoryIdentifierTypeGhsa:
		return true
	}
	return false
}

func (e SecurityAdvisoryIdentifierType) String() string {
	return string(e)
}

func (e *SecurityAdvisoryIdentifierType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SecurityAdvisoryIdentifierType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SecurityAdvisoryIdentifierType", str)
	}
	return nil
}

func (e SecurityAdvisoryIdentifierType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Properties by which security advisory connections can be ordered.
type SecurityAdvisoryOrderField string

const (
	// Order advisories by publication time
	SecurityAdvisoryOrderFieldPublishedAt SecurityAdvisoryOrderField = "PUBLISHED_AT"
	// Order advisories by update time
	SecurityAdvisoryOrderFieldUpdatedAt SecurityAdvisoryOrderField = "UPDATED_AT"
)

var AllSecurityAdvisoryOrderField = []SecurityAdvisoryOrderField{
	SecurityAdvisoryOrderFieldPublishedAt,
	SecurityAdvisoryOrderFieldUpdatedAt,
}

func (e SecurityAdvisoryOrderField) IsValid() bool {
	switch e {
	case SecurityAdvisoryOrderFieldPublishedAt, SecurityAdvisoryOrderFieldUpdatedAt:
		return true
	}
	return false
}

func (e SecurityAdvisoryOrderField) String() string {
	return string(e)
}

func (e *SecurityAdvisoryOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SecurityAdvisoryOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SecurityAdvisoryOrderField", str)
	}
	return nil
}

func (e SecurityAdvisoryOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Severity of the vulnerability.
type SecurityAdvisorySeverity string

const (
	// Low.
	SecurityAdvisorySeverityLow SecurityAdvisorySeverity = "LOW"
	// Moderate.
	SecurityAdvisorySeverityModerate SecurityAdvisorySeverity = "MODERATE"
	// High.
	SecurityAdvisorySeverityHigh SecurityAdvisorySeverity = "HIGH"
	// Critical.
	SecurityAdvisorySeverityCritical SecurityAdvisorySeverity = "CRITICAL"
)

var AllSecurityAdvisorySeverity = []SecurityAdvisorySeverity{
	SecurityAdvisorySeverityLow,
	SecurityAdvisorySeverityModerate,
	SecurityAdvisorySeverityHigh,
	SecurityAdvisorySeverityCritical,
}

func (e SecurityAdvisorySeverity) IsValid() bool {
	switch e {
	case SecurityAdvisorySeverityLow, SecurityAdvisorySeverityModerate, SecurityAdvisorySeverityHigh, SecurityAdvisorySeverityCritical:
		return true
	}
	return false
}

func (e SecurityAdvisorySeverity) String() string {
	return string(e)
}

func (e *SecurityAdvisorySeverity) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SecurityAdvisorySeverity(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SecurityAdvisorySeverity", str)
	}
	return nil
}

func (e SecurityAdvisorySeverity) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Properties by which security vulnerability connections can be ordered.
type SecurityVulnerabilityOrderField string

const (
	// Order vulnerability by update time
	SecurityVulnerabilityOrderFieldUpdatedAt SecurityVulnerabilityOrderField = "UPDATED_AT"
)

var AllSecurityVulnerabilityOrderField = []SecurityVulnerabilityOrderField{
	SecurityVulnerabilityOrderFieldUpdatedAt,
}

func (e SecurityVulnerabilityOrderField) IsValid() bool {
	switch e {
	case SecurityVulnerabilityOrderFieldUpdatedAt:
		return true
	}
	return false
}

func (e SecurityVulnerabilityOrderField) String() string {
	return string(e)
}

func (e *SecurityVulnerabilityOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SecurityVulnerabilityOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SecurityVulnerabilityOrderField", str)
	}
	return nil
}

func (e SecurityVulnerabilityOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Properties by which Sponsors tiers connections can be ordered.
type SponsorsTierOrderField string

const (
	// Order tiers by creation time.
	SponsorsTierOrderFieldCreatedAt SponsorsTierOrderField = "CREATED_AT"
	// Order tiers by their monthly price in cents
	SponsorsTierOrderFieldMonthlyPriceInCents SponsorsTierOrderField = "MONTHLY_PRICE_IN_CENTS"
)

var AllSponsorsTierOrderField = []SponsorsTierOrderField{
	SponsorsTierOrderFieldCreatedAt,
	SponsorsTierOrderFieldMonthlyPriceInCents,
}

func (e SponsorsTierOrderField) IsValid() bool {
	switch e {
	case SponsorsTierOrderFieldCreatedAt, SponsorsTierOrderFieldMonthlyPriceInCents:
		return true
	}
	return false
}

func (e SponsorsTierOrderField) String() string {
	return string(e)
}

func (e *SponsorsTierOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SponsorsTierOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SponsorsTierOrderField", str)
	}
	return nil
}

func (e SponsorsTierOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Properties by which sponsorship connections can be ordered.
type SponsorshipOrderField string

const (
	// Order sponsorship by creation time.
	SponsorshipOrderFieldCreatedAt SponsorshipOrderField = "CREATED_AT"
)

var AllSponsorshipOrderField = []SponsorshipOrderField{
	SponsorshipOrderFieldCreatedAt,
}

func (e SponsorshipOrderField) IsValid() bool {
	switch e {
	case SponsorshipOrderFieldCreatedAt:
		return true
	}
	return false
}

func (e SponsorshipOrderField) String() string {
	return string(e)
}

func (e *SponsorshipOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SponsorshipOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SponsorshipOrderField", str)
	}
	return nil
}

func (e SponsorshipOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The privacy of a sponsorship
type SponsorshipPrivacy string

const (
	// Public
	SponsorshipPrivacyPublic SponsorshipPrivacy = "PUBLIC"
	// Private
	SponsorshipPrivacyPrivate SponsorshipPrivacy = "PRIVATE"
)

var AllSponsorshipPrivacy = []SponsorshipPrivacy{
	SponsorshipPrivacyPublic,
	SponsorshipPrivacyPrivate,
}

func (e SponsorshipPrivacy) IsValid() bool {
	switch e {
	case SponsorshipPrivacyPublic, SponsorshipPrivacyPrivate:
		return true
	}
	return false
}

func (e SponsorshipPrivacy) String() string {
	return string(e)
}

func (e *SponsorshipPrivacy) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SponsorshipPrivacy(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SponsorshipPrivacy", str)
	}
	return nil
}

func (e SponsorshipPrivacy) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Properties by which star connections can be ordered.
type StarOrderField string

const (
	// Allows ordering a list of stars by when they were created.
	StarOrderFieldStarredAt StarOrderField = "STARRED_AT"
)

var AllStarOrderField = []StarOrderField{
	StarOrderFieldStarredAt,
}

func (e StarOrderField) IsValid() bool {
	switch e {
	case StarOrderFieldStarredAt:
		return true
	}
	return false
}

func (e StarOrderField) String() string {
	return string(e)
}

func (e *StarOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = StarOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid StarOrderField", str)
	}
	return nil
}

func (e StarOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The possible commit status states.
type StatusState string

const (
	// Status is expected.
	StatusStateExpected StatusState = "EXPECTED"
	// Status is errored.
	StatusStateError StatusState = "ERROR"
	// Status is failing.
	StatusStateFailure StatusState = "FAILURE"
	// Status is pending.
	StatusStatePending StatusState = "PENDING"
	// Status is successful.
	StatusStateSuccess StatusState = "SUCCESS"
)

var AllStatusState = []StatusState{
	StatusStateExpected,
	StatusStateError,
	StatusStateFailure,
	StatusStatePending,
	StatusStateSuccess,
}

func (e StatusState) IsValid() bool {
	switch e {
	case StatusStateExpected, StatusStateError, StatusStateFailure, StatusStatePending, StatusStateSuccess:
		return true
	}
	return false
}

func (e StatusState) String() string {
	return string(e)
}

func (e *StatusState) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = StatusState(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid StatusState", str)
	}
	return nil
}

func (e StatusState) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The possible states of a subscription.
type SubscriptionState string

const (
	// The User is only notified when participating or @mentioned.
	SubscriptionStateUnsubscribed SubscriptionState = "UNSUBSCRIBED"
	// The User is notified of all conversations.
	SubscriptionStateSubscribed SubscriptionState = "SUBSCRIBED"
	// The User is never notified.
	SubscriptionStateIgnored SubscriptionState = "IGNORED"
)

var AllSubscriptionState = []SubscriptionState{
	SubscriptionStateUnsubscribed,
	SubscriptionStateSubscribed,
	SubscriptionStateIgnored,
}

func (e SubscriptionState) IsValid() bool {
	switch e {
	case SubscriptionStateUnsubscribed, SubscriptionStateSubscribed, SubscriptionStateIgnored:
		return true
	}
	return false
}

func (e SubscriptionState) String() string {
	return string(e)
}

func (e *SubscriptionState) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SubscriptionState(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SubscriptionState", str)
	}
	return nil
}

func (e SubscriptionState) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Properties by which team discussion comment connections can be ordered.
type TeamDiscussionCommentOrderField string

const (
	// Allows sequential ordering of team discussion comments (which is equivalent to chronological ordering).
	TeamDiscussionCommentOrderFieldNumber TeamDiscussionCommentOrderField = "NUMBER"
)

var AllTeamDiscussionCommentOrderField = []TeamDiscussionCommentOrderField{
	TeamDiscussionCommentOrderFieldNumber,
}

func (e TeamDiscussionCommentOrderField) IsValid() bool {
	switch e {
	case TeamDiscussionCommentOrderFieldNumber:
		return true
	}
	return false
}

func (e TeamDiscussionCommentOrderField) String() string {
	return string(e)
}

func (e *TeamDiscussionCommentOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TeamDiscussionCommentOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TeamDiscussionCommentOrderField", str)
	}
	return nil
}

func (e TeamDiscussionCommentOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Properties by which team discussion connections can be ordered.
type TeamDiscussionOrderField string

const (
	// Allows chronological ordering of team discussions.
	TeamDiscussionOrderFieldCreatedAt TeamDiscussionOrderField = "CREATED_AT"
)

var AllTeamDiscussionOrderField = []TeamDiscussionOrderField{
	TeamDiscussionOrderFieldCreatedAt,
}

func (e TeamDiscussionOrderField) IsValid() bool {
	switch e {
	case TeamDiscussionOrderFieldCreatedAt:
		return true
	}
	return false
}

func (e TeamDiscussionOrderField) String() string {
	return string(e)
}

func (e *TeamDiscussionOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TeamDiscussionOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TeamDiscussionOrderField", str)
	}
	return nil
}

func (e TeamDiscussionOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Properties by which team member connections can be ordered.
type TeamMemberOrderField string

const (
	// Order team members by login
	TeamMemberOrderFieldLogin TeamMemberOrderField = "LOGIN"
	// Order team members by creation time
	TeamMemberOrderFieldCreatedAt TeamMemberOrderField = "CREATED_AT"
)

var AllTeamMemberOrderField = []TeamMemberOrderField{
	TeamMemberOrderFieldLogin,
	TeamMemberOrderFieldCreatedAt,
}

func (e TeamMemberOrderField) IsValid() bool {
	switch e {
	case TeamMemberOrderFieldLogin, TeamMemberOrderFieldCreatedAt:
		return true
	}
	return false
}

func (e TeamMemberOrderField) String() string {
	return string(e)
}

func (e *TeamMemberOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TeamMemberOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TeamMemberOrderField", str)
	}
	return nil
}

func (e TeamMemberOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The possible team member roles; either 'maintainer' or 'member'.
type TeamMemberRole string

const (
	// A team maintainer has permission to add and remove team members.
	TeamMemberRoleMaintainer TeamMemberRole = "MAINTAINER"
	// A team member has no administrative permissions on the team.
	TeamMemberRoleMember TeamMemberRole = "MEMBER"
)

var AllTeamMemberRole = []TeamMemberRole{
	TeamMemberRoleMaintainer,
	TeamMemberRoleMember,
}

func (e TeamMemberRole) IsValid() bool {
	switch e {
	case TeamMemberRoleMaintainer, TeamMemberRoleMember:
		return true
	}
	return false
}

func (e TeamMemberRole) String() string {
	return string(e)
}

func (e *TeamMemberRole) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TeamMemberRole(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TeamMemberRole", str)
	}
	return nil
}

func (e TeamMemberRole) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Defines which types of team members are included in the returned list. Can be one of IMMEDIATE, CHILD_TEAM or ALL.
type TeamMembershipType string

const (
	// Includes only immediate members of the team.
	TeamMembershipTypeImmediate TeamMembershipType = "IMMEDIATE"
	// Includes only child team members for the team.
	TeamMembershipTypeChildTeam TeamMembershipType = "CHILD_TEAM"
	// Includes immediate and child team members for the team.
	TeamMembershipTypeAll TeamMembershipType = "ALL"
)

var AllTeamMembershipType = []TeamMembershipType{
	TeamMembershipTypeImmediate,
	TeamMembershipTypeChildTeam,
	TeamMembershipTypeAll,
}

func (e TeamMembershipType) IsValid() bool {
	switch e {
	case TeamMembershipTypeImmediate, TeamMembershipTypeChildTeam, TeamMembershipTypeAll:
		return true
	}
	return false
}

func (e TeamMembershipType) String() string {
	return string(e)
}

func (e *TeamMembershipType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TeamMembershipType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TeamMembershipType", str)
	}
	return nil
}

func (e TeamMembershipType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Properties by which team connections can be ordered.
type TeamOrderField string

const (
	// Allows ordering a list of teams by name.
	TeamOrderFieldName TeamOrderField = "NAME"
)

var AllTeamOrderField = []TeamOrderField{
	TeamOrderFieldName,
}

func (e TeamOrderField) IsValid() bool {
	switch e {
	case TeamOrderFieldName:
		return true
	}
	return false
}

func (e TeamOrderField) String() string {
	return string(e)
}

func (e *TeamOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TeamOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TeamOrderField", str)
	}
	return nil
}

func (e TeamOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The possible team privacy values.
type TeamPrivacy string

const (
	// A secret team can only be seen by its members.
	TeamPrivacySecret TeamPrivacy = "SECRET"
	// A visible team can be seen and @mentioned by every member of the organization.
	TeamPrivacyVisible TeamPrivacy = "VISIBLE"
)

var AllTeamPrivacy = []TeamPrivacy{
	TeamPrivacySecret,
	TeamPrivacyVisible,
}

func (e TeamPrivacy) IsValid() bool {
	switch e {
	case TeamPrivacySecret, TeamPrivacyVisible:
		return true
	}
	return false
}

func (e TeamPrivacy) String() string {
	return string(e)
}

func (e *TeamPrivacy) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TeamPrivacy(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TeamPrivacy", str)
	}
	return nil
}

func (e TeamPrivacy) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Properties by which team repository connections can be ordered.
type TeamRepositoryOrderField string

const (
	// Order repositories by creation time
	TeamRepositoryOrderFieldCreatedAt TeamRepositoryOrderField = "CREATED_AT"
	// Order repositories by update time
	TeamRepositoryOrderFieldUpdatedAt TeamRepositoryOrderField = "UPDATED_AT"
	// Order repositories by push time
	TeamRepositoryOrderFieldPushedAt TeamRepositoryOrderField = "PUSHED_AT"
	// Order repositories by name
	TeamRepositoryOrderFieldName TeamRepositoryOrderField = "NAME"
	// Order repositories by permission
	TeamRepositoryOrderFieldPermission TeamRepositoryOrderField = "PERMISSION"
	// Order repositories by number of stargazers
	TeamRepositoryOrderFieldStargazers TeamRepositoryOrderField = "STARGAZERS"
)

var AllTeamRepositoryOrderField = []TeamRepositoryOrderField{
	TeamRepositoryOrderFieldCreatedAt,
	TeamRepositoryOrderFieldUpdatedAt,
	TeamRepositoryOrderFieldPushedAt,
	TeamRepositoryOrderFieldName,
	TeamRepositoryOrderFieldPermission,
	TeamRepositoryOrderFieldStargazers,
}

func (e TeamRepositoryOrderField) IsValid() bool {
	switch e {
	case TeamRepositoryOrderFieldCreatedAt, TeamRepositoryOrderFieldUpdatedAt, TeamRepositoryOrderFieldPushedAt, TeamRepositoryOrderFieldName, TeamRepositoryOrderFieldPermission, TeamRepositoryOrderFieldStargazers:
		return true
	}
	return false
}

func (e TeamRepositoryOrderField) String() string {
	return string(e)
}

func (e *TeamRepositoryOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TeamRepositoryOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TeamRepositoryOrderField", str)
	}
	return nil
}

func (e TeamRepositoryOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The role of a user on a team.
type TeamRole string

const (
	// User has admin rights on the team.
	TeamRoleAdmin TeamRole = "ADMIN"
	// User is a member of the team.
	TeamRoleMember TeamRole = "MEMBER"
)

var AllTeamRole = []TeamRole{
	TeamRoleAdmin,
	TeamRoleMember,
}

func (e TeamRole) IsValid() bool {
	switch e {
	case TeamRoleAdmin, TeamRoleMember:
		return true
	}
	return false
}

func (e TeamRole) String() string {
	return string(e)
}

func (e *TeamRole) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TeamRole(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TeamRole", str)
	}
	return nil
}

func (e TeamRole) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Reason that the suggested topic is declined.
type TopicSuggestionDeclineReason string

const (
	// The suggested topic is not relevant to the repository.
	TopicSuggestionDeclineReasonNotRelevant TopicSuggestionDeclineReason = "NOT_RELEVANT"
	// The suggested topic is too specific for the repository (e.g. #ruby-on-rails-version-4-2-1).
	TopicSuggestionDeclineReasonTooSpecific TopicSuggestionDeclineReason = "TOO_SPECIFIC"
	// The viewer does not like the suggested topic.
	TopicSuggestionDeclineReasonPersonalPreference TopicSuggestionDeclineReason = "PERSONAL_PREFERENCE"
	// The suggested topic is too general for the repository.
	TopicSuggestionDeclineReasonTooGeneral TopicSuggestionDeclineReason = "TOO_GENERAL"
)

var AllTopicSuggestionDeclineReason = []TopicSuggestionDeclineReason{
	TopicSuggestionDeclineReasonNotRelevant,
	TopicSuggestionDeclineReasonTooSpecific,
	TopicSuggestionDeclineReasonPersonalPreference,
	TopicSuggestionDeclineReasonTooGeneral,
}

func (e TopicSuggestionDeclineReason) IsValid() bool {
	switch e {
	case TopicSuggestionDeclineReasonNotRelevant, TopicSuggestionDeclineReasonTooSpecific, TopicSuggestionDeclineReasonPersonalPreference, TopicSuggestionDeclineReasonTooGeneral:
		return true
	}
	return false
}

func (e TopicSuggestionDeclineReason) String() string {
	return string(e)
}

func (e *TopicSuggestionDeclineReason) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TopicSuggestionDeclineReason(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TopicSuggestionDeclineReason", str)
	}
	return nil
}

func (e TopicSuggestionDeclineReason) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The possible durations that a user can be blocked for.
type UserBlockDuration string

const (
	// The user was blocked for 1 day
	UserBlockDurationOneDay UserBlockDuration = "ONE_DAY"
	// The user was blocked for 3 days
	UserBlockDurationThreeDays UserBlockDuration = "THREE_DAYS"
	// The user was blocked for 7 days
	UserBlockDurationOneWeek UserBlockDuration = "ONE_WEEK"
	// The user was blocked for 30 days
	UserBlockDurationOneMonth UserBlockDuration = "ONE_MONTH"
	// The user was blocked permanently
	UserBlockDurationPermanent UserBlockDuration = "PERMANENT"
)

var AllUserBlockDuration = []UserBlockDuration{
	UserBlockDurationOneDay,
	UserBlockDurationThreeDays,
	UserBlockDurationOneWeek,
	UserBlockDurationOneMonth,
	UserBlockDurationPermanent,
}

func (e UserBlockDuration) IsValid() bool {
	switch e {
	case UserBlockDurationOneDay, UserBlockDurationThreeDays, UserBlockDurationOneWeek, UserBlockDurationOneMonth, UserBlockDurationPermanent:
		return true
	}
	return false
}

func (e UserBlockDuration) String() string {
	return string(e)
}

func (e *UserBlockDuration) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserBlockDuration(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserBlockDuration", str)
	}
	return nil
}

func (e UserBlockDuration) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Properties by which user status connections can be ordered.
type UserStatusOrderField string

const (
	// Order user statuses by when they were updated.
	UserStatusOrderFieldUpdatedAt UserStatusOrderField = "UPDATED_AT"
)

var AllUserStatusOrderField = []UserStatusOrderField{
	UserStatusOrderFieldUpdatedAt,
}

func (e UserStatusOrderField) IsValid() bool {
	switch e {
	case UserStatusOrderFieldUpdatedAt:
		return true
	}
	return false
}

func (e UserStatusOrderField) String() string {
	return string(e)
}

func (e *UserStatusOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserStatusOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserStatusOrderField", str)
	}
	return nil
}

func (e UserStatusOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
