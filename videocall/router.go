package videocall

import (
	"github.com/gofiber/fiber/v2"
)

func AddRouterForVoiceCalls(router fiber.Router) {
	// add routers for api
	router.Get("/log", nil)
	router.Post("/log", nil)
}

// func main() {
// 	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()
// 	log.SetOutput(file)
// 	router := gin.Default()

// 	// sender to channel of track
// 	peerConnectionMap := make(map[string]chan *webrtc.Track)

// 	m := webrtc.MediaEngine{}

// 	// Setup the codecs you want to use.
// 	// Only support VP8(video compression), this makes our proxying code simpler
// 	m.RegisterCodec(webrtc.NewRTPVP8Codec(webrtc.DefaultPayloadTypeVP8, 90000))

// 	api := webrtc.NewAPI(webrtc.WithMediaEngine(m))

// 	peerConnectionConfig := webrtc.Configuration{
// 		ICEServers: []webrtc.ICEServer{
// 			{
// 				URLs: []string{"stun:stun.l.google.com:19302"},
// 			},
// 		},
// 	}

// 	router.POST("/webrtc/sdp/m/:meetingId/c/:userID/p/:peerId/s/:isSender", func(c *gin.Context) {
// 		isSender, _ := strconv.ParseBool(c.Param("isSender"))
// 		userID := c.Param("userID")
// 		peerID := c.Param("peerId")

// 		var session Sdp
// 		if err := c.ShouldBindJSON(&session); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		offer := webrtc.SessionDescription{}
// 		Decode(session.Sdp, &offer)

// 		// Create a new RTCPeerConnection
// 		// this is the gist of webrtc, generates and process SDP
// 		peerConnection, err := api.NewPeerConnection(peerConnectionConfig)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		if !isSender {
// 			recieveTrack(peerConnection, peerConnectionMap, peerID)
// 		} else {
// 			createTrack(peerConnection, peerConnectionMap, userID)
// 		}
// 		// Set the SessionDescription of remote peer
// 		peerConnection.SetRemoteDescription(offer)

// 		// Create answer
// 		answer, err := peerConnection.CreateAnswer(nil)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		// Sets the LocalDescription, and starts our UDP listeners
// 		err = peerConnection.SetLocalDescription(answer)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		c.JSON(http.StatusOK, Sdp{Sdp: Encode(answer)})
// 	})

// 	router.Run(":8080")
// }
