package fileAvailability

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
	"github.com/matryer/try"
	"log"
	"os"
	"strings"
	"time"
)

type Service interface {
	GetFilesInPath(path string) ([]string, error)
	pathToMostRecentFile(dirPath, fileContains string) (bool, string, error)
	ConfirmUgandaFileAvailability()
	CreateJSONResponse() map[string]bool
}

type service struct {
}

var (
	ZimbabweStatus   bool
	BotswanaStatus   bool
	KenyaStatus      bool
	MalawiStatus     bool
	NamibiaStatus    bool
	GhanaStatus      bool
	UgandaStatus     bool
	UgandaDRStatus   bool
	ZambiaStatus     bool
	ZambiaDRStatus   bool
	ZambiaProdStatus bool
)

type File struct {
	Name         string
	Path         string
	Size         int64
	LastModified time.Time
}

func NewService() Service {
	s := &service{}
	s.ConfirmUgandaFileAvailability()
	//s.ConfirmBotswanaFileAvailability()
	//s.ConfirmGhanaFileAvailability()
	//s.ConfirmKenyaFileAvailability()
	//s.ConfirmMalawiFileAvailability()
	//s.ConfirmNamibiaFileAvailability()
	//s.ConfirmUgandaDRFileAvailability()
	//s.ConfirmZambiaFileAvailability()
	//s.ConfirmZambiaDRFileAvailability()
	//s.ConfirmZambiaProdFileAvailability()
	//s.ConfirmZimbabweFileAvailability()

	return s
}

func (s *service) schedule() {
	confirmUgandaAvailability := gocron.NewScheduler()
	confirmBotswanaAvailability := gocron.NewScheduler()
	confirmMalawiAvailability := gocron.NewScheduler()
	confirmGhanaAvailability := gocron.NewScheduler()
	confirmKenyaAvailability := gocron.NewScheduler()
	confirmNamibiaAvailability := gocron.NewScheduler()
	confirmUgandaDRAvailability := gocron.NewScheduler()
	confirmZambiaAvailability := gocron.NewScheduler()
	confirmZambiaDRAvailability := gocron.NewScheduler()
	confirmZambiaProdAvailability := gocron.NewScheduler()
	confirmZimbabweAvailability := gocron.NewScheduler()

	go func() {
		confirmUgandaAvailability.Every(1).Day().At("00:00").Do(s.ConfirmUgandaFileAvailability)
		<-confirmUgandaAvailability.Start()
	}()
	go func() {
		confirmBotswanaAvailability.Every(1).Day().At("00:00").Do(s.ConfirmBotswanaFileAvailability)
		<-confirmBotswanaAvailability.Start()
	}()
	go func() {
		confirmMalawiAvailability.Every(1).Day().At("00:00").Do(s.ConfirmMalawiFileAvailability)
		<-confirmMalawiAvailability.Start()
	}()
	go func() {
		confirmGhanaAvailability.Every(1).Day().At("00:00").Do(s.ConfirmGhanaFileAvailability)
		<-confirmGhanaAvailability.Start()
	}()
	go func() {
		confirmKenyaAvailability.Every(1).Day().At("00:00").Do(s.ConfirmKenyaFileAvailability)
		<-confirmKenyaAvailability.Start()
	}()
	go func() {
		confirmNamibiaAvailability.Every(1).Day().At("00:00").Do(s.ConfirmNamibiaFileAvailability)
		<-confirmNamibiaAvailability.Start()
	}()
	go func() {
		confirmUgandaDRAvailability.Every(1).Day().At("00:00").Do(s.ConfirmUgandaDRFileAvailability)
		<-confirmUgandaDRAvailability.Start()
	}()
	go func() {
		confirmZambiaAvailability.Every(1).Day().At("00:00").Do(s.ConfirmZambiaFileAvailability)
		<-confirmZambiaAvailability.Start()
	}()
	go func() {
		confirmZambiaDRAvailability.Every(1).Day().At("00:00").Do(s.ConfirmZambiaDRFileAvailability)
		<-confirmZambiaDRAvailability.Start()
	}()
	go func() {
		confirmZambiaProdAvailability.Every(1).Day().At("00:00").Do(s.ConfirmZambiaProdFileAvailability)
		<-confirmZambiaProdAvailability.Start()
	}()
	go func() {
		confirmZimbabweAvailability.Every(1).Day().At("00:00").Do(s.ConfirmZimbabweFileAvailability)
		<-confirmZimbabweAvailability.Start()
	}()
}

