package cmd

import (
	"fmt"
	"os"

	"git.proserve.nl/statuspage/email"
	"github.com/spf13/cobra"
)

var subscriberCmd = &cobra.Command{
	Use:   "subscriber",
	Short: "Manipulate subscribers",
}

var subscriberListCmd = &cobra.Command{
	Use:     "list",
	Example: "statuspage subscriber list [type]",
	Short:   "list subscribers",
	Long:    `Lists subscribers`,
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		listAllSubscribers := func() {
			subscribers, err := app.Client.GetAllSubscribers()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			for _, subscriber := range subscribers {
				fmt.Println(&subscriber)
			}
		}

		if len(args) < 1 {
			listAllSubscribers()
		} else {
			switch args[0] {
			case "sms": // TODO: Implement SMS subscriber list
				fmt.Println("not implemented yet")
			case "webhook": // TODO: Impelement webhook subscriber list
				fmt.Println("not implemented yet")
			default:
				listAllSubscribers()
			}
		}
	},
}

var subscriberCreateCmd = &cobra.Command{
	Use:     "create",
	Example: "statuspage subscriber create [email@example.org]",
	Short:   "create a subscriber",
	Long: `Create a subscriber through email adres. The subsciber *needs* to
		   confirm the email from statuspage to receive notifications. The
		   subscriber is then added to all the components by default.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		create := args[0]

		if err := email.ValidateFormat(create); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if _, err := app.Client.SearchEmailSubscriber(create); err == nil {
			// If the request could not be processed, it likely already exists.
			fmt.Printf("subscriber already exists: %v\n", create)
			os.Exit(1)
		}

		_, err := app.Client.CreateSubscriber(create)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("succesfully added: %v\n", create)
	},
}

var subscriberDeleteCmd = &cobra.Command{
	Use:     "delete",
	Example: "statuspage subscriber delete [email@example.org]",
	Short:   "Delete a subscriber",
	Long:    "Delete a subscriber through email adres",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		delete := args[0]

		if err := email.ValidateFormat(delete); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		subscriber, err := app.Client.SearchEmailSubscriber(delete)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		_, err = app.Client.DeleteSubscriber(subscriber)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("succesfully deleted: %v\n", *subscriber.Email)
	},
}

var subscriberSearchCmd = &cobra.Command{
	Use:     "search",
	Example: "statuspage subscriber search [email@example.org]",
	Short:   "Search a subscriber",
	Long:    "Search a subscriber through email adres",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		search := args[0]

		if err := email.ValidateFormat(search); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		subscriber, err := app.Client.SearchEmailSubscriber(search)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf(subscriber.String())
	},
}

var subscriberResendCmd = &cobra.Command{
	Use:     "resend",
	Example: "statuspage subscriber resend [email@example.org]",
	Short:   "Resend conformation email",
	Long:    "Resend the conformation email to a subscriber",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		address := args[0]

		if err := email.ValidateFormat(address); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		subscriber, err := app.Client.SearchEmailSubscriber(address)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if app.Debug {
			fmt.Println("Found: ", subscriber)
		}

		_, err = app.Client.ResendConformationEmail(subscriber)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("succesfully resend conformation email to: %v\n", *subscriber.Email)
	},
}

func init() {
	RootCmd.AddCommand(subscriberCmd)
	subscriberCmd.AddCommand(subscriberCreateCmd)
	subscriberCmd.AddCommand(subscriberDeleteCmd)
	subscriberCmd.AddCommand(subscriberListCmd)
	subscriberCmd.AddCommand(subscriberSearchCmd)
	subscriberCmd.AddCommand(subscriberResendCmd)
}
