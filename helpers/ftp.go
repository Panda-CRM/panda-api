package helpers

import (
	"errors"
	"github.com/jlaffaye/ftp"
	"io/ioutil"
	"log"
	"os"
)

var (
	errConnectFTP   = errors.New("Houve um erro ao conectar ao FTP")
	errAuthFTP      = errors.New("Houve um erro ao autenticar no FTP")
	errGetListFiles = errors.New("Houve um erro ao listar arquivos do FTP")
	errReadDataFTP  = errors.New("Houve um erro ao ler os dados do arquivo FTP")
	errWriteFileFTP = errors.New("Houve um erro ao gravar os dados do FTP")
	errImportFTP    = errors.New("Houve um erro ao fazer download FTP")
)

type FTPClient struct {
	Address        string
	Username       string
	Password       string
	FtpModePassive bool
}

func (m *FTPClient) buildClientFTP() (*ftp.ServerConn, error) {
	client, err := ftp.Dial(m.Address)
	if err != nil {
		log.Println(err)
		return nil, errConnectFTP
	}

	client.DisableEPSV = m.FtpModePassive
	if err := client.Login(m.Username, m.Password); err != nil {
		log.Println(err)
		return nil, errAuthFTP
	}

	return client, nil
}

func (m *FTPClient) DownloadFiles(path, distination string) error {
	client, err := m.buildClientFTP()
	if err != nil {
		return err
	}

	entries, _ := client.List(path)
	for _, entry := range entries {
		if entry.Type == ftp.EntryTypeFile {
			reader, err := client.Retr(entry.Name)
			if err != nil {
				log.Println(err)
				return errGetListFiles
			}

			buf, err := ioutil.ReadAll(reader)
			reader.Close()
			if err != nil {
				log.Println(err)
				return errReadDataFTP
			}

			m.writeLocalFile(distination+entry.Name, buf, 0644)
		}
	}

	defer client.Quit()
	return nil
}

func (m *FTPClient) writeLocalFile(filename string, data []byte, perm os.FileMode) error {
	err := ioutil.WriteFile(filename, data, perm)
	if err != nil {
		log.Println(err)
		return errWriteFileFTP
	}
	return nil
}

func (m *FTPClient) MoveRemoteFiles(entries []*ftp.Entry, destination string) error {
	client, _ := m.buildClientFTP()

	defer client.Quit()
	return nil
}
