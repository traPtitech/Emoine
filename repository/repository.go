package repository

type Repository interface {
	StateRepository
	PresentationRepository
	ReactionRepository
	CommentRepository
	ReviewRepository
}
