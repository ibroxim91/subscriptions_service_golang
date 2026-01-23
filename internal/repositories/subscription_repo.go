package repositories

import (
	"gorm.io/gorm"
	"subscriptions_service_golang/internal/models"
)


type SubscriptionRepository interface {
    Create(sub *models.Subscription) error
    GetByID(id uint) (*models.Subscription, error)
    List(filter map[string]interface{}) ([]models.Subscription, error)
    Update(sub *models.Subscription) error
    Delete(id uint) error
    TotalPrice(userID string, serviceName string, from, to string) (int, error)
}

type subscriptionRepository struct {
    db *gorm.DB
}

func NewSubscriptionRepository(db *gorm.DB) SubscriptionRepository {
    return &subscriptionRepository{db: db}
}

func (r *subscriptionRepository) Create(sub *models.Subscription) error {
    return r.db.Create(sub).Error
}

func (r *subscriptionRepository) GetByID(id uint) (*models.Subscription, error) {
    var sub models.Subscription
    if err := r.db.First(&sub, id).Error; err != nil {
        return nil, err
    }
    return &sub, nil
}

func (r *subscriptionRepository) List(filter map[string]interface{}) ([]models.Subscription, error) {
    var subs []models.Subscription
    query := r.db.Model(&models.Subscription{})
    for k, v := range filter {
        query = query.Where(k+" = ?", v)
    }
    if err := query.Find(&subs).Error; err != nil {
        return nil, err
    }
    return subs, nil
}


func (r *subscriptionRepository) Update(sub *models.Subscription) error {
    return r.db.Save(sub).Error
}


func (r *subscriptionRepository) Delete(id uint) error {
    return r.db.Delete(&models.Subscription{}, id).Error
}


func (r *subscriptionRepository) TotalPrice(userID string, serviceName string, from, to string) (int, error) {
    var total int
    query := r.db.Model(&models.Subscription{}).Select("SUM(price)")

    if userID != "" {
        query = query.Where("user_id = ?", userID)
    }
    if serviceName != "" {
        query = query.Where("service_name = ?", serviceName)
    }
    if from != "" {
        query = query.Where("start_date >= ?", from)
    }
    if to != "" {
        query = query.Where("start_date <= ?", to)
    }

    if err := query.Scan(&total).Error; err != nil {
        return 0, err
    }
    return total, nil
}
