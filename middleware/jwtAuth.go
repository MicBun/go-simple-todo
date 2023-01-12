package middleware

import (
	"github.com/MicBun/go-simple-todo/utils/jwtAuth"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := jwtAuth.TokenValid(c)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
			return
		}
		c.Next()
	}
}

//func JwtAndClockInMiddleware() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		err := jwtAuth.TokenValid(c)
//		if err != nil {
//			c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
//			return
//		}
//
//		userID, err := jwtAuth.ExtractTokenID(c)
//		if err != nil {
//			c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
//			return
//		}
//
//		attendance := models.Attendance{UserID: userID}
//		attendanceResult, err := attendance.GetAttendanceByDate(c, time.Now().Format("2006-01-02"))
//		if err == nil {
//			if attendanceResult.UpdatedAt == attendanceResult.CreatedAt {
//				c.Next()
//			}
//		}
//
//		c.AbortWithStatusJSON(400, gin.H{"error": "user have not clocked in yet or already clocked out"})
//	}
//}
