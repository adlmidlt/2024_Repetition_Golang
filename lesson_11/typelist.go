package lesson_11

import "headfirstgo/lesson_11/gadget"

func playList(device gadget.TapePlayer, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func StartPlayList() {
	player := gadget.TapePlayer{}
	//player := gadget.TapeRecorder{}
	mixtape := []string{"Linkin Park", "System of Down", "Amatory", "Stray of kids"}
	playList(player, mixtape)
}
