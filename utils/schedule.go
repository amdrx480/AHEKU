package utils

// "backend-golang/businesses/admin"
// "fmt"
// "log"
// "time"
// "github.com/go-co-op/gocron"
// "backend-golang/drivers/mysql/admin"

// func ScheduleNotification(reminder admin.ReminderPurchaseOrder, packagingOfficer admin.PackagingOfficer) {
// 	// Format pesan yang akan dikirim
// 	message := fmt.Sprintf("Reminder for Packaging: Pre-order for product %s needs to be prepared. Quantity: %d. Customer Name: %s.",
// 		reminder.CartItem.Stocks.StockName, reminder.CartItem.Quantity, reminder.CartItem.Customers.CustomerName)

// 	// Hitung durasi sampai ReminderTime
// 	duration := time.Until(reminder.ReminderTime)

// 	// Jalankan goroutine untuk menunggu sampai ReminderTime dan kirim pesan
// 	go func() {
// 		time.Sleep(duration)

// 		// Kirim pesan WhatsApp ke Packaging Officer
// 		err := SendWhatsAppMessage(packagingOfficer.OfficerPhone, message)
// 		if err != nil {
// 			log.Printf("Failed to send WhatsApp message: %v\n", err)
// 		} else {
// 			log.Printf("WhatsApp message sent to %s: %s\n", packagingOfficer.OfficerPhone, message)
// 		}
// 	}()
// }

// func ScheduleNotification(reminder admin.ReminderPurchaseOrderDomain, packagingOfficer admin.PackagingOfficerDomain) {
// 	// Format pesan yang akan dikirim
// 	message := fmt.Sprintf("Reminder for Packaging: Pre-order for product %s needs to be prepared. Quantity: %d. Customer Name: %s.",
// 		reminder.CartItem.StockName, reminder.CartItem.Quantity, reminder.CartItem.CustomerName)

// 	// Hitung durasi sampai ReminderTime
// 	duration := time.Until(reminder.ReminderTime)

// 		// Log nama dan nomor dari Packaging Officer
// 	log.Printf("Scheduling notification for Packaging Officer: %s, Phone: %s", packagingOfficer.OfficerName, packagingOfficer.OfficerPhone)

// 	// Jalankan goroutine untuk menunggu sampai ReminderTime dan kirim pesan
// 	go func() {
// 		time.Sleep(duration)

// 		// Kirim pesan WhatsApp ke Packaging Officer
// 		err := SendWhatsAppMessage(packagingOfficer.OfficerPhone, message)
// 		if err != nil {
// 			log.Printf("Failed to send WhatsApp message: %v\n", err)
// 		} else {
// 			log.Printf("WhatsApp message sent to %s: %s\n", packagingOfficer.OfficerPhone, message)
// 		}
// 	}()
// }

// func ScheduleNotification(reminderPurchaseOrderDomain admin.ReminderPurchaseOrderDomain, admin admin.AdminDomain) {
// 	s := gocron.NewScheduler(time.UTC)

// 	s.Every(1).Day().At(reminderPurchaseOrderDomain.ReminderTime.Format("15:04")).Do(func() {
// 		message := fmt.Sprintf("Reminder for Packaging: Pre-order for product %s needs to be prepared. Quantity: %d. Customer Name: %s.",
// 			reminderPurchaseOrderDomain.CartItem.StockName, reminderPurchaseOrderDomain.CartItem.Quantity, reminderPurchaseOrderDomain.CartItem.CustomerName)

// 		err := SendWhatsAppMessage(admin.Phone, reminderPurchaseOrderDomain.Packaging.OfficerPhone, message)
// 		if err != nil {
// 			fmt.Printf("Failed to send WhatsApp message: %v\n", err)
// 		}
// 	})

// 	s.StartAsync()
// }
