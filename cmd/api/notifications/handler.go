package notifications

import (
	"net/http"
)

func EventsLog(w http.ResponseWriter, r *http.Request) {
	// userID, err := middleware.GetContextPersonID(r.Context())
	// if err != nil {
	// 	w.WriteHeader(401)
	// 	return
	// }
	// typeName := r.URL.Query().Get("type")
	// if typeName == "" {
	// 	w.WriteHeader(400)
	// 	return
	// }
	// service := cmd.NotificationServiceInstance()
	// switch typeName {
	// case "user":
	// 	eventsLog, err := service.UserLog(r.Context(), userID)
	// 	if err != nil {
	// 		w.WriteHeader(500)
	// 		return
	// 	}
	// 	if err := json.NewEncoder(w).Encode(eventsLog); err != nil {
	// 		log.Println("could not write result to response: ", err)
	// 	}
	// default:
	// 	w.WriteHeader(400)
	// 	return
}

// }
// func ProvideNotifications(w http.ResponseWriter, r *http.Request) {
// 	userID, err := middleware.GetContextPersonID(r.Context())
// 	if err != nil {
// 		w.WriteHeader(401)
// 		return
// 	}
// 	typeName := r.URL.Query().Get("type")
// 	if typeName == "" {
// 		w.WriteHeader(400)
// 		return
// 	}
// 	eventsRepository := repositories.NewEventRepository(datasource.EventSchema())
// 	service := application.NewNotificationService(
// 		eventsRepository,
// 	)
// 	switch typeName {
// 	case "enterprise":
// 		enterpriseID := r.URL.Query().Get("id")
// 		if enterpriseID == "" {
// 			w.WriteHeader(400)
// 			return
// 		}
// 		events, err := service.EnterpriseNotifications(r.Context(), enterpriseID)
// 		if err != nil {
// 			log.Println("could not get enterprise notifications: ", err)
// 			w.WriteHeader(500)
// 			return
// 		}
// 		if err := json.NewEncoder(w).Encode(events); err != nil {
// 			log.Println("could not write result to response: ", err)
// 		}
// 	case "user":
// 		notifications, err := service.UserNotifications(r.Context(), userID)
// 		if err != nil {
// 			w.WriteHeader(500)
// 			return
// 		}
// 		if err := json.NewEncoder(w).Encode(notifications); err != nil {
// 			log.Println("could not write result to response: ", err)
// 		}
// 	case "employee":
// 		return
// 	default:
// 		w.WriteHeader(400)
// 		return
// 	}
// }

// func Notification(w http.ResponseWriter, r *http.Request) {
// 	_, err := middleware.GetContextPersonID(r.Context())
// 	if err != nil {
// 		w.WriteHeader(401)
// 		return
// 	}
// 	id := r.URL.Query().Get("id")
// 	publisher := domain.DomainEventPublisherInstance()
// 	publisher.Reset()
// 	eventProccesor := application.NewEventProcessor()
// 	publisher.Subscribe(eventProccesor.Subscriber())
// 	eventsRepository := repositories.NewEventRepository(datasource.EventSchema())
// 	service := application.NewNotificationService(
// 		eventsRepository,
// 	)
// 	notification, err := service.Notification(r.Context(), id)
// 	if err != nil {
// 		w.WriteHeader(500)
// 		return
// 	}
// 	if _, err := w.Write(notification); err != nil {
// 		w.WriteHeader(500)
// 		return
// 	}
// }
