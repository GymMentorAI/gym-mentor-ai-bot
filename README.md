# Project Name: GymMentorAI

## Overview

GymMentorAI is an innovative Telegram chatbot powered by OpenAI's GPT-4 technology, designed to assist gym-goers by creating personalized workout routines. Developed using GoLang, this bot offers real-time conversation, workout tracking, and tailored fitness advice, making it an essential tool for enhancing fitness experiences through advanced technology.

## Features

- **Personalized Training Routines:** Generate customized workout plans based on user preferences and fitness goals.
- **Interactive Chatbot Experience:** Engage with users through Telegram to provide workout guidance, motivational support, and fitness education.
- **Routine Management:** Users can create, modify, and delete their training routines as their needs evolve.
- **Daily Exercise Reminders:** Automatic reminders to keep users on track with their daily workout schedules.
- **Progress Tracking:** Monitor and report on the user’s progress over time to adjust the training plan as needed.
- **Health and Fitness Tips:** Provide useful tips on health, nutrition, and wellness to complement the physical training.

## Technologies Used

- **OpenAI GPT-4:** Employs the latest in AI technology for sophisticated natural language processing.
- **Telegram API:** Integrates seamlessly with Telegram for user interaction.
- **Aiven DB:** DB SaaS.
- **GoLang:** The primary programming language used for developing the bot and handling backend operations.

## Getting Started

To get started with GymMentorAI, follow these steps:

1. Clone the repository:

   ```bash
   git clone https://github.com/GymMentorAI/gym-mentor-ai-bot.git
   ```

2. Install the required Go packages:

   ```bash
   go get ./
   ```

3. Set up your Telegram bot with the @BotFather account to obtain your API token.

4. Set up a OpenAI ApiKey to use the an AI model and Assistants API V2.

5. Create DB, we are using Aiven DB like a Saas to easy setup. With Free Tier.

6. Create .env with environment variables (copy .env.example and change to .env). Make sure to keep the .env file private and never commit it to your version control system, as it contains sensitive information.

7. Run the bot:

   ```bash
   go run .
   ```

## How to Use

To use the GymMentorAI, send a message to the bot on Telegram. Begin by typing `/start` to initiate interaction. Follow the prompts to set up your profile and preferences. The bot will guide you through creating your personalized workout plan.

For each telegram user, a TelegramUser-ThreadId relationship will be set up.
The OpenAI Assistant Api consists of Assistants, Threads, messages.

## Contributing

Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

To contribute to GymMentorAI, please follow these steps:

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

Distributed under the MIT License. See `LICENSE` for more information.

## Authors & Contact

- n1rocket
- antikorps

Write in issues tab.

Project Link: [https://github.com/GymMentorAI/gym-mentor-ai-bot](https://github.com/GymMentorAI/gym-mentor-ai-bot)
