package library

import "testing"

func TestOps(t *testing.T) {
	mm := NewMusicManager()
	if mm == nil {
		t.Error(" NewMusicManger failed,")
	}
	if mm.Len() != 0 {
		t.Error("New failed, not empty")
	}
	m0 := &MusicEntry{"1", "My Heart Will Go On", "Celion Dion", "Pop", "http://qbox.me/24501234", "MP3"}
	mm.Add(m0)

	if mm.Len() != 1 {
		t.Error("Music Add faild")
	}
	m := mm.Find(m0.Name)
	if m == nil {
		t.Error("Music Find() faild.")
	}
	if m.Id != m0.Id || m.Artist !=m.Artist || m.Name != m0.Name ||
		m.Genre != m0.Genre || m.Source != m0.Source || m.Type !=m0.Type{
			t.Error("MusicManager.Find() faild . Found it mismatch.")
	}
	m,err :=mm.Get(0)
	if m == nil{
		t.Error("MusicManger.Get() faild .", err)
	}
	m = mm.Remove(0)
	if m ==nil || mm.Len() !=0{
		t.Error("MusicManager.Remove faild", err)
	}
}