func (s *service) GetFilesInPath(path string) ([]string, error) {

	dir, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer dir.Close()

	list, _ := dir.Readdirnames(0)

	return list, nil
}

func (s *service) pathToMostRecentFile(dirPath, fileContains string) (bool, string, error) {

	fileList, err := s.GetFilesInPath(dirPath)
	if err != nil || len(fileList) == 0 {
		log.Println(fmt.Sprintf("Unable to access %v", dirPath))
	}

	currentDate := time.Now().Format("20060102")

	for _, file := range fileList {
		cont := strings.Contains(file, fileContains)
		recent := strings.Contains(file, currentDate)

		if recent == true && cont == true {
			return true, file, nil
		}
	}
	return false, "", fmt.Errorf("%v file has not arrived yet", fileContains)
}

func (s *service) ConfirmFileAvailabilityMethod(path string) error {
	fileReceived, _, err := s.pathToMostRecentFile(path, ".TXT")

	if err != nil {
		return err
	}

	switch path {
	case "/mnt/zimbabwe":
		ZimbabweStatus = fileReceived
	case "/mnt/botswana":
		BotswanaStatus = fileReceived
	case "/mnt/ghana":
		GhanaStatus = fileReceived
	case "/mnt/kenya":
		KenyaStatus = fileReceived
	case "/mnt/malawi":
		MalawiStatus = fileReceived
	case "/mnt/namibia":
		NamibiaStatus = fileReceived
	case "/mnt/uganda":
		UgandaStatus = fileReceived
	case "/mnt/ugandadr":
		UgandaDRStatus = fileReceived
	case "/mnt/zambia":
		ZambiaStatus = fileReceived
	case "/mnt/zambiadr":
		ZambiaDRStatus = fileReceived
	case "/mnt/zambiaprod":
		ZambiaProdStatus = fileReceived

	}
	return nil
}

func (s *service) CreateJSONResponse() map[string]bool {

	resp := map[string]bool{
		"ZimbabweStatus":   ZimbabweStatus,
		"BotswanaStatus":   BotswanaStatus,
		"KenyaStatus":      KenyaStatus,
		"MalawiStatus":     MalawiStatus,
		"NamibiaStatus":    NamibiaStatus,
		"GhanaStatus":      GhanaStatus,
		"UgandaStatus":     UgandaStatus,
		"UgandaDRStatus":   UgandaDRStatus,
		"ZambiaStatus":     ZambiaStatus,
		"ZambiaDRStatus":   ZambiaDRStatus,
		"ZambiaProdStatus": ZambiaProdStatus,
	}

	return resp
}

func (s *service) ConfirmZimbabweFileAvailability() {
	err := try.Do(func(attempt int) (bool, error) {
		try.MaxRetries = 240
		var err error
		err = s.ConfirmFileAvailabilityMethod("/mnt/zimbabwe")
		if err != nil {
			log.Println("Zimbabwe file not yet detected. Next attempt in 2 minutes...")
			time.Sleep(2 * time.Minute) // wait 2 minutes
		}
		return true, err
	})
	if err != nil {
		log.Println(err)
	}
}

func (s *service) ConfirmBotswanaFileAvailability() {
	err := try.Do(func(attempt int) (bool, error) {
		try.MaxRetries = 240
		var err error
		err = s.ConfirmFileAvailabilityMethod("/mnt/botswana")
		if err != nil {
			log.Println("Botswana file not yet detected. Next attempt in 2 minutes...")
			time.Sleep(2 * time.Minute) // wait 2 minutes
		}
		return true, err
	})
	if err != nil {
		log.Println(err)
	}
}

func (s *service) ConfirmGhanaFileAvailability() {
	err := try.Do(func(attempt int) (bool, error) {
		try.MaxRetries = 240
		var err error
		err = s.ConfirmFileAvailabilityMethod("/mnt/ghana")
		if err != nil {
			log.Println("Ghana file not yet detected. Next attempt in 2 minutes...")
			time.Sleep(2 * time.Minute) // wait 2 minutes
		}
		return true, err
	})
	if err != nil {
		log.Println(err)
	}
}

