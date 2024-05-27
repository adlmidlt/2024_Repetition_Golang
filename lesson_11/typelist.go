package lesson_11

import "headfirstgo/lesson_11/gadget"

func playList(device gadget.Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func StartPlayList() {
	mixtape := []string{"Linkin Park", "System of Down", "Amatory", "Stray of kids"}
	var player gadget.Player = gadget.TapePlayer{}
	playList(player, mixtape)
	player = gadget.TapeRecorder{}
	playList(player, mixtape)
}
