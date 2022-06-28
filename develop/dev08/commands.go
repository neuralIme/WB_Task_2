package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

// echo выводит аргумент в stdout
func echo(cmd []string) error {
	_, err := fmt.Println(strings.Join(cmd, " "))
	if err != nil {
		return err
	}
	return nil
}

// pwd выводит путь до текущего каталога
func pwd() error {
	path, err := os.Getwd()
	if err != nil {
		return errors.New("не удалось определить путь")
	}
	fmt.Println(path)
	return nil
}

// cd смена директории (без аргумента переходит в корневой каталог)
func cd(cmd []string) error {
	if len(cmd) == 0 {
		_ = os.Chdir("/Users/korolev")
		return nil
	}
	err := os.Chdir(cmd[0])
	if err != nil {
		return errors.New("не удалось сменить каталог")
	}
	return nil
}

// ps выводит общую информацию по запущенным процессам
func ps() error {
	cmd := exec.Command("ps")
	proc, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New("не удалось получить данные о процессах")
	}
	fmt.Printf("%s", proc)
	return nil
}

// ls выводит содержимое каталога
func ls() error {
	cmd := exec.Command("ls")
	dir, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New("не удалось определить содержимое каталога")
	}
	fmt.Printf("%s", dir)
	return nil
}

// kill процесс, переданный в качесте аргумента по PID
func kill(arg []string) error {
	pid, err := strconv.Atoi(arg[0])
	if err != nil {
		return errors.New("некорректный ввод: " + arg[0])
	}
	proc, err := os.FindProcess(pid)
	if err != nil {
		return errors.New("процесс не найден")
	}
	err = proc.Kill()
	if err != nil {
		return errors.New("не удалось убить процесс")
	}
	return nil
}

// exe заменяет процесс на указанный бинарный файл
func exe(args []string) error {
	bin, err := exec.LookPath(args[0])
	if err != nil {
		return errors.New("исполняемый файл не найден")
	}

	env := os.Environ()

	err = syscall.Exec(bin, args, env)
	if err != nil {
		return errors.New("не удалось заменить процесс")
	}

	return nil
}

// fork создает дочерний процесс
func fork() {
	pid, _, err := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
	if err != 0 {
		log.Println("неудалось создать процесс", err)
	}
	fmt.Printf("%d\n", pid)
}