func (s *service) ConfirmKenyaFileAvailability() {
	err := try.Do(func(attempt int) (bool, error) {
		try.MaxRetries = 240
		var err error
		err = s.ConfirmFileAvailabilityMethod("/mnt/kenya")
		if err != nil {
			log.Println("Kenya file not yet detected. Next attempt in 2 minutes...")
			time.Sleep(2 * time.Minute) // wait 2 minutes
		}
		return true, err
	})
	if err != nil {
		log.Println(err)
	}
}

func (s *service) ConfirmMalawiFileAvailability() {
	err := try.Do(func(attempt int) (bool, error) {
		try.MaxRetries = 240
		var err error
		err = s.ConfirmFileAvailabilityMethod("/mnt/malawi")
		if err != nil {
			log.Println("Malawi file not yet detected. Next attempt in 2 minutes...")
			time.Sleep(2 * time.Minute) // wait 2 minutes
		}
		return true, err
	})
	if err != nil {
		log.Println(err)
	}
}

func (s *service) ConfirmNamibiaFileAvailability() {
	err := try.Do(func(attempt int) (bool, error) {
		try.MaxRetries = 240
		var err error
		err = s.ConfirmFileAvailabilityMethod("/mnt/namibia")
		if err != nil {
			log.Println("Namibia file not yet detected. Next attempt in 2 minutes...")
			time.Sleep(2 * time.Minute) // wait 2 minutes
		}
		return true, err
	})
	if err != nil {
		log.Println(err)
	}
}

func (s *service) ConfirmUgandaFileAvailability() {
	err := try.Do(func(attempt int) (bool, error) {
		try.MaxRetries = 240
		var err error
		err = s.ConfirmFileAvailabilityMethod("/mnt/uganda")
		if err != nil {
			log.Println("Uganda file not yet detected. Next attempt in 2 minutes...")
			time.Sleep(2 * time.Minute) // wait 2 minutes
		}
		return true, err
	})
	if err != nil {
		log.Println(err)
	}
}

func (s *service) ConfirmUgandaDRFileAvailability() {
	err := try.Do(func(attempt int) (bool, error) {
		try.MaxRetries = 240
		var err error
		err = s.ConfirmFileAvailabilityMethod("/mnt/ugandadr")
		if err != nil {
			log.Println("UgandaDR file not yet detected. Next attempt in 2 minutes...")
			time.Sleep(2 * time.Minute) // wait 2 minutes
		}
		return true, err
	})
	if err != nil {
		log.Println(err)
	}
}

func (s *service) ConfirmZambiaFileAvailability() {
	err := try.Do(func(attempt int) (bool, error) {
		try.MaxRetries = 240
		var err error
		err = s.ConfirmFileAvailabilityMethod("/mnt/zambia")
		if err != nil {
			log.Println("Zambia file not yet detected. Next attempt in 2 minutes...")
			time.Sleep(2 * time.Minute) // wait 2 minutes
		}
		return true, err
	})
	if err != nil {
		log.Println(err)
	}
}

func (s *service) ConfirmZambiaDRFileAvailability() {
	err := try.Do(func(attempt int) (bool, error) {
		try.MaxRetries = 240
		var err error
		err = s.ConfirmFileAvailabilityMethod("/mnt/zambiadr")
		if err != nil {
			log.Println("ZambiaDR file not yet detected. Next attempt in 2 minutes...")
			time.Sleep(2 * time.Minute) // wait 2 minutes
		}
		return true, err
	})
	if err != nil {
		log.Println(err)
	}
}

func (s *service) ConfirmZambiaProdFileAvailability() {
	err := try.Do(func(attempt int) (bool, error) {
		try.MaxRetries = 240
		var err error
		err = s.ConfirmFileAvailabilityMethod("/mnt/zambiaprod")
		if err != nil {
			log.Println("ZambiaProd file not yet detected. Next attempt in 2 minutes...")
			time.Sleep(2 * time.Minute) // wait 2 minutes
		}
		return true, err
	})
	if err != nil {
		log.Println(err)
	}
}
