(enum
  :package assignment
  :type ApprovalStatus
  :code-type ApprovalStatusCode
  :string-type ApprovalStatusString
  :const-prefix Approval
  :all AllApprovalStatuses
  :code-all AllApprovalStatusCodes
  (value ApprovalPending "pending")
  (value ApprovalApproved "approved")
  (value ApprovalDenied "denied")
  (value ApprovalTimedOut "timed_out"))
