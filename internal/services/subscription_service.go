package services

import (
	"subscriptions_service_golang/internal/models"
	"subscriptions_service_golang/internal/repositories"
	"time"
)

type SubscriptionService interface {
	Create(sub models.Subscription) (*models.Subscription, error)
	GetByID(id uint) (*models.Subscription, error)
	List(userID string, serviceName string, from, to *time.Time) ([]models.Subscription, error)
	Update(sub models.Subscription) (*models.Subscription, error)
	Delete(id uint) error
	TotalPrice(userID string, serviceName string, from, to *time.Time) (int, error)
}

type subscriptionService struct {
	repo repositories.SubscriptionRepository
}

func NewSubscriptionService(repo repositories.SubscriptionRepository) SubscriptionService {
	return &subscriptionService{repo: repo}
}

// Create yangi subscription yaratadi
func (s *subscriptionService) Create(sub models.Subscription) (*models.Subscription, error) {
	if err := s.repo.Create(&sub); err != nil {
		return nil, err
	}
	return &sub, nil
}

// GetByID subscriptionni ID bo‘yicha qaytaradi
func (s *subscriptionService) GetByID(id uint) (*models.Subscription, error) {
	return s.repo.GetByID(id)
}

// List subscriptionlarni filter bilan qaytaradi
func (s *subscriptionService) List(userID string, serviceName string, from, to *time.Time) ([]models.Subscription, error) {
	filter := make(map[string]interface{})
	if userID != "" {
		filter["user_id"] = userID
	}
	if serviceName != "" {
		filter["service_name"] = serviceName
	}
	// vaqt filterini repositoryda ishlatish mumkin
	return s.repo.List(filter)
}

// Update subscriptionni yangilaydi
func (s *subscriptionService) Update(sub models.Subscription) (*models.Subscription, error) {
	if err := s.repo.Update(&sub); err != nil {
		return nil, err
	}
	return &sub, nil
}

// Delete subscriptionni o‘chiradi
func (s *subscriptionService) Delete(id uint) error {
	return s.repo.Delete(id)
}

// TotalPrice — foydalanuvchi va davr bo‘yicha umumiy narxni hisoblaydi
func (s *subscriptionService) TotalPrice(userID string, serviceName string, from, to *time.Time) (int, error) {
	fromStr, toStr := "", ""
	if from != nil {
		fromStr = from.Format("2006-01-02")
	}
	if to != nil {
		toStr = to.Format("2006-01-02")
	}
	return s.repo.TotalPrice(userID, serviceName, fromStr, toStr)
}
